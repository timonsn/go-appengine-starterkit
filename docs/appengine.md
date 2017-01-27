#App Engine YAML configuration

## Runtime
`runtime` is one of the required keys for VM hosting environment setup. For standard runtimes, use values `python27`, `go`, `java`, `php55`. For custom runtimes, use value `custom`.

The `runtime`, `vm`, `api_version` keys are required to setup the VM hosting environment. Note that when using custom runtime, although `api_version` must be given, the value is ignored.

## Resources allocation
`resources` defines the machine type that App Engine will use to serve your app based on the amount of CPU and memory you've specified. The machine is guaranteed to have at least the level of resources you've specified, it might have more.

`cpu` is the number of CPU cores that the machine should have.

`memory_gb` is the amount of RAM in Gigabytes that the machine type should have.

`disk_size_gb` is the size of Persistent Disk in Gigabytes that the machine type should have.

## Scaling settings
`automatic_scaling` defines how your application should scale to use more or less VM instances based on the value you specify for target CPU utilization.

`min_number_instances` is the minumum number of VM instances that your app will be served from as it is scaled down.

`max_number_instances` is the maximum number of VM instances that your app will be served from as it is scaled up.

`cool_down_period_sec` is the time interval in seconds between auto scaling checks. The cool-down period must be greater than or equal to 60 seconds.

`target_utilization` is used by the autoscaling service to decide when to reduce or increase the number of VM instances based on the average CPU utilization across all running VM instances.
