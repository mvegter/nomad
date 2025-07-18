## 1.10.3 (July 08, 2025)

IMPROVEMENTS:

* consul: Added kind field to service block for Consul service registrations [[GH-26170](https://github.com/hashicorp/nomad/issues/26170)]
* docker: Added support for cgroup namespaces in the task config [[GH-25927](https://github.com/hashicorp/nomad/issues/25927)]
* task environment: new NOMAD_UNIX_ADDR env var points to the task API unix socket, for use with workload identity [[GH-25598](https://github.com/hashicorp/nomad/issues/25598)]

BUG FIXES:

* agent: Fixed a bug to prevent a possible panic during graceful shutdown [[GH-26018](https://github.com/hashicorp/nomad/issues/26018)]
* agent: Fixed a bug to prevent panic during graceful server shutdown [[GH-26171](https://github.com/hashicorp/nomad/issues/26171)]
* agent: Fixed bug where agent would exit early from graceful shutdown when managed by systemd [[GH-26023](https://github.com/hashicorp/nomad/issues/26023)]
* cli: Fix panic when restarting stopped job with no scaling policies [[GH-26131](https://github.com/hashicorp/nomad/issues/26131)]
* cli: Fixed a bug in the `tls cert create` command that always added ``"<role>.global.nomad"` to the certificate DNS names, even when the specified region was not ``"global"`. [[GH-26086](https://github.com/hashicorp/nomad/issues/26086)]
* cli: Fixed a bug where the `acl token self` command only performed lookups for tokens set as environment variables and not by the `-token` flag. [[GH-26183](https://github.com/hashicorp/nomad/issues/26183)]
* client: Attempt to rollback directory creation when the `mkdir` plugin fails to perform ownership changes on it [[GH-26194](https://github.com/hashicorp/nomad/issues/26194)]
* client: Fixed bug where drained batch jobs would not be rescheduled if no eligible nodes were immediately available [[GH-26025](https://github.com/hashicorp/nomad/issues/26025)]
* docker: Fixed a bug where very low resources.cpu values could generate invalid cpu weights on hosts with very large client.cpu_total_compute values [[GH-26081](https://github.com/hashicorp/nomad/issues/26081)]
* host volumes: Fixed a bug where volumes with server-terminal allocations could be deleted from clients but not the state store [[GH-26213](https://github.com/hashicorp/nomad/issues/26213)]
* tls: Fixed a bug where reloading the Nomad server process with an updated `tls.verify_server_hostname` configuration parameter would not apply an update to internal RPC handler verification and require a full server restart [[GH-26107](https://github.com/hashicorp/nomad/issues/26107)]
* vault: Fixed a bug where non-periodic tokens would not have their TTL incremented to the lease duration [[GH-26041](https://github.com/hashicorp/nomad/issues/26041)]

## 1.10.2 (June 09, 2025)

BREAKING CHANGES:

* template: Support for the following non-hermetic sprig functions has been removed: sprig_date, sprig_dateInZone, sprig_dateModify, sprig_htmlDate, sprig_htmlDateInZone, sprig_dateInZone, sprig_dateModify, sprig_randAlphaNum, sprig_randAlpha, sprig_randAscii, sprig_randNumeric, sprig_randBytes, sprig_uuidv4, sprig_env, sprig_expandenv, and sprig_getHostByName. [[GH-25998](https://github.com/hashicorp/nomad/issues/25998)]

SECURITY:

* identity: Fixed bug where workflow identity policies are matched by job ID prefix (CVE-2025-4922) [[GH-25869](https://github.com/hashicorp/nomad/issues/25869)]
* template: Bump the consul-template version to resolve CVE-2025-27144, CVE-2025-22869, CVE-2025-22870 and CVE-2025-22872. [[GH-25998](https://github.com/hashicorp/nomad/issues/25998)]
* template: Removed support to the non-hermetic sprig_env, sprig_expandenv, and sprig_getHostByName sprig functions to prevent potential leakage of environment or network information, since they can allow reading environment variables or resolving domain names to IP addresses. [[GH-25998](https://github.com/hashicorp/nomad/issues/25998)]

IMPROVEMENTS:

* cli: Added job start command to allow starting a stopped job from the cli [[GH-24150](https://github.com/hashicorp/nomad/issues/24150)]
* client: Add gc_volumes_on_node_gc configuration to delete host volumes when nodes are garbage collected [[GH-25903](https://github.com/hashicorp/nomad/issues/25903)]
* client: add ability to set maximum allocation count by adding node_max_allocs to client configuration [[GH-25785](https://github.com/hashicorp/nomad/issues/25785)]
* host volumes: Add -force flag to volume delete command for removing volumes from GC'd nodes [[GH-25902](https://github.com/hashicorp/nomad/issues/25902)]
* identity: Allow ACL policies to be applied to a namespace [[GH-25871](https://github.com/hashicorp/nomad/issues/25871)]
* ipv6: bind and advertise addresses are now made to adhere to RFC-5942 §4 (reference: https://www.rfc-editor.org/rfc/rfc5952.html#section-4) [[GH-25921](https://github.com/hashicorp/nomad/issues/25921)]
* reporting (Enterprise): Added support for offline utilization reporting [[GH-25844](https://github.com/hashicorp/nomad/issues/25844)]
* template: adds ability to specify once mode for job templates [[GH-25922](https://github.com/hashicorp/nomad/issues/25922)]
* wi: new API endpoint for listing workload-attached ACL policies [[GH-25588](https://github.com/hashicorp/nomad/issues/25588)]

BUG FIXES:

* api: Fixed pagination bug which could result in duplicate results [[GH-25792](https://github.com/hashicorp/nomad/issues/25792)]
* client: Fixed a bug where disconnect.stop_on_client_after timeouts were extended or ignored [[GH-25946](https://github.com/hashicorp/nomad/issues/25946)]
* csi: Fixed -secret values not being sent with the `nomad volume snapshot delete` command [[GH-26022](https://github.com/hashicorp/nomad/issues/26022)]
* disconnect: Fixed a bug where pending evals for reconnected allocs were not cancelled [[GH-25923](https://github.com/hashicorp/nomad/issues/25923)]
* driver: Allow resources.cpu values above the maximum cpu.share value on Linux [[GH-25963](https://github.com/hashicorp/nomad/issues/25963)]
* job: Ensure sidecar task volume_mounts are added to planning diff object [[GH-25878](https://github.com/hashicorp/nomad/issues/25878)]
* reconnecting client: fix issue where reconcile strategy was sometimes ignored [[GH-25799](https://github.com/hashicorp/nomad/issues/25799)]
* scaling: Set the scaling policies to disabled when a job is stopped [[GH-25911](https://github.com/hashicorp/nomad/issues/25911)]
* scheduler: Fixed a bug where a node with no affinity could be selected over a node with low affinity [[GH-25800](https://github.com/hashicorp/nomad/issues/25800)]
* scheduler: Fixed a bug where planning or running a system job with constraints & previously running allocations would return a failed allocation error [[GH-25850](https://github.com/hashicorp/nomad/issues/25850)]
* telemetry: Fix excess CPU consumption from alloc stats collection [[GH-25870](https://github.com/hashicorp/nomad/issues/25870)]
* telemetry: Fixed a bug where alloc stats were still collected (but not published) if telemetry.publish_allocation_metrics=false. [[GH-25870](https://github.com/hashicorp/nomad/issues/25870)]
* ui: Fix incorrect calculation of permissions when ACLs are disabled which meant actions such as client drains were incorrectly blocked [[GH-25881](https://github.com/hashicorp/nomad/issues/25881)]

## 1.10.1 (May 13, 2025)

BREAKING CHANGES:

* api: The non-functional option -peer-address has been removed from the operator raft remove-peer command and equivalent API [[GH-25599](https://github.com/hashicorp/nomad/issues/25599)]
* core: Errors encountered when reloading agent configuration will now cause agents to exit. Before configuration errors during reloads were only logged. This could lead to agents running but unable to communicate [[GH-25721](https://github.com/hashicorp/nomad/issues/25721)]

SECURITY:

* build: Update Go to 1.24.3 to address CVE-2025-22873 [[GH-25818](https://github.com/hashicorp/nomad/issues/25818)]
* sentinel (Enterprise): Fixed a bug where in some cases hard-mandatory policies could be overridden with -policy-override. CVE-2025-3744 [[GH-2618](https://github.com/hashicorp/nomad-enterprise/pull/2618)]

IMPROVEMENTS:

* command: added priority flag to job dispatch command [[GH-25622](https://github.com/hashicorp/nomad/issues/25622)]

BUG FIXES:

* agent: Fixed a bug where reloading the agent with systemd notification enabled would cause the agent to be killed by system [[GH-25636](https://github.com/hashicorp/nomad/issues/25636)]
* cli: Respect NOMAD_REGION environment variable in operator debug command [[GH-25716](https://github.com/hashicorp/nomad/issues/25716)]
* client: fix failure cleaning up namespace on batch jobs [[GH-25714](https://github.com/hashicorp/nomad/issues/25714)]
* docker: Fix missing stats for rss, cache and swap memory for cgroups v1 [[GH-25741](https://github.com/hashicorp/nomad/issues/25741)]
* encrypter: Refactor startup decryption task handling to avoid timing problems with task addition on FSM restore [[GH-25795](https://github.com/hashicorp/nomad/issues/25795)]
* java: Fixed a bug where the default task user was set to 'nobody' on Windows [[GH-25648](https://github.com/hashicorp/nomad/issues/25648)]
* metrics: Fixed a bug where RSS and cache stats would not be reported for docker, exec, and java drivers under Linux cgroups v2 [[GH-25751](https://github.com/hashicorp/nomad/issues/25751)]
* scheduler: Fixed a bug in accounting for resources.cores that could prevent placements on nodes with available cores [[GH-25705](https://github.com/hashicorp/nomad/issues/25705)]
* scheduler: Fixed a bug where draining a node with canaries could result in a stuck deployment [[GH-25726](https://github.com/hashicorp/nomad/issues/25726)]
* scheduler: Fixed a bug where updating the rescheduler tracker could corrupt the state store [[GH-25698](https://github.com/hashicorp/nomad/issues/25698)]
* scheduler: Use core ID when selecting cores. This fixes a panic in the scheduler when the `reservable_cores` is not a contiguous list of core IDs. [[GH-25340](https://github.com/hashicorp/nomad/issues/25340)]
* server: Added a new server configuration option named `start_timeout` with a default value of `30s`. This duration is used to monitor the server setup and startup processes which must complete before it is considered healthy, such as keyring decryption. If these processes do not complete before the timeout is reached, the server process will exit. [[GH-25803](https://github.com/hashicorp/nomad/issues/25803)]
* ui: Fixed a bug where the job list page incorrectly calculated if a job had paused tasks. [[GH-25742](https://github.com/hashicorp/nomad/issues/25742)]

## 1.10.0 (April 09, 2025)

FEATURES:

* **Dynamic Host Volumes:** Nomad now supports creating host volumes via the API [[GH-24479](https://github.com/hashicorp/nomad/issues/24479)]
* **OIDC Login:** Nomad now enables PKCE for OIDC logins, and supports the private key JWT / client assertion option in the OIDC authentication flow. [[GH-25231](https://github.com/hashicorp/nomad/issues/25231)]
* **Stateful Deployments:** Nomad now supports stateful deployments when using dynamic host volumes. [[GH-24993](https://github.com/hashicorp/nomad/issues/24993)]

BREAKING CHANGES:

* agent: Plugins stored within the `plugin_dir` will now only be executed when they have a corresponding `plugin` configuration block. Any plugin found without a corresponding configuration block will be skipped. [[GH-18530](https://github.com/hashicorp/nomad/issues/18530)]
* api: QuotaSpec.RegionLimit is now of type QuotaResources instead of Resources [[GH-24785](https://github.com/hashicorp/nomad/issues/24785)]
* consul: Identities are no longer added to tasks by default when they include a template block.
Please see [Nomad's upgrade guide](https://developer.hashicorp.com/nomad/docs/upgrade/upgrade-specific)
for more detail. [[GH-25298](https://github.com/hashicorp/nomad/issues/25298)]
* consul: The deprecated token-based authentication workflow for allocations has been removed. Please see [Nomad's upgrade guide](https://developer.hashicorp.com/nomad/docs/upgrade/upgrade-specific) for more detail. [[GH-25217](https://github.com/hashicorp/nomad/issues/25217)]
* disconnected nodes: ignore the previously deprecated disconnect group fields in favor of the disconnect block introduced in Nomad 1.8 [[GH-25284](https://github.com/hashicorp/nomad/issues/25284)]
* drivers: remove remote task support for task drivers [[GH-24909](https://github.com/hashicorp/nomad/issues/24909)]
* sentinel: The sentinel apply command now requires the -scope option [[GH-24601](https://github.com/hashicorp/nomad/issues/24601)]
* vault: The deprecated token-based authentication workflow for allocations has been removed. Please
see [Nomad's upgrade guide](https://developer.hashicorp.com/nomad/docs/upgrade/upgrade-specific) for
more detail. [[GH-25155](https://github.com/hashicorp/nomad/issues/25155)]

IMPROVEMENTS:

* build: Updated Go to 1.24.2 [[GH-25623](https://github.com/hashicorp/nomad/issues/25623)]
* cli: Add -group option to `alloc exec`, `alloc logs`, `alloc fs` commands [[GH-25568](https://github.com/hashicorp/nomad/issues/25568)]
* cli: Added UI URL hints to the end of common CLI commands and a `-ui` flag to auto-open them [[GH-24454](https://github.com/hashicorp/nomad/issues/24454)]
* client: Fixed a bug where JSON formatted logs would not show the requested and overlapping cores when failing to reserve cores [[GH-25523](https://github.com/hashicorp/nomad/issues/25523)]
* client: Improve memory usage by dropping references to task environment [[GH-25373](https://github.com/hashicorp/nomad/issues/25373)]
* cni: Add a warning log when CNI check commands fail [[GH-25581](https://github.com/hashicorp/nomad/issues/25581)]
* csi: Accept ID prefixes and wildcard namespace for the volume delete command [[GH-24997](https://github.com/hashicorp/nomad/issues/24997)]
* csi: Added CSI volume and plugin events to the event stream [[GH-24724](https://github.com/hashicorp/nomad/issues/24724)]
* csi: Show volume capabilities in the volume status command [[GH-25173](https://github.com/hashicorp/nomad/issues/25173)]
* drivers/docker: adds image_pull_timeout to plugin config options [[GH-25489](https://github.com/hashicorp/nomad/issues/25489)]
* drivers/rawexec: adds denied_envvars to driver and task config options [[GH-25511](https://github.com/hashicorp/nomad/issues/25511)]
* rawexec: add support for setting the task user on windows platform [[GH-25496](https://github.com/hashicorp/nomad/issues/25496)]
* rpc: Added ability to configure yamux session parameters [[GH-25466](https://github.com/hashicorp/nomad/issues/25466)]
* ui: Added Dynamic Host Volumes to the web UI [[GH-25224](https://github.com/hashicorp/nomad/issues/25224)]
* ui: Added a scope selector for sentinel policy page [[GH-25390](https://github.com/hashicorp/nomad/issues/25390)]
* ui: Makes jobs list filtering case-insensitive [[GH-25378](https://github.com/hashicorp/nomad/issues/25378)]
* ui: Updated icons to the newest design system [[GH-25353](https://github.com/hashicorp/nomad/issues/25353)]

DEPRECATIONS:

* api: QuotaSpec.VariablesLimit field is deprecated and will be removed in Nomad 1.12.0. Use QuotaSpec.RegionLimit.Storage.Variables instead. [[GH-24785](https://github.com/hashicorp/nomad/issues/24785)]
* quotas: the variables_limit field in the quota specification is deprecated and replaced by a new storage block under the region_limit block, with a variables field. The variables_limit field will be removed in Nomad 1.12.0 [[GH-24785](https://github.com/hashicorp/nomad/issues/24785)]

BUG FIXES:

* client: fixed a bug where AMD CPUs were not correctly fingerprinting base speed [[GH-24415](https://github.com/hashicorp/nomad/issues/24415)]
* client: remove blocking call during client gc [[GH-25123](https://github.com/hashicorp/nomad/issues/25123)]
* client: skip a task groups shutdown_delay when all tasks have already been deregistered [[GH-25157](https://github.com/hashicorp/nomad/issues/25157)]
* csi: Fixed a CSI ExpandVolume bug where the namespace was left out of the staging path [[GH-25253](https://github.com/hashicorp/nomad/issues/25253)]
* csi: Fixed a bug where GC would attempt and fail to delete plugins that had volumes [[GH-25432](https://github.com/hashicorp/nomad/issues/25432)]
* csi: Fixed a bug where cleaning up volume claims on GC'd nodes would cause errors on the leader [[GH-25428](https://github.com/hashicorp/nomad/issues/25428)]
* csi: Fixed a bug where in-flight CSI RPCs would not be cancelled on client GC or dev agent shutdown [[GH-25472](https://github.com/hashicorp/nomad/issues/25472)]
* drivers: set -1 exit code in case of executor failure for the exec, raw_exec, java, and qemu task drivers [[GH-25453](https://github.com/hashicorp/nomad/issues/25453)]
* job: Ensure migrate block difference is added to planning diff object [[GH-25528](https://github.com/hashicorp/nomad/issues/25528)]
* scheduler: Fixed a bug that made affinity and spread updates destructive [[GH-25109](https://github.com/hashicorp/nomad/issues/25109)]
* server: Validate `num_schedulers` configuration parameter is between 0 and the number of CPUs available on the machine [[GH-25441](https://github.com/hashicorp/nomad/issues/25441)]
* services: Fixed a bug where Nomad native services would not be correctly interpolated during in-place updates [[GH-25373](https://github.com/hashicorp/nomad/issues/25373)]
* services: Fixed a bug where task-level services, checks, and identities could interpolate jobspec values from other tasks in the same group [[GH-25373](https://github.com/hashicorp/nomad/issues/25373)]

## 1.9.11 Enterprise (July 08, 2025)

BUG FIXES:

* agent: Fixed a bug to prevent a possible panic during graceful shutdown [[GH-26018](https://github.com/hashicorp/nomad/issues/26018)]
* agent: Fixed a bug to prevent panic during graceful server shutdown [[GH-26171](https://github.com/hashicorp/nomad/issues/26171)]
* agent: Fixed bug where agent would exit early from graceful shutdown when managed by systemd [[GH-26023](https://github.com/hashicorp/nomad/issues/26023)]
* cli: Fixed a bug in the `tls cert create` command that always added ``"<role>.global.nomad"` to the certificate DNS names, even when the specified region was not ``"global"`. [[GH-26086](https://github.com/hashicorp/nomad/issues/26086)]
* client: Fixed bug where drained batch jobs would not be rescheduled if no eligible nodes were immediately available [[GH-26025](https://github.com/hashicorp/nomad/issues/26025)]
* docker: Fixed a bug where very low resources.cpu values could generate invalid cpu weights on hosts with very large client.cpu_total_compute values [[GH-26081](https://github.com/hashicorp/nomad/issues/26081)]
* tls: Fixed a bug where reloading the Nomad server process with an updated `tls.verify_server_hostname` configuration parameter would not apply an update to internal RPC handler verification and require a full server restart [[GH-26107](https://github.com/hashicorp/nomad/issues/26107)]
* vault: Fixed a bug where non-periodic tokens would not have their TTL incremented to the lease duration [[GH-26041](https://github.com/hashicorp/nomad/issues/26041)]

## 1.9.10 Enterprise (June 10, 2025)

BREAKING CHANGES:

* template: Support for the following non-hermetic sprig functions has been removed: sprig_date, sprig_dateInZone, sprig_dateModify, sprig_htmlDate, sprig_htmlDateInZone, sprig_dateInZone, sprig_dateModify, sprig_randAlphaNum, sprig_randAlpha, sprig_randAscii, sprig_randNumeric, sprig_randBytes, sprig_uuidv4, sprig_env, sprig_expandenv, and sprig_getHostByName. [[GH-25998](https://github.com/hashicorp/nomad/issues/25998)]

SECURITY:

* identity: Fixed bug where workflow identity policies are matched by job ID prefix (CVE-2025-4922) [[GH-25869](https://github.com/hashicorp/nomad/issues/25869)]
* template: Bump the consul-template version to resolve CVE-2025-27144, CVE-2025-22869, CVE-2025-22870 and CVE-2025-22872. [[GH-25998](https://github.com/hashicorp/nomad/issues/25998)]
* template: Removed support to the non-hermetic sprig_env, sprig_expandenv, and sprig_getHostByName sprig functions to prevent potential leakage of environment or network information, since they can allow reading environment variables or resolving domain names to IP addresses. [[GH-25998](https://github.com/hashicorp/nomad/issues/25998)]

IMPROVEMENTS:

* reporting (Enterprise): Added support for offline utilization reporting [[GH-25844](https://github.com/hashicorp/nomad/issues/25844)]

BUG FIXES:

* client: Fixed a bug where disconnect.stop_on_client_after timeouts were extended or ignored [[GH-25946](https://github.com/hashicorp/nomad/issues/25946)]
* csi: Fixed -secret values not being sent with the `nomad volume snapshot delete` command [[GH-26022](https://github.com/hashicorp/nomad/issues/26022)]
* disconnect: Fixed a bug where pending evals for reconnected allocs were not cancelled [[GH-25923](https://github.com/hashicorp/nomad/issues/25923)]
* driver: Allow resources.cpu values above the maximum cpu.share value on Linux [[GH-25963](https://github.com/hashicorp/nomad/issues/25963)]
* job: Ensure sidecar task volume_mounts are added to planning diff object [[GH-25878](https://github.com/hashicorp/nomad/issues/25878)]
* reconnecting client: fix issue where reconcile strategy was sometimes ignored [[GH-25799](https://github.com/hashicorp/nomad/issues/25799)]
* scaling: Set the scaling policies to disabled when a job is stopped [[GH-25911](https://github.com/hashicorp/nomad/issues/25911)]
* scheduler: Fixed a bug where a node with no affinity could be selected over a node with low affinity [[GH-25800](https://github.com/hashicorp/nomad/issues/25800)]
* scheduler: Fixed a bug where planning or running a system job with constraints & previously running allocations would return a failed allocation error [[GH-25850](https://github.com/hashicorp/nomad/issues/25850)]
* telemetry: Fix excess CPU consumption from alloc stats collection [[GH-25870](https://github.com/hashicorp/nomad/issues/25870)]
* telemetry: Fixed a bug where alloc stats were still collected (but not published) if telemetry.publish_allocation_metrics=false. [[GH-25870](https://github.com/hashicorp/nomad/issues/25870)]
* vault: Fixed a bug where poststop tasks could not obtain Vault tokens after the main task failed

## 1.9.9 Enterprise (May 13, 2025)

BREAKING CHANGES:

* core: Errors encountered when reloading agent configuration will now cause agents to exit. Before configuration errors during reloads were only logged. This could lead to agents running but unable to communicate [[GH-25721](https://github.com/hashicorp/nomad/issues/25721)]

SECURITY:

* build: Update Go to 1.24.3 to address CVE-2025-22873 [[GH-25818](https://github.com/hashicorp/nomad/issues/25818)]
* sentinel (Enterprise): Fixed a bug where in some cases hard-mandatory policies could be overridden with -policy-override. CVE-2025-3744.

BUG FIXES:

* agent: Fixed a bug where reloading the agent with systemd notification enabled would cause the agent to be killed by system [[GH-25636](https://github.com/hashicorp/nomad/issues/25636)]
* api: Fixed pagination bug which could result in duplicate results [[GH-25792](https://github.com/hashicorp/nomad/issues/25792)]
* cli: Respect NOMAD_REGION environment variable in operator debug command [[GH-25716](https://github.com/hashicorp/nomad/issues/25716)]
* client: fix failure cleaning up namespace on batch jobs [[GH-25714](https://github.com/hashicorp/nomad/issues/25714)]
* docker: Fix missing stats for rss, cache and swap memory for cgroups v1 [[GH-25741](https://github.com/hashicorp/nomad/issues/25741)]
* encrypter: Refactor startup decryption task handling to avoid timing problems with task addition on FSM restore [[GH-25795](https://github.com/hashicorp/nomad/issues/25795)]
* metrics: Fixed a bug where RSS and cache stats would not be reported for docker, exec, and java drivers under Linux cgroups v2 [[GH-25751](https://github.com/hashicorp/nomad/issues/25751)]
* scheduler: Fixed a bug in accounting for resources.cores that could prevent placements on nodes with available cores [[GH-25705](https://github.com/hashicorp/nomad/issues/25705)]
* scheduler: Fixed a bug where draining a node with canaries could result in a stuck deployment [[GH-25726](https://github.com/hashicorp/nomad/issues/25726)]
* scheduler: Fixed a bug where updating the rescheduler tracker could corrupt the state store [[GH-25698](https://github.com/hashicorp/nomad/issues/25698)]
* scheduler: Use core ID when selecting cores. This fixes a panic in the scheduler when the `reservable_cores` is not a contiguous list of core IDs. [[GH-25340](https://github.com/hashicorp/nomad/issues/25340)]
* server: Added a new server configuration option named `start_timeout` with a default value of `30s`. This duration is used to monitor the server setup and startup processes which must complete before it is considered healthy, such as keyring decryption. If these processes do not complete before the timeout is reached, the server process will exit. [[GH-25803](https://github.com/hashicorp/nomad/issues/25803)]
* ui: Fixed a bug where the job list page incorrectly calculated if a job had paused tasks. [[GH-25742](https://github.com/hashicorp/nomad/issues/25742)]

## 1.9.8 Enterprise (April 9, 2025)

IMPROVEMENTS:

* build: Updated Go to 1.24.2 [[GH-25623](https://github.com/hashicorp/nomad/issues/25623)]
* client: Improve memory usage by dropping references to task environment [[GH-25373](https://github.com/hashicorp/nomad/issues/25373)]
* cni: Add a warning log when CNI check commands fail [[GH-25581](https://github.com/hashicorp/nomad/issues/25581)]
* ui: Makes jobs list filtering case-insensitive [[GH-25378](https://github.com/hashicorp/nomad/issues/25378)]

BUG FIXES:

* client: remove blocking call during client gc [[GH-25123](https://github.com/hashicorp/nomad/issues/25123)]
* client: skip a task groups shutdown_delay when all tasks have already been deregistered [[GH-25157](https://github.com/hashicorp/nomad/issues/25157)]
* csi: Fixed a CSI ExpandVolume bug where the namespace was left out of the staging path [[GH-25253](https://github.com/hashicorp/nomad/issues/25253)]
* csi: Fixed a bug where GC would attempt and fail to delete plugins that had volumes [[GH-25432](https://github.com/hashicorp/nomad/issues/25432)]
* csi: Fixed a bug where cleaning up volume claims on GC'd nodes would cause errors on the leader [[GH-25428](https://github.com/hashicorp/nomad/issues/25428)]
* csi: Fixed a bug where in-flight CSI RPCs would not be cancelled on client GC or dev agent shutdown [[GH-25472](https://github.com/hashicorp/nomad/issues/25472)]
* drivers: set -1 exit code in case of executor failure for the exec, raw_exec, java, and qemu task drivers [[GH-25453](https://github.com/hashicorp/nomad/issues/25453)]
* job: Ensure migrate block difference is added to planning diff object [[GH-25528](https://github.com/hashicorp/nomad/issues/25528)]
* server: Validate `num_schedulers` configuration parameter is between 0 and the number of CPUs available on the machine [[GH-25441](https://github.com/hashicorp/nomad/issues/25441)]
* services: Fixed a bug where Nomad native services would not be correctly interpolated during in-place updates [[GH-25373](https://github.com/hashicorp/nomad/issues/25373)]
* services: Fixed a bug where task-level services, checks, and identities could interpolate jobspec values from other tasks in the same group [[GH-25373](https://github.com/hashicorp/nomad/issues/25373)]

## 1.9.7 (March 11, 2025)

BREAKING CHANGES:

* node: The node attribute `consul.addr.dns` has been changed to `unique.consul.addr.dns`. The node attribute `nomad.advertise.address` has been changed to `unique.advertise.address`. [[GH-24942](https://github.com/hashicorp/nomad/issues/24942)]

SECURITY:

* auth: Redact OIDC client secret from API responses and event stream ([CVE-2025-1296](https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2025-1296)) [[GH-25328](https://github.com/hashicorp/nomad/issues/25328)]

IMPROVEMENTS:

* build: Updated Go to 1.24.1 [[GH-25249](https://github.com/hashicorp/nomad/issues/25249)]
* config: Allow disabling `wait` in client config [[GH-25255](https://github.com/hashicorp/nomad/issues/25255)]
* cpustats: Add config "cpu_disable_dmidecode" to disable cpu detection using dmidecode [[GH-25108](https://github.com/hashicorp/nomad/issues/25108)]
* metrics: Fix the process lookup for raw_exec when running rootless [[GH-25198](https://github.com/hashicorp/nomad/issues/25198)]
* ui: System, Batch and Sysbatch jobs get a "Revert to prev version" button on their main pages [[GH-25104](https://github.com/hashicorp/nomad/issues/25104)]

BUG FIXES:

* cli: Add node_prefix read when setting up the task workload identity Consul policy [[GH-25310](https://github.com/hashicorp/nomad/issues/25310)]
* cni: Fixed a bug where CNI state was not migrated after upgrade, resulting in IP collisions [[GH-25093](https://github.com/hashicorp/nomad/issues/25093)]
* csi: Fixed a bug where plugins that failed initial fingerprints would not be restarted [[GH-25307](https://github.com/hashicorp/nomad/issues/25307)]
* fingerprint: Fixed a bug where Consul/Vault would never be fingerprinted if not available on agent start [[GH-25102](https://github.com/hashicorp/nomad/issues/25102)]
* hcl: Avoid panics by checking null values on durations [[GH-25294](https://github.com/hashicorp/nomad/issues/25294)]
* rpc: Fixed a bug that would cause the reader side of RPC connections to hang indefinitely [[GH-25201](https://github.com/hashicorp/nomad/issues/25201)]
* scheduler: Fixed a bug where node class hashes included unique attributes, making scheduling more costly [[GH-24942](https://github.com/hashicorp/nomad/issues/24942)]
* template: Fixed a bug where unset client.template retry blocks ignored defaults [[GH-25113](https://github.com/hashicorp/nomad/issues/25113)]
* template: Updated the consul-template dependency to v0.40.0 which included a bug fix in the
quiescence timers. This bug could cause increased Nomad client CPU usage for tasks which use two or
more template blocks. [[GH-25140](https://github.com/hashicorp/nomad/issues/25140)]

## 1.9.6 (February 11, 2025)

BREAKING CHANGES:

* fingerprint: Consul and Vault fingerprints no longer reload periodically [[GH-24526](https://github.com/hashicorp/nomad/issues/24526)]

SECURITY:

* api: sanitize the SignedIdentities in allocations of events to clean the identity token. [[GH-24966](https://github.com/hashicorp/nomad/issues/24966)]
* build: Updated Go to 1.23.6 [[GH-25041](https://github.com/hashicorp/nomad/issues/25041)]
* event stream: fixes vulnerability CVE-2025-0937, where using a wildcard namespace to subscribe to the events API grants a user with "read" capabilites on any namespace, the ability to read events from all namespaces. [[GH-25089](https://github.com/hashicorp/nomad/issues/25089)]

IMPROVEMENTS:

* auth: adds `VerboseLogging` option to auth-method config for debugging SSO [[GH-24892](https://github.com/hashicorp/nomad/issues/24892)]
* cli: Added actions available to a job when running nomad job status command [[GH-24959](https://github.com/hashicorp/nomad/issues/24959)]
* event stream: adds ability to authenticate using workload identities [[GH-24849](https://github.com/hashicorp/nomad/issues/24849)]
* services: Nomad service checks now support the `tls_skip_verify` parameter [[GH-24781](https://github.com/hashicorp/nomad/issues/24781)]
* task schedule: The task being paused no longer impacts restart attempts [[GH-25085](https://github.com/hashicorp/nomad/issues/25085)]
* ui: Contextualizes the Start Job button on whether it is startable, revertable, or not [[GH-24985](https://github.com/hashicorp/nomad/issues/24985)]

BUG FIXES:

* agent: Fixed a bug where Nomad error log messages within syslog showed via the notice priority [[GH-24820](https://github.com/hashicorp/nomad/issues/24820)]
* agent: Fixed a bug where all syslog entries were marked as notice when using JSON logging format [[GH-24865](https://github.com/hashicorp/nomad/issues/24865)]
* client: Fixed a bug where temporary RPC errors cause the client to poll for changes more frequently thereafter [[GH-25039](https://github.com/hashicorp/nomad/issues/25039)]
* csi: Fixed a bug where volume context from the plugin would be erased on volume updates [[GH-24922](https://github.com/hashicorp/nomad/issues/24922)]
* docker: Fixed a bug that prevented image_pull_timeout from being applied [[GH-24991](https://github.com/hashicorp/nomad/issues/24991)]
* docker: Fixed a bug where "error reading image pull progress" caused the allocation to get stuck pending [[GH-24981](https://github.com/hashicorp/nomad/issues/24981)]
* reporting (Enterprise): Updated the reporting metric to utilize node active heartbeat count. [[GH-24919](https://github.com/hashicorp/nomad/issues/24919)]
* state store: fix for setting correct status for a job version when reverting, and also fixes an issue where jobs were briefly marked dead during restarts [[GH-24974](https://github.com/hashicorp/nomad/issues/24974)]
* taskrunner: fix panic when a task with dynamic user is recovered [[GH-24739](https://github.com/hashicorp/nomad/issues/24739)]
* ui: Ensure pending service check blocks are filled [[GH-24818](https://github.com/hashicorp/nomad/issues/24818)]
* ui: Remove unrequired node read API call when attempting to stream task logs [[GH-24973](https://github.com/hashicorp/nomad/issues/24973)]
* vault: Fixed a bug where successful renewal was logged as an error [[GH-25040](https://github.com/hashicorp/nomad/issues/25040)]

## 1.9.5 (January 14, 2025)

IMPROVEMENTS:

* client: Add noswap mount option to secrets directory where supported on Linux [[GH-24645](https://github.com/hashicorp/nomad/issues/24645)]
* deps: Upgraded aws-sdk-go from v1 to v2 [[GH-24720](https://github.com/hashicorp/nomad/issues/24720)]
* keyring: Warn if deleting a key previously used to encrypt an existing variable [[GH-24766](https://github.com/hashicorp/nomad/issues/24766)]
* ui: Added possibility to supply HCL variable values on job submission [[GH-24622](https://github.com/hashicorp/nomad/issues/24622)]
* ui: add leadership status for servers in other regions [[GH-24723](https://github.com/hashicorp/nomad/issues/24723)]

BUG FIXES:

* docker: Fix a bug where images with port number and no tags weren't parsed correctly [[GH-24547](https://github.com/hashicorp/nomad/issues/24547)]
* driver/docker: Fix container CPU stats collection where previous CPU stats were missing and causing incorrect calculations [[GH-24768](https://github.com/hashicorp/nomad/issues/24768)]
* drivers: validate logmon plugin during reattach [[GH-24798](https://github.com/hashicorp/nomad/issues/24798)]
* networking: check network namespaces on Linux during client restarts and fail the allocation if an existing namespace is invalid [[GH-24658](https://github.com/hashicorp/nomad/issues/24658)]

## 1.9.4 (December 18, 2024)

SECURITY:

* api: sanitize the SignedIdentities in allocations to prevent privilege escalation through unredacted workload identity token impersonation associated with ACL policies. [[GH-24683](https://github.com/hashicorp/nomad/issues/24683)]
* security: Added more host environment variables to the default deny list for tasks [[GH-24540](https://github.com/hashicorp/nomad/issues/24540)]
* security: Explicitly set 'Content-Type' header to mitigate XSS vulnerability [[GH-24489](https://github.com/hashicorp/nomad/issues/24489)]
* security: add executeTemplate to default template function_denylist [[GH-24541](https://github.com/hashicorp/nomad/issues/24541)]

IMPROVEMENTS:

* actions: Nomad Actions names now accept a wider range of names [[GH-24642](https://github.com/hashicorp/nomad/issues/24642)]
* api: Sanitise hcl variables before storage on JobSubmission [[GH-24423](https://github.com/hashicorp/nomad/issues/24423)]
* client: Emit telemetry from prerun and prestart hooks for monitoring and alerting [[GH-24556](https://github.com/hashicorp/nomad/issues/24556)]
* cni: Add Nomad specific workload information to CNI_ARGS [[GH-24319](https://github.com/hashicorp/nomad/issues/24319)]
* core: add the possibility to scale system jobs between 0 and 1 [[GH-24363](https://github.com/hashicorp/nomad/issues/24363)]
* ui: Add an Edit From Version button as an option when reverting from an older job version [[GH-24168](https://github.com/hashicorp/nomad/issues/24168)]
* ui: Adds metadata tables to Task Group and Task pages [[GH-24594](https://github.com/hashicorp/nomad/issues/24594)]

BUG FIXES:

* agent: Fixed a bug where `retry_join` gave up after a single failure, rather than retrying until max attempts had been reached [[GH-24561](https://github.com/hashicorp/nomad/issues/24561)]
* api: Fixed a bug where alloc exec/logs/fs APIs would return errors for non-global regions [[GH-24644](https://github.com/hashicorp/nomad/issues/24644)]
* cli: Ensure the `operator autopilot health` command only outputs JSON when the `json` flag is supplied [[GH-24655](https://github.com/hashicorp/nomad/issues/24655)]
* consul: Fixed a bug where failures when syncing Consul checks could panic the Nomad agent [[GH-24513](https://github.com/hashicorp/nomad/issues/24513)]
* consul: Fixed a bug where non-root Nomad agents could not recreate a task's Consul token on task restart [[GH-24410](https://github.com/hashicorp/nomad/issues/24410)]
* csi: Fixed a bug where drivers that emit multiple topology segments would cause placements to fail [[GH-24522](https://github.com/hashicorp/nomad/issues/24522)]
* csi: Removed redundant namespace output from volume status command [[GH-24432](https://github.com/hashicorp/nomad/issues/24432)]
* discovery: Fixed a bug where IPv6 addresses would not be accepted from cloud autojoin [[GH-24649](https://github.com/hashicorp/nomad/issues/24649)]
* drivers: fix executor leak when drivers error starting tasks [[GH-24495](https://github.com/hashicorp/nomad/issues/24495)]
* executor: validate executor on reattach to avoid possibility of killing non-Nomad processes [[GH-24538](https://github.com/hashicorp/nomad/issues/24538)]
* keyring: Fixed a bug when decrypting aead with an empty RSA block on state upserts [[GH-24442](https://github.com/hashicorp/nomad/issues/24442)]
* networking: use a tmpfs location for the state of CNI IPAM plugin used by bridge mode, to fix a bug where allocations would fail to restore after host reboot [[GH-24650](https://github.com/hashicorp/nomad/issues/24650)]
* scheduler: Fix bug where forced garbage collection does not ignore GC thresholds [[GH-24456](https://github.com/hashicorp/nomad/issues/24456)]
* scheduler: take all assigned cpu cores into account instead of only those part of the largest lifecycle [[GH-24304](https://github.com/hashicorp/nomad/issues/24304)]
* ui: Fix a bug where namespaced jobs wouldn't show diffs on the versions page [[GH-24466](https://github.com/hashicorp/nomad/issues/24466)]
* ui: Fix an issue where 2 parent jobs would see the others dispatches if it were otherwise empty [[GH-24668](https://github.com/hashicorp/nomad/issues/24668)]
* ui: Fix an issue where cmd+click or ctrl+click would double-open a var [[GH-24316](https://github.com/hashicorp/nomad/issues/24316)]
* ui: Fix an issue where system jobs with garbage-collected allocations were showing as Scaled Down [[GH-24620](https://github.com/hashicorp/nomad/issues/24620)]
* ui: Fix an issue where volumes weren't navigable [[GH-24542](https://github.com/hashicorp/nomad/issues/24542)]
* vault: Fixed a bug where expired secret leases were treated as non-fatal and retried [[GH-24409](https://github.com/hashicorp/nomad/issues/24409)]
* windows: Restore process accounting logic from Nomad 1.6.x [[GH-24494](https://github.com/hashicorp/nomad/issues/24494)]

## 1.9.3 (November 11, 2024)

BUG FIXES:

* fix: handles consul template re-renders on client restart [[GH-24399](https://github.com/hashicorp/nomad/issues/24399)]
* fsm: Fix bug in snapshot restore that could partially wipe state [[GH-24412](https://github.com/hashicorp/nomad/issues/24412)]

## 1.9.2 (November 08, 2024)

SECURITY:

* csi: Fixed a bug where a user with csi-write-volume permissions to one namespace can create volumes in another namespace (CVE-2024-10975) [[GH-24396](https://github.com/hashicorp/nomad/issues/24396)]

IMPROVEMENTS:

* api: new parameterized dispatch endpoint sends raw HTTP request body as Payload [[GH-24312](https://github.com/hashicorp/nomad/issues/24312)]
* connect: Able to accept go-sockaddr address for consul grpc address [[GH-24280](https://github.com/hashicorp/nomad/issues/24280)]
* consul: add support for service weight [[GH-24186](https://github.com/hashicorp/nomad/issues/24186)]
* drivers: Add work_dir config to exec/rawexec/java drivers for setting the working directory of processes in a task [[GH-24249](https://github.com/hashicorp/nomad/issues/24249)]
* drivers: Move executor process out of task cgroup after task starts on cgroups v1 [[GH-24340](https://github.com/hashicorp/nomad/issues/24340)]
* drivers: add posibility to restrict user and group for exec and rawexec [[GH-20073](https://github.com/hashicorp/nomad/issues/20073)]
* fingerprint gce: fingerprint preemptibility [[GH-24169](https://github.com/hashicorp/nomad/issues/24169)]
* getter: Added option to chown artifact(s) to task user [[GH-24157](https://github.com/hashicorp/nomad/issues/24157)]
* ui: Indicates prestart/poststart tasks by running/failed/pending status [[GH-24133](https://github.com/hashicorp/nomad/issues/24133)]
* ui: Show region in header when only one region exists, and set it immediately upon logging in with a token [[GH-24320](https://github.com/hashicorp/nomad/issues/24320)]
* ui: When your token expires, upon signing back in, redirect to your original route [[GH-24374](https://github.com/hashicorp/nomad/issues/24374)]

BUG FIXES:

* cli: Corrected an ordering mistake in job tag unset help text [[GH-24272](https://github.com/hashicorp/nomad/issues/24272)]
* connect: add validation to ensure that connect native services specify a port [[GH-24329](https://github.com/hashicorp/nomad/issues/24329)]
* deps: Fixed a bug where restarting Nomad could cause an unrelated process with the same PID as a failed executor to be killed [[GH-24265](https://github.com/hashicorp/nomad/issues/24265)]
* keyring: Fixed a panic on server startup when decrypting AEAD key data with empty RSA block [[GH-24383](https://github.com/hashicorp/nomad/issues/24383)]
* scheduler: fixed a bug where resource calculation did not account correctly for poststart tasks [[GH-24297](https://github.com/hashicorp/nomad/issues/24297)]
* state: Fixed setting GC threshold to more than 72hrs being ignored [[GH-24112](https://github.com/hashicorp/nomad/issues/24112)]

## 1.9.1 (October 21, 2024)

IMPROVEMENTS:

* cli: Added synopsis for `operator root` and `operator gossip` command [[GH-23671](https://github.com/hashicorp/nomad/issues/23671)]
* cli: Updated example job specifications in nomad job init [[GH-24232](https://github.com/hashicorp/nomad/issues/24232)]

BUG FIXES:

* consul: Fixed a bug where broken Consul ACL tokens could block registration and deregistration of services and checks [[GH-24166](https://github.com/hashicorp/nomad/issues/24166)]
* consul: Fixed a bug where service deregistration could fail because Consul ACL tokens were revoked during allocation GC [[GH-24166](https://github.com/hashicorp/nomad/issues/24166)]
* docker: Always negotiate API version when initializing clients [[GH-24237](https://github.com/hashicorp/nomad/issues/24237)]
* docker: Fix incorrect auth parsing for private registries [[GH-24215](https://github.com/hashicorp/nomad/issues/24215)]
* docker: Fixed a bug where alloc exec could leak a goroutine [[GH-24244](https://github.com/hashicorp/nomad/issues/24244)]
* docker: Fixed a bug where alloc exec with stdin would hang [[GH-24202](https://github.com/hashicorp/nomad/issues/24202)]
* docker: Fixed a bug where task CPU stats were reported incorrectly [[GH-24229](https://github.com/hashicorp/nomad/issues/24229)]
* heartbeat: Fixed a bug where failed nodes would not be marked down [[GH-24241](https://github.com/hashicorp/nomad/issues/24241)]
* scheduler: fixes reconnecting allocations not getting picked correctly when replacements failed [[GH-24165](https://github.com/hashicorp/nomad/issues/24165)]
* ui: Fix an issue where a dropdown on the variables page would appear underneath table headers [[GH-24162](https://github.com/hashicorp/nomad/issues/24162)]
* ui: Put a max-width on token name so it doesn't collide with the search box in the top nav [[GH-24240](https://github.com/hashicorp/nomad/issues/24240)]
* windows: Fixed a bug where a crashed executor would orphan task processes [[GH-24214](https://github.com/hashicorp/nomad/issues/24214)]

## 1.9.0 (October 10, 2024)

BREAKING CHANGES:

* heartbeats: clients older than 1.6.0 will fail heartbeats to 1.9.0+ servers [[GH-23838](https://github.com/hashicorp/nomad/issues/23838)]
* jobspec: Removed support for HCLv1 [[GH-23912](https://github.com/hashicorp/nomad/issues/23912)]
* services: Clients older than 1.5.0 will fail to read Nomad native services via template blocks [[GH-23910](https://github.com/hashicorp/nomad/issues/23910)]
* tls: Removed deprecated `tls.prefer_server_cipher_suites` field from agent configuration [[GH-23712](https://github.com/hashicorp/nomad/issues/23712)]

SECURITY:

* security: Fixed a bug in client FS API where the check to prevent reads from the secrets dir could be bypassed on case-insensitive file systems [[GH-24125](https://github.com/hashicorp/nomad/issues/24125)]

IMPROVEMENTS:

* cli: Added redaction options to operator snapshot commands [[GH-24023](https://github.com/hashicorp/nomad/issues/24023)]
* cli: Increase default log level and duration when capturing logs with `operator debug` [[GH-23850](https://github.com/hashicorp/nomad/issues/23850)]
* deps: Upgraded yamux to v0.1.2 to fix a bug where RPC connections could deadlock [[GH-24058](https://github.com/hashicorp/nomad/issues/24058)]
* docker: Use official docker SDK instead of a 3rd party client [[GH-23966](https://github.com/hashicorp/nomad/issues/23966)]
* identity: Added filepath parameter to identity block for persisting workload identities [[GH-24038](https://github.com/hashicorp/nomad/issues/24038)]
* jobs: Added Version Tags to job versions, to prevent them from being garbage collected and allow for diffs [[GH-24055](https://github.com/hashicorp/nomad/issues/24055)]
* keyring: Stored wrapped data encryption keys in Raft [[GH-23977](https://github.com/hashicorp/nomad/issues/23977)]
* metrics: introduce client config to include alloc metadata as part of the base labels [[GH-23964](https://github.com/hashicorp/nomad/issues/23964)]
* networking: Added an option to ignore static port collisions when scheduling, for programs that use the SO_REUSEPORT unix socket option [[GH-23956](https://github.com/hashicorp/nomad/issues/23956)]
* networking: IPv6 can now be enabled on the Nomad bridge network mode [[GH-23882](https://github.com/hashicorp/nomad/issues/23882)]
* quotas (Enterprise): Added the possibility to set device count limits [[GH-23894](https://github.com/hashicorp/nomad/issues/23894)]
* raft: Bump raft to v1.7.1 which includes pre-vote. This should make servers more stable after network partitions [[GH-24029](https://github.com/hashicorp/nomad/issues/24029)]

BUG FIXES:

* bug: Allow client template config block to be parsed when using json config [[GH-24007](https://github.com/hashicorp/nomad/issues/24007)]
* cli: Fixed a bug in job status command where -t would act as though -json was also set [[GH-24054](https://github.com/hashicorp/nomad/issues/24054)]
* scaling: Fixed a bug where scaling policies would not get created during job submission unless namespace field was set in jobspec [[GH-24065](https://github.com/hashicorp/nomad/issues/24065)]
* state: Fixed a bug where compatibility updates for node topology for nodes older than 1.7.0 were not being correctly applied [[GH-24127](https://github.com/hashicorp/nomad/issues/24127)]
* task: adds node.pool attribute to interpretable values in task env [[GH-24052](https://github.com/hashicorp/nomad/issues/24052)]
* template: Fixed a panic on client restart when using change_mode=script [[GH-24057](https://github.com/hashicorp/nomad/issues/24057)]
* ui: Fixes an issue where variables paths would not let namespaced users write variables unless they also had wildcard namespace variable write permissions [[GH-24073](https://github.com/hashicorp/nomad/issues/24073)]

## 1.8.15 Enterprise (July 08, 2025)

BUG FIXES:

* agent: Fixed a bug to prevent a possible panic during graceful shutdown [[GH-26018](https://github.com/hashicorp/nomad/issues/26018)]
* agent: Fixed a bug to prevent panic during graceful server shutdown [[GH-26171](https://github.com/hashicorp/nomad/issues/26171)]
* agent: Fixed bug where agent would exit early from graceful shutdown when managed by systemd [[GH-26023](https://github.com/hashicorp/nomad/issues/26023)]
* cli: Fixed a bug in the `tls cert create` command that always added ``"<role>.global.nomad"` to the certificate DNS names, even when the specified region was not ``"global"`. [[GH-26086](https://github.com/hashicorp/nomad/issues/26086)]
* client: Fixed bug where drained batch jobs would not be rescheduled if no eligible nodes were immediately available [[GH-26025](https://github.com/hashicorp/nomad/issues/26025)]
* docker: Fixed a bug where very low resources.cpu values could generate invalid cpu weights on hosts with very large client.cpu_total_compute values [[GH-26081](https://github.com/hashicorp/nomad/issues/26081)]
* encrypter: Fixes a bug where waiting for the active keyset wouldn't return correctly
* tls: Fixed a bug where reloading the Nomad server process with an updated `tls.verify_server_hostname` configuration parameter would not apply an update to internal RPC handler verification and require a full server restart [[GH-26107](https://github.com/hashicorp/nomad/issues/26107)]
* vault: Fixed a bug where non-periodic tokens would not have their TTL incremented to the lease duration [[GH-26041](https://github.com/hashicorp/nomad/issues/26041)]

## 1.8.14 Enterprise (June 10, 2025)

BREAKING CHANGES:

* template: Support for the following non-hermetic sprig functions has been removed: sprig_date, sprig_dateInZone, sprig_dateModify, sprig_htmlDate, sprig_htmlDateInZone, sprig_dateInZone, sprig_dateModify, sprig_randAlphaNum, sprig_randAlpha, sprig_randAscii, sprig_randNumeric, sprig_randBytes, sprig_uuidv4, sprig_env, sprig_expandenv, and sprig_getHostByName. [[GH-25998](https://github.com/hashicorp/nomad/issues/25998)]

SECURITY:

* identity: Fixed bug where workflow identity policies are matched by job ID prefix (CVE-2025-4922) [[GH-25869](https://github.com/hashicorp/nomad/issues/25869)]
* template: Bump the consul-template version to resolve CVE-2025-27144, CVE-2025-22869, CVE-2025-22870 and CVE-2025-22872. [[GH-25998](https://github.com/hashicorp/nomad/issues/25998)]
* template: Removed support to the non-hermetic sprig_env, sprig_expandenv, and sprig_getHostByName sprig functions to prevent potential leakage of environment or network information, since they can allow reading environment variables or resolving domain names to IP addresses. [[GH-25998](https://github.com/hashicorp/nomad/issues/25998)]

IMPROVEMENTS:

* reporting (Enterprise): Added support for offline utilization reporting [[GH-25844](https://github.com/hashicorp/nomad/issues/25844)]

BUG FIXES:

* client: Fixed a bug where disconnect.stop_on_client_after timeouts were extended or ignored [[GH-25946](https://github.com/hashicorp/nomad/issues/25946)]
* csi: Fixed -secret values not being sent with the `nomad volume snapshot delete` command [[GH-26022](https://github.com/hashicorp/nomad/issues/26022)]
* disconnect: Fixed a bug where pending evals for reconnected allocs were not cancelled [[GH-25923](https://github.com/hashicorp/nomad/issues/25923)]
* driver: Allow resources.cpu values above the maximum cpu.share value on Linux [[GH-25963](https://github.com/hashicorp/nomad/issues/25963)]
* job: Ensure sidecar task volume_mounts are added to planning diff object [[GH-25878](https://github.com/hashicorp/nomad/issues/25878)]
* reconnecting client: fix issue where reconcile strategy was sometimes ignored [[GH-25799](https://github.com/hashicorp/nomad/issues/25799)]
* scaling: Set the scaling policies to disabled when a job is stopped [[GH-25911](https://github.com/hashicorp/nomad/issues/25911)]
* scheduler: Fixed a bug where a node with no affinity could be selected over a node with low affinity [[GH-25800](https://github.com/hashicorp/nomad/issues/25800)]
* scheduler: Fixed a bug where planning or running a system job with constraints & previously running allocations would return a failed allocation error [[GH-25850](https://github.com/hashicorp/nomad/issues/25850)]
* telemetry: Fix excess CPU consumption from alloc stats collection [[GH-25870](https://github.com/hashicorp/nomad/issues/25870)]
* telemetry: Fixed a bug where alloc stats were still collected (but not published) if telemetry.publish_allocation_metrics=false. [[GH-25870](https://github.com/hashicorp/nomad/issues/25870)]
* vault: Fixed a bug where poststop tasks could not obtain Vault tokens after the main task failed

## 1.8.13 Enterprise (May 13, 2025)

BREAKING CHANGES:

* core: Errors encountered when reloading agent configuration will now cause agents to exit. Before configuration errors during reloads were only logged. This could lead to agents running but unable to communicate [[GH-25721](https://github.com/hashicorp/nomad/issues/25721)]

SECURITY:

* build: Update Go to 1.24.3 to address CVE-2025-22873 [[GH-25818](https://github.com/hashicorp/nomad/issues/25818)]
* sentinel (Enterprise): Fixed a bug where in some cases hard-mandatory policies could be overridden with -policy-override. CVE-2025-3744.

BUG FIXES:

* agent: Fixed a bug where reloading the agent with systemd notification enabled would cause the agent to be killed by system [[GH-25636](https://github.com/hashicorp/nomad/issues/25636)]
* api: Fixed pagination bug which could result in duplicate results [[GH-25792](https://github.com/hashicorp/nomad/issues/25792)]
* cli: Respect NOMAD_REGION environment variable in operator debug command [[GH-25716](https://github.com/hashicorp/nomad/issues/25716)]
* client: fix failure cleaning up namespace on batch jobs [[GH-25714](https://github.com/hashicorp/nomad/issues/25714)]
* metrics: Fixed a bug where RSS and cache stats would not be reported for docker, exec, and java drivers under Linux cgroups v2 [[GH-25751](https://github.com/hashicorp/nomad/issues/25751)]
* scheduler: Fixed a bug in accounting for resources.cores that could prevent placements on nodes with available cores [[GH-25705](https://github.com/hashicorp/nomad/issues/25705)]
* scheduler: Fixed a bug where draining a node with canaries could result in a stuck deployment [[GH-25726](https://github.com/hashicorp/nomad/issues/25726)]
* scheduler: Fixed a bug where updating the rescheduler tracker could corrupt the state store [[GH-25698](https://github.com/hashicorp/nomad/issues/25698)]
* scheduler: Use core ID when selecting cores. This fixes a panic in the scheduler when the `reservable_cores` is not a contiguous list of core IDs. [[GH-25340](https://github.com/hashicorp/nomad/issues/25340)]
* ui: Fixed a bug where the job list page incorrectly calculated if a job had paused tasks. [[GH-25742](https://github.com/hashicorp/nomad/issues/25742)]

## 1.8.12 Enterprise (April 9, 2025)

IMPROVEMENTS:

* build: Updated Go to 1.24.2 [[GH-25623](https://github.com/hashicorp/nomad/issues/25623)]
* client: Improve memory usage by dropping references to task environment [[GH-25373](https://github.com/hashicorp/nomad/issues/25373)]
* cni: Add a warning log when CNI check commands fail [[GH-25581](https://github.com/hashicorp/nomad/issues/25581)]

BUG FIXES:

* client: remove blocking call during client gc [[GH-25123](https://github.com/hashicorp/nomad/issues/25123)]
* client: skip a task groups shutdown_delay when all tasks have already been deregistered [[GH-25157](https://github.com/hashicorp/nomad/issues/25157)]
* csi: Fixed a CSI ExpandVolume bug where the namespace was left out of the staging path [[GH-25253](https://github.com/hashicorp/nomad/issues/25253)]
* csi: Fixed a bug where GC would attempt and fail to delete plugins that had volumes [[GH-25432](https://github.com/hashicorp/nomad/issues/25432)]
* csi: Fixed a bug where cleaning up volume claims on GC'd nodes would cause errors on the leader [[GH-25428](https://github.com/hashicorp/nomad/issues/25428)]
* csi: Fixed a bug where in-flight CSI RPCs would not be cancelled on client GC or dev agent shutdown [[GH-25472](https://github.com/hashicorp/nomad/issues/25472)]
* drivers: set -1 exit code in case of executor failure for the exec, raw_exec, java, and qemu task drivers [[GH-25453](https://github.com/hashicorp/nomad/issues/25453)]
* job: Ensure migrate block difference is added to planning diff object [[GH-25528](https://github.com/hashicorp/nomad/issues/25528)]
* server: Validate `num_schedulers` configuration parameter is between 0 and the number of CPUs available on the machine [[GH-25441](https://github.com/hashicorp/nomad/issues/25441)]
* services: Fixed a bug where Nomad native services would not be correctly interpolated during in-place updates [[GH-25373](https://github.com/hashicorp/nomad/issues/25373)]
* services: Fixed a bug where task-level services, checks, and identities could interpolate jobspec values from other tasks in the same group [[GH-25373](https://github.com/hashicorp/nomad/issues/25373)]

## 1.8.11 Enterprise (March 11, 2025)

BREAKING CHANGES:

* node: The node attribute `consul.addr.dns` has been changed to `unique.consul.addr.dns`. The node attribute `nomad.advertise.address` has been changed to `unique.advertise.address`. [[GH-24942](https://github.com/hashicorp/nomad/issues/24942)]

SECURITY:

* auth: Redact OIDC client secret from API responses and event stream ([CVE-2025-1296](https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2025-1296)) [[GH-25328](https://github.com/hashicorp/nomad/issues/25328)]

IMPROVEMENTS:

* build: Updated Go to 1.24.1 [[GH-25249](https://github.com/hashicorp/nomad/issues/25249)]
* metrics: Fix the process lookup for raw_exec when running rootless [[GH-25198](https://github.com/hashicorp/nomad/issues/25198)]

BUG FIXES:

* cli: Add node_prefix read when setting up the task workload identity Consul policy [[GH-25310](https://github.com/hashicorp/nomad/issues/25310)]
* cni: Fixed a bug where CNI state was not migrated after upgrade, resulting in IP collisions [[GH-25093](https://github.com/hashicorp/nomad/issues/25093)]
* csi: Fixed a bug where plugins that failed initial fingerprints would not be restarted [[GH-25307](https://github.com/hashicorp/nomad/issues/25307)]
* rpc: Fixed a bug that would cause the reader side of RPC connections to hang indefinitely [[GH-25201](https://github.com/hashicorp/nomad/issues/25201)]
* scheduler: Fixed a bug where node class hashes included unique attributes, making scheduling more costly [[GH-24942](https://github.com/hashicorp/nomad/issues/24942)]
* template: Fixed a bug where unset client.template retry blocks ignored defaults [[GH-25113](https://github.com/hashicorp/nomad/issues/25113)]
* template: Updated the consul-template dependency to v0.40.0 which included a bug fix in the quiescence timers. This bug could cause increased Nomad client CPU usage for tasks which use two or more template blocks. [[GH-25140](https://github.com/hashicorp/nomad/issues/25140)]

## 1.8.10 (February 11, 2025)

SECURITY:

* api: sanitize the SignedIdentities in allocations of events to clean the identity token. [[GH-24966](https://github.com/hashicorp/nomad/issues/24966)]
* build: Updated Go to 1.23.6 [[GH-25041](https://github.com/hashicorp/nomad/issues/25041)]
* event stream: fixes vulnerability CVE-2025-0937, where using a wildcard namespace to subscribe to the events API grants a user with "read" capabilites on any namespace, the ability to read events from all namespaces. [[GH-25089](https://github.com/hashicorp/nomad/issues/25089)]

IMPROVEMENTS:

* auth: adds `VerboseLogging` option to auth-method config for debugging SSO [[GH-24892](https://github.com/hashicorp/nomad/issues/24892)]
* event stream: adds ability to authenticate using workload identities [[GH-24849](https://github.com/hashicorp/nomad/issues/24849)]

BUG FIXES:

* agent: Fixed a bug where Nomad error log messages within syslog showed via the notice priority [[GH-24820](https://github.com/hashicorp/nomad/issues/24820)]
* agent: Fixed a bug where all syslog entries were marked as notice when using JSON logging format [[GH-24865](https://github.com/hashicorp/nomad/issues/24865)]
* client: Fixed a bug where temporary RPC errors cause the client to poll for changes more frequently thereafter [[GH-25039](https://github.com/hashicorp/nomad/issues/25039)]
* csi: Fixed a bug where volume context from the plugin would be erased on volume updates [[GH-24922](https://github.com/hashicorp/nomad/issues/24922)]
* networking: check network namespaces on Linux during client restarts and fail the allocation if an existing namespace is invalid [[GH-24658](https://github.com/hashicorp/nomad/issues/24658)]
* reporting (Enterprise): Updated the reporting metric to utilize node active heartbeat count. [[GH-24919](https://github.com/hashicorp/nomad/issues/24919)]
* state store: fix for setting correct status for a job version when reverting, and also fixes an issue where jobs were briefly marked dead during restarts [[GH-24974](https://github.com/hashicorp/nomad/issues/24974)]
* taskrunner: fix panic when a task with dynamic user is recovered [[GH-24739](https://github.com/hashicorp/nomad/issues/24739)]
* ui: Ensure pending service check blocks are filled [[GH-24818](https://github.com/hashicorp/nomad/issues/24818)]
* ui: Remove unrequired node read API call when attempting to stream task logs [[GH-24973](https://github.com/hashicorp/nomad/issues/24973)]
* vault: Fixed a bug where successful renewal was logged as an error [[GH-25040](https://github.com/hashicorp/nomad/issues/25040)]


## 1.8.9 (January 14, 2025)


IMPROVEMENTS:

* api: Sanitise hcl variables before storage on JobSubmission [[GH-24423](https://github.com/hashicorp/nomad/issues/24423)]
* deps: Upgraded aws-sdk-go from v1 to v2 [[GH-24720](https://github.com/hashicorp/nomad/issues/24720)]

BUG FIXES:

* drivers: validate logmon plugin during reattach [[GH-24798](https://github.com/hashicorp/nomad/issues/24798)]


## 1.8.8 Enterprise (December 18, 2024)

SECURITY:

* api: sanitize the SignedIdentities in allocations to prevent privilege escalation through unredacted workload identity token impersonation associated with ACL policies. [[GH-24683](https://github.com/hashicorp/nomad/issues/24683)]
* security: Added more host environment variables to the default deny list for tasks [[GH-24540](https://github.com/hashicorp/nomad/issues/24540)]
* security: Explicitly set 'Content-Type' header to mitigate XSS vulnerability [[GH-24489](https://github.com/hashicorp/nomad/issues/24489)]
* security: add executeTemplate to default template function_denylist [[GH-24541](https://github.com/hashicorp/nomad/issues/24541)]

BUG FIXES:

* agent: Fixed a bug where `retry_join` gave up after a single failure, rather than retrying until max attempts had been reached [[GH-24561](https://github.com/hashicorp/nomad/issues/24561)]
* api: Fixed a bug where alloc exec/logs/fs APIs would return errors for non-global regions [[GH-24644](https://github.com/hashicorp/nomad/issues/24644)]
* cli: Ensure the `operator autopilot health` command only outputs JSON when the `json` flag is supplied [[GH-24655](https://github.com/hashicorp/nomad/issues/24655)]
* consul: Fixed a bug where failures when syncing Consul checks could panic the Nomad agent [[GH-24513](https://github.com/hashicorp/nomad/issues/24513)]
* consul: Fixed a bug where non-root Nomad agents could not recreate a task's Consul token on task restart [[GH-24410](https://github.com/hashicorp/nomad/issues/24410)]
* csi: Fixed a bug where drivers that emit multiple topology segments would cause placements to fail [[GH-24522](https://github.com/hashicorp/nomad/issues/24522)]
* csi: Removed redundant namespace output from volume status command [[GH-24432](https://github.com/hashicorp/nomad/issues/24432)]
* discovery: Fixed a bug where IPv6 addresses would not be accepted from cloud autojoin [[GH-24649](https://github.com/hashicorp/nomad/issues/24649)]
* drivers: fix executor leak when drivers error starting tasks [[GH-24495](https://github.com/hashicorp/nomad/issues/24495)]
* executor: validate executor on reattach to avoid possibility of killing non-Nomad processes [[GH-24538](https://github.com/hashicorp/nomad/issues/24538)]
* fix: handles consul template re-renders on client restart [[GH-24399](https://github.com/hashicorp/nomad/issues/24399)]
* networking: use a tmpfs location for the state of CNI IPAM plugin used by bridge mode, to fix a bug where allocations would fail to restore after host reboot [[GH-24650](https://github.com/hashicorp/nomad/issues/24650)]
* scheduler: take all assigned cpu cores into account instead of only those part of the largest lifecycle [[GH-24304](https://github.com/hashicorp/nomad/issues/24304)]
* vault: Fixed a bug where expired secret leases were treated as non-fatal and retried [[GH-24409](https://github.com/hashicorp/nomad/issues/24409)]

## 1.8.7 Enterprise (November 8, 2024)

SECURITY:

* csi: Fixed a bug where a user with csi-write-volume permissions to one namespace can create volumes in another namespace (CVE-2024-10975) [[GH-24396](https://github.com/hashicorp/nomad/issues/24396)]

BUG FIXES:

* connect: add validation to ensure that connect native services specify a port [[GH-24329](https://github.com/hashicorp/nomad/issues/24329)]
* keyring: Fixed a panic on server startup when decrypting AEAD key data with empty RSA block [[GH-24383](https://github.com/hashicorp/nomad/issues/24383)]
* scheduler: fixed a bug where resource calculation did not account correctly for poststart tasks [[GH-24297](https://github.com/hashicorp/nomad/issues/24297)]

## 1.8.6 Enterprise(October 21, 2024)

IMPROVEMENTS:

* cli: Added synopsis for `operator root` and `operator gossip` command [[GH-23671](https://github.com/hashicorp/nomad/issues/23671)]

BUG FIXES:

* consul: Fixed a bug where broken Consul ACL tokens could block registration and deregistration of services and checks [[GH-24166](https://github.com/hashicorp/nomad/issues/24166)]
* consul: Fixed a bug where service deregistration could fail because Consul ACL tokens were revoked during allocation GC [[GH-24166](https://github.com/hashicorp/nomad/issues/24166)]
* deps: Fixed a bug where restarting Nomad could cause an unrelated process with the same PID as a failed executor to be killed [[GH-24265](https://github.com/hashicorp/nomad/issues/24265)]
* scheduler: fixes reconnecting allocations not getting picked correctly when replacements failed [[GH-24165](https://github.com/hashicorp/nomad/issues/24165)]
* windows: Fixed a bug where a crashed executor would orphan task processes [[GH-24214](https://github.com/hashicorp/nomad/issues/24214)]

## 1.8.5 Enterprise (October 10, 2024)

SECURITY:

* security: Fixed a bug in client FS API where the check to prevent reads from the secrets dir could be bypassed on case-insensitive file systems [[GH-24125](https://github.com/hashicorp/nomad/issues/24125)]

IMPROVEMENTS:

* cli: Increase default log level and duration when capturing logs with `operator debug` [[GH-23850](https://github.com/hashicorp/nomad/issues/23850)]

BUG FIXES:

* bug: Allow client template config block to be parsed when using json config [[GH-24007](https://github.com/hashicorp/nomad/issues/24007)]
* cli: Fixed a bug in job status command where -t would act as though -json was also set [[GH-24054](https://github.com/hashicorp/nomad/issues/24054)]
* licensing: Fixed a bug where environment variable to opt-out of reporting was not respected
* scaling: Fixed a bug where scaling policies would not get created during job submission unless namespace field was set in jobspec [[GH-24065](https://github.com/hashicorp/nomad/issues/24065)]
* state: Fixed a bug where compatibility updates for node topology for nodes older than 1.7.0 were not being correctly applied [[GH-24127](https://github.com/hashicorp/nomad/issues/24127)]
* task: adds node.pool attribute to interpretable values in task env [[GH-24052](https://github.com/hashicorp/nomad/issues/24052)]
* template: Fixed a panic on client restart when using change_mode=script [[GH-24057](https://github.com/hashicorp/nomad/issues/24057)]

## 1.8.4 (September 17, 2024)

BREAKING CHANGES:

* docker: The default infra_image for pause containers is now registry.k8s.io/pause [[GH-23927](https://github.com/hashicorp/nomad/issues/23927)]

IMPROVEMENTS:

* build: update to go1.22.6 [[GH-23805](https://github.com/hashicorp/nomad/issues/23805)]
* cgroups: Allow clients with delegated cgroups check that required cgroup v2 controllers exist [[GH-23803](https://github.com/hashicorp/nomad/issues/23803)]
* docker: Disable cpuset management for non-root clients [[GH-23804](https://github.com/hashicorp/nomad/issues/23804)]
* identity: Added support for server-configured additional claims on the Vault default_identity block [[GH-23675](https://github.com/hashicorp/nomad/issues/23675)]
* namespaces: Allow enabling/disabling allowed network modes per namespace [[GH-23813](https://github.com/hashicorp/nomad/issues/23813)]
* ui: Badge added for Scaled Down jobs [[GH-23829](https://github.com/hashicorp/nomad/issues/23829)]

DEPRECATIONS:

* api: the JobParseRequest.HCLv1 field will be removed in Nomad 1.9.0 [[GH-23913](https://github.com/hashicorp/nomad/issues/23913)]
* jobspec: using the -hcl1 flag for HCLv1 job specifications will now emit a warning at the command line. This feature will be removed in Nomad 1.9.0 [[GH-23913](https://github.com/hashicorp/nomad/issues/23913)]

BUG FIXES:

* identity: Fixed a bug where dispatch and periodic jobs would have their job ID and not parent job ID used when creating the subject claim [[GH-23902](https://github.com/hashicorp/nomad/issues/23902)]
* identity: Fixed a bug where dispatch and periodic jobs would have their job ID and not parent job ID used when interpolating vault.default_identity.extra_claims [[GH-23817](https://github.com/hashicorp/nomad/issues/23817)]
* node: Fixed bug where sysbatch allocations were started prematurely [[GH-23858](https://github.com/hashicorp/nomad/issues/23858)]
* ui: Fix an issue where cmd+click or ctrl+click would double-open a job [[GH-23832](https://github.com/hashicorp/nomad/issues/23832)]

## 1.8.3 (August 13, 2024)

SECURITY:

* security: Fix symlink escape during unarchiving by removing existing paths within the same allocdir. Compromising the Nomad client agent at the source allocation first is a prerequisite for leveraging this issue. [[GH-23738](https://github.com/hashicorp/nomad/issues/23738)]

IMPROVEMENTS:

* acl: Submitting a policy with a leading `/` in a variable path will now return an error to prevent improperly working policies. [[GH-23757](https://github.com/hashicorp/nomad/issues/23757)]
* cli: Added option to return original HCL in `job inspect` command [[GH-23699](https://github.com/hashicorp/nomad/issues/23699)]
* cli: Added support for updating the roles for an ACL token [[GH-18532](https://github.com/hashicorp/nomad/issues/18532)]
* cli: `acl token create` will now emit a warning if the token has a policy that does not yet exist [[GH-16437](https://github.com/hashicorp/nomad/issues/16437)]
* keyring: Added support for encrypting the keyring via Vault transit or external KMS [[GH-23580](https://github.com/hashicorp/nomad/issues/23580)]
* keyring: Added support for prepublishing keys [[GH-23577](https://github.com/hashicorp/nomad/issues/23577)]
* identity: Added support for server-configured additional claims on the Vault default_identity block [[GH-23675](https://github.com/hashicorp/nomad/issues/23675)]
* metrics: Added `client.tasks` metrics to track task states [[GH-23773](https://github.com/hashicorp/nomad/issues/23773)]
* resources: Added `resources.secrets` field to configure size of secrets directory on Linux [[GH-23696](https://github.com/hashicorp/nomad/issues/23696)]
* tls: Allow setting the `tls_min_version` field to `"tls13"` [[GH-23713](https://github.com/hashicorp/nomad/issues/23713)]
* ui: added a Pack badge to the jobs index page for jobs run via Nomad Pack [[GH-23404](https://github.com/hashicorp/nomad/issues/23404)]

BUG FIXES:

* api: Fixed a bug where an `api.Config` targeting a unix domain socket could not be reused between clients [[GH-23785](https://github.com/hashicorp/nomad/issues/23785)]
* cni: .conf and .json config files are now parsed properly [[GH-23629](https://github.com/hashicorp/nomad/issues/23629)]
* cni: network.cni jobspec updates now replace allocs to apply the new network config [[GH-23764](https://github.com/hashicorp/nomad/issues/23764)]
* docker: Fixed a bug where plugin SELinux labels would conflict with read-only `volume` options [[GH-23750](https://github.com/hashicorp/nomad/issues/23750)]
* identity: Fixed a bug where a missing default task identity could panic the leader [[GH-23763](https://github.com/hashicorp/nomad/issues/23763)]
* keyring: Fixed a bug where keys could be garbage collected before workload identities expire [[GH-23577](https://github.com/hashicorp/nomad/issues/23577)]
* keyring: Fixed a bug where keys would never exit the "rekeying" state after a rotation with the `-full` flag [[GH-23577](https://github.com/hashicorp/nomad/issues/23577)]
* keyring: Fixed a bug where periodic key rotation would not occur [[GH-23577](https://github.com/hashicorp/nomad/issues/23577)]
* networking: The same static port can now be used more than once on host networks with multiple IPs [[GH-23693](https://github.com/hashicorp/nomad/issues/23693)]
* scaling: Fixed a bug where state store corruption could occur when writing scaling events [[GH-23673](https://github.com/hashicorp/nomad/issues/23673)]
* template: Fixed a bug where change_mode = "script" would not execute after a client restart [[GH-23663](https://github.com/hashicorp/nomad/issues/23663)]
* ui: Fixed storage/plugin 404s by unescaping a slash character in the request URL [[GH-23625](https://github.com/hashicorp/nomad/issues/23625)]
* windows: Fix bug with containers capabilities on Docker CE [[GH-23599](https://github.com/hashicorp/nomad/issues/23599)]

## 1.8.2 (July 16, 2024)

BREAKING CHANGES:

* docker: default to hyper-v isolation mode on Windows [[GH-23452](https://github.com/hashicorp/nomad/issues/23452)]

SECURITY:

* build: Updated Go to 1.22.5 to address CVE-2024-24791 [[GH-23498](https://github.com/hashicorp/nomad/issues/23498)]
* migration: Added a check for relative paths escaping the allocation directory when unpacking archive during migration, to harden clients against compromised peer clients sending malicious archives [[GH-23319](https://github.com/hashicorp/nomad/issues/23319)]
* security: Removed insecure TLS cipher suites: `TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256`, `TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA25` and `TLS_RSA_WITH_AES_128_CBC_SHA256`. [[GH-23551](https://github.com/hashicorp/nomad/issues/23551)]

IMPROVEMENTS:

* client: add a preferred_address_family config to prefer ipv4 or ipv6 when deducing IP from network interface [[GH-23389](https://github.com/hashicorp/nomad/issues/23389)]
* cni: allow users to input CNI args in job specification [[GH-23538](https://github.com/hashicorp/nomad/issues/23538)]
* deps: Updated Consul API to 1.29.1. [[GH-23436](https://github.com/hashicorp/nomad/issues/23436)]
* deps: Updated consul-template to 0.39 to allow admin partition and sameness groups queries. [[GH-23436](https://github.com/hashicorp/nomad/issues/23436)]
* docker: Validate that unprivileged containers aren't running as ContainerAdmin on Windows [[GH-23443](https://github.com/hashicorp/nomad/issues/23443)]
* namespaces: Added warnings if deleting namespaces that have existing objects associated with them [[GH-23499](https://github.com/hashicorp/nomad/issues/23499)]
* quota (Enterprise): Allow CPU cores to be configured within a quota [[GH-23543](https://github.com/hashicorp/nomad/issues/23543)]
* scaling: Added `-check-index` support to `job scale` command [[GH-23457](https://github.com/hashicorp/nomad/issues/23457)]
* ui: Allow users to create Global ACL tokens from the Administration UI [[GH-23506](https://github.com/hashicorp/nomad/issues/23506)]
* ui: Update headers in the Admin section to use the HashiCorp Design System [[GH-23366](https://github.com/hashicorp/nomad/issues/23366)]
* ui: allow for multiple namespaces in jobs index filters [[GH-23468](https://github.com/hashicorp/nomad/issues/23468)]

BUG FIXES:

* api: Fixed bug where newlines in JobSubmission vars weren't encoded correctly [[GH-23560](https://github.com/hashicorp/nomad/issues/23560)]
* cli: Fixed bug where the `plugin status` command would fail if the plugin ID was a prefix of another plugin ID [[GH-23502](https://github.com/hashicorp/nomad/issues/23502)]
* cli: Fixed bug where the `quota status` and `quota inspect` commands would fail if the quota name was a prefix of another quota name [[GH-23502](https://github.com/hashicorp/nomad/issues/23502)]
* cli: Fixed bug where the `scaling policy info` command would fail if the policy ID was a prefix of another policy ID [[GH-23502](https://github.com/hashicorp/nomad/issues/23502)]
* cli: Fixed bug where the `service info` command would fail if the service name was a prefix of another service name in the same namespace [[GH-23502](https://github.com/hashicorp/nomad/issues/23502)]
* cli: Fixed bug where the `volume deregister`, `volume detach`, and `volume status` commands would fail if the volume ID was a prefix of another volume ID in the same namespace [[GH-23502](https://github.com/hashicorp/nomad/issues/23502)]
* consul: Fixed a bug where service registration and Envoy bootstrap would not wait for Consul ACL tokens and services to be replicated to the local agent [[GH-23381](https://github.com/hashicorp/nomad/issues/23381)]
* plugins: Fix panic on systems that don't support NUMA [[GH-23399](https://github.com/hashicorp/nomad/issues/23399)]
* qemu: Fixed a bug that prevented `qemu` tasks from running on Linux [[GH-23466](https://github.com/hashicorp/nomad/issues/23466)]
* quota (Enterprise): Fixed a bug where a task's resource core count was not translated to CPU MHz and checked against its quota when performing a job plan [[GH-18876](https://github.com/hashicorp/nomad/issues/18876)]
* scheduler: Fix a bug where reserved resources are not calculated correctly [[GH-23386](https://github.com/hashicorp/nomad/issues/23386)]
* server: Fixed a bug where expiring heartbeats for garbage collected nodes could panic the server [[GH-23383](https://github.com/hashicorp/nomad/issues/23383)]
* template: Fix template rendering on Windows [[GH-23432](https://github.com/hashicorp/nomad/issues/23432)]
* ui: Actions run from jobs with explicit name properties now work from the web UI [[GH-23553](https://github.com/hashicorp/nomad/issues/23553)]
* ui: Don't show keyboard nav hints when taking a screenshot [[GH-23365](https://github.com/hashicorp/nomad/issues/23365)]
* ui: Fix an issue where a remotely purged job would prevent redirect from taking place in the web UI [[GH-23492](https://github.com/hashicorp/nomad/issues/23492)]
* ui: Fix an issue where access to Job Templates in the UI was restricted to variable.write access [[GH-23458](https://github.com/hashicorp/nomad/issues/23458)]
* ui: Fix the Upload Jobspec button on the Run Job page [[GH-23548](https://github.com/hashicorp/nomad/issues/23548)]
* ui: Fixed support for namespace parameter on job statuses API [[GH-23456](https://github.com/hashicorp/nomad/issues/23456)]
* ui: fix an issue where gateway timeouts would cause the jobs list to revert to null, gives users a Pause Fetch option [[GH-23427](https://github.com/hashicorp/nomad/issues/23427)]
* vault: Fixed a bug where requests to derive or renew tokens could be sent to the wrong namespace [[GH-23491](https://github.com/hashicorp/nomad/issues/23491)]

## 1.8.1 (June 19, 2024)

SECURITY:

* build: Updated Go to 1.22.4 to address Go stdlib vulnerabilities CVE-2024-24789 and CVE-2024-24790 [[GH-23172](https://github.com/hashicorp/nomad/issues/23172)]

IMPROVEMENTS:

* api: Add support for setting Notes field for Consul health checks [[GH-22397](https://github.com/hashicorp/nomad/issues/22397)]
* cli: `operator snapshot inspect` now includes details of data in snapshot [[GH-18372](https://github.com/hashicorp/nomad/issues/18372)]
* docker: Added container_exists_attempts plugin configuration variable [[GH-22419](https://github.com/hashicorp/nomad/issues/22419)]
* docker: Added support for oom_score_adj [[GH-23297](https://github.com/hashicorp/nomad/issues/23297)]
* exec: Fixed a bug where `exec` driver tasks would fail on older versions of glibc [[GH-23331](https://github.com/hashicorp/nomad/issues/23331)]
* metrics (Enterprise): Publish quota utilization as metrics [[GH-22912](https://github.com/hashicorp/nomad/issues/22912)]
* raw_exec: Added support for oom_score_adj [[GH-23308](https://github.com/hashicorp/nomad/issues/23308)]
* ui: adds a Stopped label for jobs that a user has manually stopped [[GH-23328](https://github.com/hashicorp/nomad/issues/23328)]
* ui: namespace dropdown gets a search field and supports many namespaces [[GH-20626](https://github.com/hashicorp/nomad/issues/20626)]
* ui: shorten client/node metadata/attributes display and make parent-terminal attributes show up [[GH-23290](https://github.com/hashicorp/nomad/issues/23290)]

BUG FIXES:

* acl: Fix plugin policy validation when checking write permissions [[GH-23274](https://github.com/hashicorp/nomad/issues/23274)]
* api: (Enterprise) fixed Allocations.GetPauseState method discarding the task argument [[GH-23377](https://github.com/hashicorp/nomad/issues/23377)]
* client: Fixed a bug where empty task directories would be left behind [[GH-23237](https://github.com/hashicorp/nomad/issues/23237)]
* connect: fix validation with multiple socket paths [[GH-22312](https://github.com/hashicorp/nomad/issues/22312)]
* consul: (Enterprise) Fixed a bug where gateway config entries were written before Sentinel policies were enforced [[GH-22228](https://github.com/hashicorp/nomad/issues/22228)]
* consul: Fixed a bug where Consul admin partition was not used to login via Consul JWT auth method [[GH-22226](https://github.com/hashicorp/nomad/issues/22226)]
* consul: Fixed a bug where gateway config entries were written to the Nomad server agent's Consul partition and not the client's partition [[GH-22228](https://github.com/hashicorp/nomad/issues/22228)]
* driver: Fixed a bug where the exec, java, and raw_exec drivers would not configure cgroups to allow access to devices provided by device plugins [[GH-22518](https://github.com/hashicorp/nomad/issues/22518)]
* scheduler: Fixed a bug where rescheduled allocations that could not be placed would later ignore their reschedule policy limits [[GH-12319](https://github.com/hashicorp/nomad/issues/12319)]
* task schedule: Fixed a bug where schedules wrongly errored as invalid on the last day of the month [[GH-23329](https://github.com/hashicorp/nomad/issues/23329)]
* ui: unbind job detail running allocations count from job-summary endpoint [[GH-23306](https://github.com/hashicorp/nomad/issues/23306)]

## 1.8.0 (May 28, 2024)

IMPROVEMENTS:

* agent: Added support for systemd readiness notifications [[GH-20528](https://github.com/hashicorp/nomad/issues/20528)]
* api: new /v1/jobs/statuses endpoint collates details about jobs' allocs and latest deployment, intended for use in the updated UI jobs index page [[GH-20130](https://github.com/hashicorp/nomad/issues/20130)]
* artifact: Added support for downloading artifacts without validating the TLS certificate [[GH-20126](https://github.com/hashicorp/nomad/issues/20126)]
* autopilot: Added `operator autopilot health` command to review Autopilot health data [[GH-20156](https://github.com/hashicorp/nomad/issues/20156)]
* cli: Add `-jwks-ca-file` argument to `setup consul/vault` commands [[GH-20518](https://github.com/hashicorp/nomad/issues/20518)]
* client/volumes: Add a mount volume level option for selinux tags on volumes [[GH-19839](https://github.com/hashicorp/nomad/issues/19839)]
* client: expose network namespace bridge/cni configuration values as task env vars [[GH-11810](https://github.com/hashicorp/nomad/issues/11810)]
* connect: Added support for `volume_mount` blocks on sidecar task overrides [[GH-20575](https://github.com/hashicorp/nomad/issues/20575)]
* consul/connect: Attempt autodetection of podman task driver for Connect gateways [[GH-20611](https://github.com/hashicorp/nomad/issues/20611)]
* consul: provide tasks that have Consul tokens the CONSUL_HTTP_TOKEN environment variable [[GH-20519](https://github.com/hashicorp/nomad/issues/20519)]
* core: Do not create evaluations within batch deregister endpoint during job garbage collection [[GH-20510](https://github.com/hashicorp/nomad/issues/20510)]
* csi: Added support for wildcard namespace to `plugin status` command [[GH-20551](https://github.com/hashicorp/nomad/issues/20551)]
* deps: Update msgpack to v2 [[GH-20173](https://github.com/hashicorp/nomad/issues/20173)]
* deps: Updated `docker` dependency to 26.0.1 [[GH-20389](https://github.com/hashicorp/nomad/issues/20389)]
* driver/rawexec: Allow specifying custom cgroups [[GH-20481](https://github.com/hashicorp/nomad/issues/20481)]
* func: Allow custom paths to be added the the getter landlock [[GH-20315](https://github.com/hashicorp/nomad/issues/20315)]
* jobspec: Add a schedule{} block for time based task execution (Enterprise) [[GH-22201](https://github.com/hashicorp/nomad/issues/22201)]
* metrics: Added tracking of enqueue and dequeue times of evaluations to the broker [[GH-20329](https://github.com/hashicorp/nomad/issues/20329)]
* networking: Inject constraints on CNI plugins when using bridge networking [[GH-15473](https://github.com/hashicorp/nomad/issues/15473)]
* scheduler: Added a new configuration to avoid rescheduling allocations if a nodes misses one or more heartbits [[GH-19101](https://github.com/hashicorp/nomad/issues/19101)]
* server: Add new options for reconcilation in case of disconnected nodes [[GH-20029](https://github.com/hashicorp/nomad/issues/20029)]
* ui: Added a UI for creating, editing and deleting Sentinel Policies [[GH-20483](https://github.com/hashicorp/nomad/issues/20483)]
* ui: Added a copy button on Action output [[GH-19496](https://github.com/hashicorp/nomad/issues/19496)]
* ui: Added a new UI block to job spec in order to provide description and links in the Web UI [[GH-18292](https://github.com/hashicorp/nomad/issues/18292)]
* ui: Added token.name information to the top nav for ease of operator debugging [[GH-20539](https://github.com/hashicorp/nomad/issues/20539)]
* ui: Improve error and warning messages for invalid variable and job template paths/names [[GH-19989](https://github.com/hashicorp/nomad/issues/19989)]
* ui: Overhaul of the Jobs Index list page, with live updates, more informative statuses, filter expressions, and pagination [[GH-20452](https://github.com/hashicorp/nomad/issues/20452)]
* ui: Prompt a user before they close an exec window to prevent accidental close-browser-tab shortcuts that overlap with terminal ones [[GH-19985](https://github.com/hashicorp/nomad/issues/19985)]
* ui: Replaced single-line variable value fields with multi-line textarea blocks [[GH-19544](https://github.com/hashicorp/nomad/issues/19544)]
* ui: Updated the style of components in the Variables web ui [[GH-19544](https://github.com/hashicorp/nomad/issues/19544)]
* ui: change the State filter on clients page to split out eligibility and drain status [[GH-18607](https://github.com/hashicorp/nomad/issues/18607)]

BUG FIXES:

* cli: Fix handling of scaling jobs which don't generate evals [[GH-20479](https://github.com/hashicorp/nomad/issues/20479)]
* client: Fix unallocated CPU metric calculation when client reserved CPU is set [[GH-20543](https://github.com/hashicorp/nomad/issues/20543)]
* client: terminate old exec task processes before starting new ones, to avoid accidentally leaving running processes in case of an error [[GH-20500](https://github.com/hashicorp/nomad/issues/20500)]
* config: Fixed a panic triggered by registering a job specifying a Vault cluster that has not been configured within the server [[GH-22227](https://github.com/hashicorp/nomad/issues/22227)]
* core: Fix multiple incorrect type conversion for potential overflows [[GH-20553](https://github.com/hashicorp/nomad/issues/20553)]
* csi: Fixed a bug where concurrent mount and unmount operations could unstage volumes needed by another allocation [[GH-20550](https://github.com/hashicorp/nomad/issues/20550)]
* csi: Fixed a bug where plugins would not be deleted on GC if their job updated the plugin ID [[GH-20555](https://github.com/hashicorp/nomad/issues/20555)]
* csi: Fixed a bug where volumes in different namespaces but the same ID would fail to stage on the same client [[GH-20532](https://github.com/hashicorp/nomad/issues/20532)]
* job endpoint: fix implicit constraint mutation for task-level services [[GH-22229](https://github.com/hashicorp/nomad/issues/22229)]
* quota (Enterprise): Fixed a bug where quota usage would not be freed if a job was purged
* services: Added retry to Nomad service deregistration RPCs during alloc stop [[GH-20596](https://github.com/hashicorp/nomad/issues/20596)]
* services: Fixed bug where Nomad services might not be deregistered when nodes are marked down or allocations are terminal [[GH-20590](https://github.com/hashicorp/nomad/issues/20590)]
* structs: Fix job canonicalization for array type fields [[GH-20522](https://github.com/hashicorp/nomad/issues/20522)]
* ui: Fix a bug where the UI would prompt a user to promote a deployment with unplaced canaries [[GH-20408](https://github.com/hashicorp/nomad/issues/20408)]
* ui: Fixed an issue where keynav would not trigger evaluation sidebar expand [[GH-20047](https://github.com/hashicorp/nomad/issues/20047)]
* ui: Show the namespace in the web UI exec command hint [[GH-20218](https://github.com/hashicorp/nomad/issues/20218)]
* windows: Fixed a regression where scanning task processes was inefficient [[GH-20619](https://github.com/hashicorp/nomad/issues/20619)]

## Unsupported Versions

Versions of Nomad before 1.8.0 are no longer supported. See [CHANGELOG-unsupported.md](./CHANGELOG-unsupported.md) for their changelogs.
