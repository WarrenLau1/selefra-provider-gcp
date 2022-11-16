# Table: gcp_sql_instances

## Primary Keys 

```
self_link
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| self_link | string | √ | √ |  | 
| kind | string | X | √ |  | 
| maintenance_version | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| service_account_email_address | string | X | √ |  | 
| database_installed_version | string | X | √ |  | 
| database_version | string | X | √ |  | 
| gce_zone | string | X | √ |  | 
| replica_configuration | json | X | √ |  | 
| root_password | string | X | √ |  | 
| backend_type | string | X | √ |  | 
| disk_encryption_configuration | json | X | √ |  | 
| etag | string | X | √ |  | 
| name | string | X | √ |  | 
| disk_encryption_status | json | X | √ |  | 
| ip_addresses | json | X | √ |  | 
| master_instance_name | string | X | √ |  | 
| server_ca_cert | json | X | √ |  | 
| settings | json | X | √ |  | 
| scheduled_maintenance | json | X | √ |  | 
| secondary_gce_zone | string | X | √ |  | 
| state | string | X | √ |  | 
| current_disk_size | big_int | X | √ |  | 
| ipv6_address | string | X | √ |  | 
| max_disk_size | big_int | X | √ |  | 
| project | string | X | √ |  | 
| satisfies_pzs | bool | X | √ |  | 
| available_maintenance_versions | string_array | X | √ |  | 
| on_premises_configuration | json | X | √ |  | 
| replica_names | string_array | X | √ |  | 
| project_id | string | X | √ |  | 
| connection_name | string | X | √ |  | 
| instance_type | string | X | √ |  | 
| suspension_reason | string_array | X | √ |  | 
| create_time | string | X | √ |  | 
| failover_replica | json | X | √ |  | 
| out_of_disk_report | json | X | √ |  | 
| region | string | X | √ |  | 


