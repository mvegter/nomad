// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

package feasible

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/nomad/ci"
	"github.com/hashicorp/nomad/helper/uuid"
	"github.com/hashicorp/nomad/nomad/mock"
	"github.com/hashicorp/nomad/nomad/structs"
	psstructs "github.com/hashicorp/nomad/plugins/shared/structs"
	"github.com/hashicorp/nomad/scheduler/tests"
	"github.com/shoenig/test/must"
)

func TestResourceDistance(t *testing.T) {
	ci.Parallel(t)

	resourceAsk := &structs.ComparableResources{
		Flattened: structs.AllocatedTaskResources{
			Cpu: structs.AllocatedCpuResources{
				CpuShares: 2048,
			},
			Memory: structs.AllocatedMemoryResources{
				MemoryMB: 512,
			},
			Networks: []*structs.NetworkResource{
				{
					Device: "eth0",
					MBits:  1024,
				},
			},
		},
		Shared: structs.AllocatedSharedResources{
			DiskMB: 4096,
		},
	}

	type testCase struct {
		allocResource    *structs.ComparableResources
		expectedDistance string
	}

	testCases := []*testCase{
		{
			&structs.ComparableResources{
				Flattened: structs.AllocatedTaskResources{
					Cpu: structs.AllocatedCpuResources{
						CpuShares: 2048,
					},
					Memory: structs.AllocatedMemoryResources{
						MemoryMB: 512,
					},
					Networks: []*structs.NetworkResource{
						{
							Device: "eth0",
							MBits:  1024,
						},
					},
				},
				Shared: structs.AllocatedSharedResources{
					DiskMB: 4096,
				},
			},
			"0.000",
		},
		{
			&structs.ComparableResources{
				Flattened: structs.AllocatedTaskResources{
					Cpu: structs.AllocatedCpuResources{
						CpuShares: 1024,
					},
					Memory: structs.AllocatedMemoryResources{
						MemoryMB: 400,
					},
					Networks: []*structs.NetworkResource{
						{
							Device: "eth0",
							MBits:  1024,
						},
					},
				},
				Shared: structs.AllocatedSharedResources{
					DiskMB: 1024,
				},
			},
			"0.928",
		},
		{
			&structs.ComparableResources{
				Flattened: structs.AllocatedTaskResources{
					Cpu: structs.AllocatedCpuResources{
						CpuShares: 8192,
					},
					Memory: structs.AllocatedMemoryResources{
						MemoryMB: 200,
					},
					Networks: []*structs.NetworkResource{
						{
							Device: "eth0",
							MBits:  512,
						},
					},
				},
				Shared: structs.AllocatedSharedResources{
					DiskMB: 1024,
				},
			},
			"3.152",
		},
		{
			&structs.ComparableResources{
				Flattened: structs.AllocatedTaskResources{
					Cpu: structs.AllocatedCpuResources{
						CpuShares: 2048,
					},
					Memory: structs.AllocatedMemoryResources{
						MemoryMB: 500,
					},
					Networks: []*structs.NetworkResource{
						{
							Device: "eth0",
							MBits:  1024,
						},
					},
				},
				Shared: structs.AllocatedSharedResources{
					DiskMB: 4096,
				},
			},
			"0.023",
		},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			actualDistance := fmt.Sprintf("%3.3f", basicResourceDistance(resourceAsk, tc.allocResource))
			must.Eq(t, tc.expectedDistance, actualDistance)
		})

	}

}

func makeDeviceInstance(instanceID, busID string) *structs.NodeDevice {
	return &structs.NodeDevice{
		ID:      instanceID,
		Healthy: true,
		Locality: &structs.NodeDeviceLocality{
			PciBusID: busID,
		},
	}
}

