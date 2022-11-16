# Table: gcp_redis_instances

## Primary Keys 

```
name
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| memory_size_gb | big_int | X | √ |  | 
| persistence_iam_identity | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| labels | json | X | √ |  | 
| create_time | timestamp | X | √ |  | 
| redis_configs | json | X | √ |  | 
| replica_count | big_int | X | √ |  | 
| location_id | string | X | √ |  | 
| redis_version | string | X | √ |  | 
| authorized_network | string | X | √ |  | 
| display_name | string | X | √ |  | 
| port | big_int | X | √ |  | 
| server_ca_certs | json | X | √ |  | 
| name | string | √ | √ |  | 
| auth_enabled | bool | X | √ |  | 
| nodes | json | X | √ |  | 
| secondary_ip_range | string | X | √ |  | 
| tier | big_int | X | √ |  | 
| status_message | string | X | √ |  | 
| maintenance_schedule | json | X | √ |  | 
| read_replicas_mode | big_int | X | √ |  | 
| alternative_location_id | string | X | √ |  | 
| reserved_ip_range | string | X | √ |  | 
| state | big_int | X | √ |  | 
| current_location_id | string | X | √ |  | 
| transit_encryption_mode | big_int | X | √ |  | 
| maintenance_policy | json | X | √ |  | 
| read_endpoint | string | X | √ |  | 
| read_endpoint_port | big_int | X | √ |  | 
| project_id | string | X | √ |  | 
| host | string | X | √ |  | 
| connect_mode | big_int | X | √ |  | 