func TestPreemption_Normal(t *testing.T) {
	ci.Parallel(t)

	type testCase struct {
		desc                 string
		currentAllocations   []*structs.Allocation
		nodeReservedCapacity *structs.NodeReservedResources
		nodeCapacity         *structs.NodeResources
		resourceAsk          *structs.Resources
		jobPriority          int
		currentPreemptions   []*structs.Allocation
		preemptedAllocIDs    map[string]struct{}
	}

	highPrioJob := mock.Job()
	highPrioJob.Priority = 100

	lowPrioJob := mock.Job()
	lowPrioJob.Priority = 30

	lowPrioJob2 := mock.Job()
	lowPrioJob2.Priority = 40

	// Create some persistent alloc ids to use in test cases
	allocIDs := []string{uuid.Generate(), uuid.Generate(), uuid.Generate(), uuid.Generate(), uuid.Generate(), uuid.Generate()}

	var deviceIDs []string
	for i := 0; i < 10; i++ {
		deviceIDs = append(deviceIDs, "dev"+strconv.Itoa(i))
	}

	legacyCpuResources, processorResources := tests.CpuResources(4000)

	defaultNodeResources := &structs.NodeResources{
		Processors: processorResources,
		Cpu:        legacyCpuResources,

		Memory: structs.NodeMemoryResources{
			MemoryMB: 8192,
		},
		Disk: structs.NodeDiskResources{
			DiskMB: 100 * 1024,
		},
		Networks: []*structs.NetworkResource{
			{
				Device: "eth0",
				CIDR:   "192.168.0.100/32",
				MBits:  1000,
			},
		},
		Devices: []*structs.NodeDeviceResource{
			{
				Type:   "gpu",
				Vendor: "nvidia",
				Name:   "1080ti",
				Attributes: map[string]*psstructs.Attribute{
					"memory":           psstructs.NewIntAttribute(11, psstructs.UnitGiB),
					"cuda_cores":       psstructs.NewIntAttribute(3584, ""),
					"graphics_clock":   psstructs.NewIntAttribute(1480, psstructs.UnitMHz),
					"memory_bandwidth": psstructs.NewIntAttribute(11, psstructs.UnitGBPerS),
				},
				Instances: []*structs.NodeDevice{
					makeDeviceInstance(deviceIDs[0], "0000:00:00.0"),
					makeDeviceInstance(deviceIDs[1], "0000:00:01.0"),
					makeDeviceInstance(deviceIDs[2], "0000:00:02.0"),
					makeDeviceInstance(deviceIDs[3], "0000:00:03.0"),
				},
			},
			{
				Type:   "gpu",
				Vendor: "nvidia",
				Name:   "2080ti",
				Attributes: map[string]*psstructs.Attribute{
					"memory":           psstructs.NewIntAttribute(11, psstructs.UnitGiB),
					"cuda_cores":       psstructs.NewIntAttribute(3584, ""),
					"graphics_clock":   psstructs.NewIntAttribute(1480, psstructs.UnitMHz),
					"memory_bandwidth": psstructs.NewIntAttribute(11, psstructs.UnitGBPerS),
				},
				Instances: []*structs.NodeDevice{
					makeDeviceInstance(deviceIDs[4], "0000:00:04.0"),
					makeDeviceInstance(deviceIDs[5], "0000:00:05.0"),
					makeDeviceInstance(deviceIDs[6], "0000:00:06.0"),
					makeDeviceInstance(deviceIDs[7], "0000:00:07.0"),
					makeDeviceInstance(deviceIDs[8], "0000:00:08.0"),
				},
			},
			{
				Type:   "fpga",
				Vendor: "intel",
				Name:   "F100",
				Attributes: map[string]*psstructs.Attribute{
					"memory": psstructs.NewIntAttribute(4, psstructs.UnitGiB),
				},
				Instances: []*structs.NodeDevice{
					makeDeviceInstance("fpga1", "0000:01:00.0"),
					makeDeviceInstance("fpga2", "0000:02:01.0"),
				},
			},
		},
	}

	reservedNodeResources := &structs.NodeReservedResources{
		Cpu: structs.NodeReservedCpuResources{
			CpuShares: 100,
		},
		Memory: structs.NodeReservedMemoryResources{
			MemoryMB: 256,
		},
		Disk: structs.NodeReservedDiskResources{
			DiskMB: 4 * 1024,
		},
	}

	testCases := []testCase{
		{
			desc: "No preemption because existing allocs are not low priority",
			currentAllocations: []*structs.Allocation{
				tests.CreateAlloc(allocIDs[0], highPrioJob, &structs.Resources{
					CPU:      3200,
					MemoryMB: 7256,
					DiskMB:   4 * 1024,
					Networks: []*structs.NetworkResource{
						{
							Device: "eth0",
							IP:     "192.168.0.100",
							MBits:  50,
						},
					},
				})},
			nodeReservedCapacity: reservedNodeResources,
			nodeCapacity:         defaultNodeResources,
			jobPriority:          100,
			resourceAsk: &structs.Resources{
				CPU:      2000,
				MemoryMB: 256,
				DiskMB:   4 * 1024,
				Networks: []*structs.NetworkResource{
					{
						Device:        "eth0",
						IP:            "192.168.0.100",
						ReservedPorts: []structs.Port{{Label: "ssh", Value: 22}},
						MBits:         1,
					},
				},
			},
		},
		{
			desc: "Preempting low priority allocs not enough to meet resource ask",
			currentAllocations: []*structs.Allocation{
				tests.CreateAlloc(allocIDs[0], lowPrioJob, &structs.Resources{
					CPU:      3200,
					MemoryMB: 7256,
					DiskMB:   4 * 1024,
					Networks: []*structs.NetworkResource{
						{
							Device: "eth0",
							IP:     "192.168.0.100",
							MBits:  50,
						},
					},
				})},
			nodeReservedCapacity: reservedNodeResources,
			nodeCapacity:         defaultNodeResources,
			jobPriority:          100,
			resourceAsk: &structs.Resources{
				CPU:      4000,
				MemoryMB: 8192,
				DiskMB:   4 * 1024,
				Networks: []*structs.NetworkResource{
					{
						Device:        "eth0",
						IP:            "192.168.0.100",
						ReservedPorts: []structs.Port{{Label: "ssh", Value: 22}},
						MBits:         1,
					},
				},
			},
		},
		{
			desc: "preemption impossible - static port needed is used by higher priority alloc",
			currentAllocations: []*structs.Allocation{
				tests.CreateAlloc(allocIDs[0], highPrioJob, &structs.Resources{
					CPU:      1200,
					MemoryMB: 2256,
					DiskMB:   4 * 1024,
					Networks: []*structs.NetworkResource{
						{
							Device: "eth0",
							IP:     "192.168.0.100",
							MBits:  150,
						},
					},
				}),
				tests.CreateAlloc(allocIDs[1], highPrioJob, &structs.Resources{
					CPU:      200,
					MemoryMB: 256,
					DiskMB:   4 * 1024,
					Networks: []*structs.NetworkResource{
						{
							Device: "eth0",
							IP:     "192.168.0.200",
							MBits:  600,
							ReservedPorts: []structs.Port{
								{
									Label: "db",
									Value: 88,
								},
							},
						},
					},
				}),
			},
			nodeReservedCapacity: reservedNodeResources,
			nodeCapacity:         defaultNodeResources,
			jobPriority:          100,
			resourceAsk: &structs.Resources{
				CPU:      600,
				MemoryMB: 1000,
				DiskMB:   25 * 1024,
				Networks: []*structs.NetworkResource{
					{
						Device: "eth0",
						IP:     "192.168.0.100",
						MBits:  700,
						ReservedPorts: []structs.Port{
							{
								Label: "db",
								Value: 88,
							},
						},
					},
				},
			},
		},
		{
			desc: "preempt only from device that has allocation with unused reserved port",
			currentAllocations: []*structs.Allocation{
				tests.CreateAlloc(allocIDs[0], highPrioJob, &structs.Resources{
					CPU:      1200,
					MemoryMB: 2256,
					DiskMB:   4 * 1024,
					Networks: []*structs.NetworkResource{
						{
							Device: "eth0",
							IP:     "192.168.0.100",
							MBits:  150,
						},
					},
				}),
				tests.CreateAlloc(allocIDs[1], highPrioJob, &structs.Resources{
					CPU:      200,
					MemoryMB: 256,
					DiskMB:   4 * 1024,
					Networks: []*structs.NetworkResource{
						{
							Device: "eth1",
							IP:     "192.168.0.200",
							MBits:  600,
							ReservedPorts: []structs.Port{
								{
									Label: "db",
									Value: 88,
								},
							},
						},
					},
				}),
				tests.CreateAlloc(allocIDs[2], lowPrioJob, &structs.Resources{
					CPU:      200,
					MemoryMB: 256,
					DiskMB:   4 * 1024,
					Networks: []*structs.NetworkResource{
						{
							Device: "eth0",
							IP:     "192.168.0.200",
							MBits:  600,
						},
					},
				}),
			},
			nodeReservedCapacity: reservedNodeResources,
			// This test sets up a node with two NICs

			nodeCapacity: &structs.NodeResources{
				Processors: processorResources,
				Cpu:        legacyCpuResources,
				Memory: structs.NodeMemoryResources{
					MemoryMB: 8192,
				},
				Disk: structs.NodeDiskResources{
					DiskMB: 100 * 1024,
				},
				Networks: []*structs.NetworkResource{
					{
						Device: "eth0",
						CIDR:   "192.168.0.100/32",
						MBits:  1000,
					},
					{
						Device: "eth1",
						CIDR:   "192.168.1.100/32",
						MBits:  1000,
					},
				},
			},
			jobPriority: 100,
			resourceAsk: &structs.Resources{
				CPU:      600,
				MemoryMB: 1000,
				DiskMB:   25 * 1024,
				Networks: []*structs.NetworkResource{
					{
						IP:    "192.168.0.100",
						MBits: 700,
						ReservedPorts: []structs.Port{
							{
								Label: "db",
								Value: 88,
							},
						},
					},
				},
			},
			preemptedAllocIDs: map[string]struct{}{
				allocIDs[2]: {},
			},
		},
		{
			desc: "Combination of high/low priority allocs, without static ports",
			currentAllocations: []*structs.Allocation{
				tests.CreateAlloc(allocIDs[0], highPrioJob, &structs.Resources{
					CPU:      2800,
					MemoryMB: 2256,
					DiskMB:   4 * 1024,
					Networks: []*structs.NetworkResource{
						{
							Device: "eth0",
							IP:     "192.168.0.100",
							MBits:  150,
						},
					},
				}),
				tests.CreateAllocWithTaskgroupNetwork(allocIDs[1], lowPrioJob, &structs.Resources{
					CPU:      200,
					MemoryMB: 256,
					DiskMB:   4 * 1024,
					Networks: []*structs.NetworkResource{
						{
							Device: "eth0",
							IP:     "192.168.0.200",
							MBits:  200,
						},
					},
				}, &structs.NetworkResource{
					Device: "eth0",
					IP:     "192.168.0.201",
					MBits:  300,
				}),
				tests.CreateAlloc(allocIDs[2], lowPrioJob, &structs.Resources{
					CPU:      200,
					MemoryMB: 256,
					DiskMB:   4 * 1024,
					Networks: []*structs.NetworkResource{
						{
							Device: "eth0",
							IP:     "192.168.0.100",
							MBits:  300,
						},
					},
				}),
				tests.CreateAlloc(allocIDs[3], lowPrioJob, &structs.Resources{
					CPU:      700,
					MemoryMB: 256,
					DiskMB:   4 * 1024,
				}),
			},
			nodeReservedCapacity: reservedNodeResources,
			nodeCapacity:         defaultNodeResources,
			jobPriority:          100,
			resourceAsk: &structs.Resources{
				CPU:      1100,
				MemoryMB: 1000,
				DiskMB:   25 * 1024,
				Networks: []*structs.NetworkResource{
					{
						Device: "eth0",
						IP:     "192.168.0.100",
						MBits:  840,
					},
				},
			},
			preemptedAllocIDs: map[string]struct{}{
				allocIDs[1]: {},
				allocIDs[2]: {},
				allocIDs[3]: {},
			},
		},
		{
			desc: "preempt allocs with network devices",
			currentAllocations: []*structs.Allocation{
				tests.CreateAlloc(allocIDs[0], lowPrioJob, &structs.Resources{
					CPU:      2800,
					MemoryMB: 2256,
					DiskMB:   4 * 1024,
				}),
				tests.CreateAlloc(allocIDs[1], lowPrioJob, &structs.Resources{
					CPU:      200,
					MemoryMB: 256,
					DiskMB:   4 * 1024,
					Networks: []*structs.NetworkResource{
						{
							Device: "eth0",
							IP:     "192.168.0.200",
							MBits:  800,
						},
					},
				}),
			},
			nodeReservedCapacity: reservedNodeResources,
			nodeCapacity:         defaultNodeResources,
			jobPriority:          100,
			resourceAsk: &structs.Resources{
				CPU:      1100,
				MemoryMB: 1000,
				DiskMB:   25 * 1024,
				Networks: []*structs.NetworkResource{
					{
						Device: "eth0",
						IP:     "192.168.0.100",
						MBits:  840,
					},
				},
			},
			preemptedAllocIDs: map[string]struct{}{
				allocIDs[1]: {},
			},
		},
		{
			desc: "ignore allocs with close enough priority for network devices",
			currentAllocations: []*structs.Allocation{
				tests.CreateAlloc(allocIDs[0], lowPrioJob, &structs.Resources{
					CPU:      2800,
					MemoryMB: 2256,
					DiskMB:   4 * 1024,
				}),
				tests.CreateAlloc(allocIDs[1], lowPrioJob, &structs.Resources{
					CPU:      200,
					MemoryMB: 256,
					DiskMB:   4 * 1024,
					Networks: []*structs.NetworkResource{
						{
							Device: "eth0",
							IP:     "192.168.0.200",
							MBits:  800,
						},
					},
				}),
			},
			nodeReservedCapacity: reservedNodeResources,
			nodeCapacity:         defaultNodeResources,
			jobPriority:          lowPrioJob.Priority + 5,
			resourceAsk: &structs.Resources{
				CPU:      1100,
				MemoryMB: 1000,
				DiskMB:   25 * 1024,
				Networks: []*structs.NetworkResource{
					{
						Device: "eth0",
						IP:     "192.168.0.100",
						MBits:  840,
					},
				},
			},
			preemptedAllocIDs: nil,
		},
		{
			desc: "Preemption needed for all resources except network",
			currentAllocations: []*structs.Allocation{
				tests.CreateAlloc(allocIDs[0], highPrioJob, &structs.Resources{
					CPU:      2800,
					MemoryMB: 2256,
					DiskMB:   40 * 1024,
					Networks: []*structs.NetworkResource{
						{
							Device: "eth0",
							IP:     "192.168.0.100",
							MBits:  150,
						},
					},
				}),
				tests.CreateAlloc(allocIDs[1], lowPrioJob, &structs.Resources{
					CPU:      200,
					MemoryMB: 256,
					DiskMB:   4 * 1024,
					Networks: []*structs.NetworkResource{
						{
							Device: "eth0",
							IP:     "192.168.0.200",
							MBits:  50,
						},
					},
				}),
				tests.CreateAlloc(allocIDs[2], lowPrioJob, &structs.Resources{
					CPU:      200,
					MemoryMB: 512,
					DiskMB:   25 * 1024,
				}),
				tests.CreateAlloc(allocIDs[3], lowPrioJob, &structs.Resources{
					CPU:      700,
					MemoryMB: 276,
					DiskMB:   20 * 1024,
				}),
			},
			nodeReservedCapacity: reservedNodeResources,
			nodeCapacity:         defaultNodeResources,
			jobPriority:          100,
			resourceAsk: &structs.Resources{
				CPU:      1000,
				MemoryMB: 3000,
				DiskMB:   50 * 1024,
				Networks: []*structs.NetworkResource{
					{
						Device: "eth0",
						IP:     "192.168.0.100",
						MBits:  50,
					},
				},
			},
			preemptedAllocIDs: map[string]struct{}{
				allocIDs[1]: {},
				allocIDs[2]: {},
				allocIDs[3]: {},
			},
		},
		{
			desc: "Only one low priority alloc needs to be preempted",
			currentAllocations: []*structs.Allocation{
				tests.CreateAlloc(allocIDs[0], highPrioJob, &structs.Resources{
					CPU:      1200,
					MemoryMB: 2256,
					DiskMB:   4 * 1024,
					Networks: []*structs.NetworkResource{
						{
							Device: "eth0",
							IP:     "192.168.0.100",
							MBits:  150,
						},
					},
				}),
				tests.CreateAlloc(allocIDs[1], lowPrioJob, &structs.Resources{
					CPU:      200,
					MemoryMB: 256,
					DiskMB:   4 * 1024,
					Networks: []*structs.NetworkResource{
						{
							Device: "eth0",
							IP:     "192.168.0.100",
							MBits:  500,
						},
					},
				}),
				tests.CreateAlloc(allocIDs[2], lowPrioJob, &structs.Resources{
					CPU:      200,
					MemoryMB: 256,
					DiskMB:   4 * 1024,
					Networks: []*structs.NetworkResource{
						{
							Device: "eth0",
							IP:     "192.168.0.200",
							MBits:  320,
						},
					},
				}),
			},
			nodeReservedCapacity: reservedNodeResources,
			nodeCapacity:         defaultNodeResources,
			jobPriority:          100,
			resourceAsk: &structs.Resources{
				CPU:      300,
				MemoryMB: 500,
				DiskMB:   5 * 1024,
				Networks: []*structs.NetworkResource{
					{
						Device: "eth0",
						IP:     "192.168.0.100",
						MBits:  320,
					},
				},
			},
			preemptedAllocIDs: map[string]struct{}{
				allocIDs[2]: {},
			},
		},
		{
			desc: "one alloc meets static port need, another meets remaining mbits needed",
			currentAllocations: []*structs.Allocation{
				tests.CreateAlloc(allocIDs[0], highPrioJob, &structs.Resources{
					CPU:      1200,
					MemoryMB: 2256,
					DiskMB:   4 * 1024,
					Networks: []*structs.NetworkResource{
						{
							Device: "eth0",
							IP:     "192.168.0.100",
							MBits:  150,
						},
					},
				}),
				tests.CreateAlloc(allocIDs[1], lowPrioJob, &structs.Resources{
					CPU:      200,
					MemoryMB: 256,
					DiskMB:   4 * 1024,
					Networks: []*structs.NetworkResource{
						{
							Device: "eth0",
							IP:     "192.168.0.200",
							MBits:  500,
							ReservedPorts: []structs.Port{
								{
									Label: "db",
									Value: 88,
								},
							},
						},
					},
				}),
				tests.CreateAlloc(allocIDs[2], lowPrioJob, &structs.Resources{
					CPU:      200,
					MemoryMB: 256,
					DiskMB:   4 * 1024,
					Networks: []*structs.NetworkResource{
						{
							Device: "eth0",
							IP:     "192.168.0.100",
							MBits:  200,
						},
					},
				}),
			},
			nodeReservedCapacity: reservedNodeResources,
			nodeCapacity:         defaultNodeResources,
			jobPriority:          100,
			resourceAsk: &structs.Resources{
				CPU:      2700,
				MemoryMB: 1000,
				DiskMB:   25 * 1024,
				Networks: []*structs.NetworkResource{
					{
						Device: "eth0",
						IP:     "192.168.0.100",
						MBits:  800,
						ReservedPorts: []structs.Port{
							{
								Label: "db",
								Value: 88,
							},
						},
					},
				},
			},
			preemptedAllocIDs: map[string]struct{}{
				allocIDs[1]: {},
				allocIDs[2]: {},
			},
		},
		{
			desc: "alloc that meets static port need also meets other needs",
			currentAllocations: []*structs.Allocation{
				tests.CreateAlloc(allocIDs[0], highPrioJob, &structs.Resources{
					CPU:      1200,
					MemoryMB: 2256,
					DiskMB:   4 * 1024,
					Networks: []*structs.NetworkResource{
						{
							Device: "eth0",
							IP:     "192.168.0.100",
							MBits:  150,
						},
					},
				}),
				tests.CreateAlloc(allocIDs[1], lowPrioJob, &structs.Resources{
					CPU:      200,
					MemoryMB: 256,
					DiskMB:   4 * 1024,
					Networks: []*structs.NetworkResource{
						{
							Device: "eth0",
							IP:     "192.168.0.200",
							MBits:  600,
							ReservedPorts: []structs.Port{
								{
									Label: "db",
									Value: 88,
								},
							},
						},
					},
				}),
				tests.CreateAlloc(allocIDs[2], lowPrioJob, &structs.Resources{
					CPU:      200,
					MemoryMB: 256,
					DiskMB:   4 * 1024,
					Networks: []*structs.NetworkResource{
						{
							Device: "eth0",
							IP:     "192.168.0.100",
							MBits:  100,
						},
					},
				}),
			},
			nodeReservedCapacity: reservedNodeResources,
			nodeCapacity:         defaultNodeResources,
			jobPriority:          100,
			resourceAsk: &structs.Resources{
				CPU:      600,
				MemoryMB: 1000,
				DiskMB:   25 * 1024,
				Networks: []*structs.NetworkResource{
					{
						Device: "eth0",
						IP:     "192.168.0.100",
						MBits:  700,
						ReservedPorts: []structs.Port{
							{
								Label: "db",
								Value: 88,
							},
						},
					},
				},
			},
			preemptedAllocIDs: map[string]struct{}{
				allocIDs[1]: {},
			},
		},
		{
			desc: "alloc from job that has existing evictions not chosen for preemption",
			currentAllocations: []*structs.Allocation{
				tests.CreateAlloc(allocIDs[0], highPrioJob, &structs.Resources{
					CPU:      1200,
					MemoryMB: 2256,
					DiskMB:   4 * 1024,
					Networks: []*structs.NetworkResource{
						{
							Device: "eth0",
							IP:     "192.168.0.100",
							MBits:  150,
						},
					},
				}),
				tests.CreateAlloc(allocIDs[1], lowPrioJob, &structs.Resources{
					CPU:      200,
					MemoryMB: 256,
					DiskMB:   4 * 1024,
					Networks: []*structs.NetworkResource{
						{
							Device: "eth0",
							IP:     "192.168.0.200",
							MBits:  500,
						},
					},
				}),
				tests.CreateAlloc(allocIDs[2], lowPrioJob2, &structs.Resources{
					CPU:      200,
					MemoryMB: 256,
					DiskMB:   4 * 1024,
					Networks: []*structs.NetworkResource{
						{
							Device: "eth0",
							IP:     "192.168.0.100",
							MBits:  300,
						},
					},
				}),
			},
			nodeReservedCapacity: reservedNodeResources,
			nodeCapacity:         defaultNodeResources,
			jobPriority:          100,
			resourceAsk: &structs.Resources{
				CPU:      300,
				MemoryMB: 500,
				DiskMB:   5 * 1024,
				Networks: []*structs.NetworkResource{
					{
						Device: "eth0",
						IP:     "192.168.0.100",
						MBits:  320,
					},
				},
			},
			currentPreemptions: []*structs.Allocation{
				tests.CreateAlloc(allocIDs[4], lowPrioJob2, &structs.Resources{
					CPU:      200,
					MemoryMB: 256,
					DiskMB:   4 * 1024,
					Networks: []*structs.NetworkResource{
						{
							Device: "eth0",
							IP:     "192.168.0.100",
							MBits:  300,
						},
					},
				}),
			},
			preemptedAllocIDs: map[string]struct{}{
				allocIDs[1]: {},
			},
		},
		{
			desc: "Preemption with one device instance per alloc",
			// Add allocations that use two device instances
			currentAllocations: []*structs.Allocation{
				tests.CreateAllocWithDevice(allocIDs[0], lowPrioJob, &structs.Resources{
					CPU:      500,
					MemoryMB: 512,
					DiskMB:   4 * 1024,
				}, &structs.AllocatedDeviceResource{
					Type:      "gpu",
					Vendor:    "nvidia",
					Name:      "1080ti",
					DeviceIDs: []string{deviceIDs[0]},
				}),
				tests.CreateAllocWithDevice(allocIDs[1], lowPrioJob, &structs.Resources{
					CPU:      200,
					MemoryMB: 512,
					DiskMB:   4 * 1024,
				}, &structs.AllocatedDeviceResource{
					Type:      "gpu",
					Vendor:    "nvidia",
					Name:      "1080ti",
					DeviceIDs: []string{deviceIDs[1]},
				})},
			nodeReservedCapacity: reservedNodeResources,
			nodeCapacity:         defaultNodeResources,
			jobPriority:          100,
			resourceAsk: &structs.Resources{
				CPU:      1000,
				MemoryMB: 512,
				DiskMB:   4 * 1024,
				Devices: []*structs.RequestedDevice{
					{
						Name:  "nvidia/gpu/1080ti",
						Count: 4,
					},
				},
			},
			preemptedAllocIDs: map[string]struct{}{
				allocIDs[0]: {},
				allocIDs[1]: {},
			},
		},
		{
			desc: "Preemption multiple devices used",
			currentAllocations: []*structs.Allocation{
				tests.CreateAllocWithDevice(allocIDs[0], lowPrioJob, &structs.Resources{
					CPU:      500,
					MemoryMB: 512,
					DiskMB:   4 * 1024,
				}, &structs.AllocatedDeviceResource{
					Type:      "gpu",
					Vendor:    "nvidia",
					Name:      "1080ti",
					DeviceIDs: []string{deviceIDs[0], deviceIDs[1], deviceIDs[2], deviceIDs[3]},
				}),
				tests.CreateAllocWithDevice(allocIDs[1], lowPrioJob, &structs.Resources{
					CPU:      200,
					MemoryMB: 512,
					DiskMB:   4 * 1024,
				}, &structs.AllocatedDeviceResource{
					Type:      "fpga",
					Vendor:    "intel",
					Name:      "F100",
					DeviceIDs: []string{"fpga1"},
				})},
			nodeReservedCapacity: reservedNodeResources,
			nodeCapacity:         defaultNodeResources,
			jobPriority:          100,
			resourceAsk: &structs.Resources{
				CPU:      1000,
				MemoryMB: 512,
				DiskMB:   4 * 1024,
				Devices: []*structs.RequestedDevice{
					{
						Name:  "nvidia/gpu/1080ti",
						Count: 4,
					},
				},
			},
			preemptedAllocIDs: map[string]struct{}{
				allocIDs[0]: {},
			},
		},
		{
			// This test cases creates allocations across two GPUs
			// Both GPUs are eligible for the task, but only allocs sharing the
			// same device should be chosen for preemption
			desc: "Preemption with allocs across multiple devices that match",
			currentAllocations: []*structs.Allocation{
				tests.CreateAllocWithDevice(allocIDs[0], lowPrioJob, &structs.Resources{
					CPU:      500,
					MemoryMB: 512,
					DiskMB:   4 * 1024,
				}, &structs.AllocatedDeviceResource{
					Type:      "gpu",
					Vendor:    "nvidia",
					Name:      "1080ti",
					DeviceIDs: []string{deviceIDs[0], deviceIDs[1]},
				}),
				tests.CreateAllocWithDevice(allocIDs[1], highPrioJob, &structs.Resources{
					CPU:      200,
					MemoryMB: 100,
					DiskMB:   4 * 1024,
				}, &structs.AllocatedDeviceResource{
					Type:      "gpu",
					Vendor:    "nvidia",
					Name:      "1080ti",
					DeviceIDs: []string{deviceIDs[2]},
				}),
				tests.CreateAllocWithDevice(allocIDs[2], lowPrioJob, &structs.Resources{
					CPU:      200,
					MemoryMB: 256,
					DiskMB:   4 * 1024,
				}, &structs.AllocatedDeviceResource{
					Type:      "gpu",
					Vendor:    "nvidia",
					Name:      "2080ti",
					DeviceIDs: []string{deviceIDs[4], deviceIDs[5]},
				}),
				tests.CreateAllocWithDevice(allocIDs[3], lowPrioJob, &structs.Resources{
					CPU:      100,
					MemoryMB: 256,
					DiskMB:   4 * 1024,
				}, &structs.AllocatedDeviceResource{
					Type:      "gpu",
					Vendor:    "nvidia",
					Name:      "2080ti",
					DeviceIDs: []string{deviceIDs[6], deviceIDs[7]},
				}),
				tests.CreateAllocWithDevice(allocIDs[4], lowPrioJob, &structs.Resources{
					CPU:      200,
					MemoryMB: 512,
					DiskMB:   4 * 1024,
				}, &structs.AllocatedDeviceResource{
					Type:      "fpga",
					Vendor:    "intel",
					Name:      "F100",
					DeviceIDs: []string{"fpga1"},
				})},
			nodeReservedCapacity: reservedNodeResources,
			nodeCapacity:         defaultNodeResources,
			jobPriority:          100,
			resourceAsk: &structs.Resources{
				CPU:      1000,
				MemoryMB: 512,
				DiskMB:   4 * 1024,
				Devices: []*structs.RequestedDevice{
					{
						Name:  "gpu",
						Count: 4,
					},
				},
			},
			preemptedAllocIDs: map[string]struct{}{
				allocIDs[2]: {},
				allocIDs[3]: {},
			},
		},
		{
			// This test cases creates allocations across two GPUs
			// Both GPUs are eligible for the task, but only allocs with the lower
			// priority are chosen
			desc: "Preemption with lower/higher priority combinations",
			currentAllocations: []*structs.Allocation{
				tests.CreateAllocWithDevice(allocIDs[0], lowPrioJob, &structs.Resources{
					CPU:      500,
					MemoryMB: 512,
					DiskMB:   4 * 1024,
				}, &structs.AllocatedDeviceResource{
					Type:      "gpu",
					Vendor:    "nvidia",
					Name:      "1080ti",
					DeviceIDs: []string{deviceIDs[0], deviceIDs[1]},
				}),
				tests.CreateAllocWithDevice(allocIDs[1], lowPrioJob2, &structs.Resources{
					CPU:      200,
					MemoryMB: 100,
					DiskMB:   4 * 1024,
				}, &structs.AllocatedDeviceResource{
					Type:      "gpu",
					Vendor:    "nvidia",
					Name:      "1080ti",
					DeviceIDs: []string{deviceIDs[2], deviceIDs[3]},
				}),
				tests.CreateAllocWithDevice(allocIDs[2], lowPrioJob, &structs.Resources{
					CPU:      200,
					MemoryMB: 256,
					DiskMB:   4 * 1024,
				}, &structs.AllocatedDeviceResource{
					Type:      "gpu",
					Vendor:    "nvidia",
					Name:      "2080ti",
					DeviceIDs: []string{deviceIDs[4], deviceIDs[5]},
				}),
				tests.CreateAllocWithDevice(allocIDs[3], lowPrioJob, &structs.Resources{
					CPU:      100,
					MemoryMB: 256,
					DiskMB:   4 * 1024,
				}, &structs.AllocatedDeviceResource{
					Type:      "gpu",
					Vendor:    "nvidia",
					Name:      "2080ti",
					DeviceIDs: []string{deviceIDs[6], deviceIDs[7]},
				}),
				tests.CreateAllocWithDevice(allocIDs[4], lowPrioJob, &structs.Resources{
					CPU:      100,
					MemoryMB: 256,
					DiskMB:   4 * 1024,
				}, &structs.AllocatedDeviceResource{
					Type:      "gpu",
					Vendor:    "nvidia",
					Name:      "2080ti",
					DeviceIDs: []string{deviceIDs[8]},
				}),
				tests.CreateAllocWithDevice(allocIDs[5], lowPrioJob, &structs.Resources{
					CPU:      200,
					MemoryMB: 512,
					DiskMB:   4 * 1024,
				}, &structs.AllocatedDeviceResource{
					Type:      "fpga",
					Vendor:    "intel",
					Name:      "F100",
					DeviceIDs: []string{"fpga1"},
				})},
			nodeReservedCapacity: reservedNodeResources,
			nodeCapacity:         defaultNodeResources,
			jobPriority:          100,
			resourceAsk: &structs.Resources{
				CPU:      1000,
				MemoryMB: 512,
				DiskMB:   4 * 1024,
				Devices: []*structs.RequestedDevice{
					{
						Name:  "gpu",
						Count: 4,
					},
				},
			},
			preemptedAllocIDs: map[string]struct{}{
				allocIDs[2]: {},
				allocIDs[3]: {},
			},
		},
		{
			desc: "Device preemption not possible due to more instances needed than available",
			currentAllocations: []*structs.Allocation{
				tests.CreateAllocWithDevice(allocIDs[0], lowPrioJob, &structs.Resources{
					CPU:      500,
					MemoryMB: 512,
					DiskMB:   4 * 1024,
				}, &structs.AllocatedDeviceResource{
					Type:      "gpu",
					Vendor:    "nvidia",
					Name:      "1080ti",
					DeviceIDs: []string{deviceIDs[0], deviceIDs[1], deviceIDs[2], deviceIDs[3]},
				}),
				tests.CreateAllocWithDevice(allocIDs[1], lowPrioJob, &structs.Resources{
					CPU:      200,
					MemoryMB: 512,
					DiskMB:   4 * 1024,
				}, &structs.AllocatedDeviceResource{
					Type:      "fpga",
					Vendor:    "intel",
					Name:      "F100",
					DeviceIDs: []string{"fpga1"},
				})},
			nodeReservedCapacity: reservedNodeResources,
			nodeCapacity:         defaultNodeResources,
			jobPriority:          100,
			resourceAsk: &structs.Resources{
				CPU:      1000,
				MemoryMB: 512,
				DiskMB:   4 * 1024,
				Devices: []*structs.RequestedDevice{
					{
						Name:  "gpu",
						Count: 6,
					},
				},
			},
		},
		// This test case exercises the code path for a final filtering step that tries to
		// minimize the number of preemptible allocations
		{
			desc: "Filter out allocs whose resource usage superset is also in the preemption list",
			currentAllocations: []*structs.Allocation{
				tests.CreateAlloc(allocIDs[0], highPrioJob, &structs.Resources{
					CPU:      1800,
					MemoryMB: 2256,
					DiskMB:   4 * 1024,
					Networks: []*structs.NetworkResource{
						{
							Device: "eth0",
							IP:     "192.168.0.100",
							MBits:  150,
						},
					},
				}),
				tests.CreateAlloc(allocIDs[1], lowPrioJob, &structs.Resources{
					CPU:      1500,
					MemoryMB: 256,
					DiskMB:   5 * 1024,
					Networks: []*structs.NetworkResource{
						{
							Device: "eth0",
							IP:     "192.168.0.100",
							MBits:  100,
						},
					},
				}),
				tests.CreateAlloc(allocIDs[2], lowPrioJob, &structs.Resources{
					CPU:      600,
					MemoryMB: 256,
					DiskMB:   5 * 1024,
					Networks: []*structs.NetworkResource{
						{
							Device: "eth0",
							IP:     "192.168.0.200",
							MBits:  300,
						},
					},
				}),
			},
			nodeReservedCapacity: reservedNodeResources,
			nodeCapacity:         defaultNodeResources,
			jobPriority:          100,
			resourceAsk: &structs.Resources{
				CPU:      1000,
				MemoryMB: 256,
				DiskMB:   5 * 1024,
				Networks: []*structs.NetworkResource{
					{
						Device: "eth0",
						IP:     "192.168.0.100",
						MBits:  50,
					},
				},
			},
			preemptedAllocIDs: map[string]struct{}{
				allocIDs[1]: {},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			node := mock.Node()
			node.NodeResources = tc.nodeCapacity
			node.ReservedResources = tc.nodeReservedCapacity

			state, ctx := MockContext(t)

			nodes := []*RankedNode{
				{
					Node: node,
				},
			}
			state.UpsertNode(structs.MsgTypeTestSetup, 1000, node)
			for _, alloc := range tc.currentAllocations {
				alloc.NodeID = node.ID
			}
			err := state.UpsertAllocs(structs.MsgTypeTestSetup, 1001, tc.currentAllocations)

			must.NoError(t, err)
			if tc.currentPreemptions != nil {
				ctx.plan.NodePreemptions[node.ID] = tc.currentPreemptions
			}
			static := NewStaticRankIterator(ctx, nodes)
			binPackIter := NewBinPackIterator(ctx, static, true, tc.jobPriority)
			job := mock.Job()
			job.Priority = tc.jobPriority
			binPackIter.SetJob(job)
			binPackIter.SetSchedulerConfiguration(testSchedulerConfig)

			taskGroup := &structs.TaskGroup{
				EphemeralDisk: &structs.EphemeralDisk{},
				Tasks: []*structs.Task{
					{
						Name:      "web",
						Resources: tc.resourceAsk,
					},
				},
			}

			binPackIter.SetTaskGroup(taskGroup)
			option := binPackIter.Next()
			if tc.preemptedAllocIDs == nil {
				must.Nil(t, option)
			} else {
				must.NotNil(t, option)
				preemptedAllocs := option.PreemptedAllocs
				must.Eq(t, len(tc.preemptedAllocIDs), len(preemptedAllocs))
				for _, alloc := range preemptedAllocs {
					_, ok := tc.preemptedAllocIDs[alloc.ID]
					must.True(t, ok, must.Sprintf("alloc %s was preempted unexpectedly", alloc.ID))
				}
			}
		})
	}
}
