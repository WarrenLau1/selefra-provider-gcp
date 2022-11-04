# Table: gcp_compute_firewalls

## Primary Keys 

```
self_link
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| log_config | json | X | √ |  | 
| source_ranges | string_array | X | √ |  | 
| source_tags | string_array | X | √ |  | 
| target_service_accounts | string_array | X | √ |  | 
| target_tags | string_array | X | √ |  | 
| allowed | json | X | √ |  | 
| id | big_int | X | √ |  | 
| destination_ranges | string_array | X | √ |  | 
| direction | string | X | √ |  | 
| network | string | X | √ |  | 
| priority | big_int | X | √ |  | 
| source_service_accounts | string_array | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| self_link | string | √ | √ |  | 
| creation_timestamp | string | X | √ |  | 
| disabled | bool | X | √ |  | 
| name | string | X | √ |  | 
| description | string | X | √ |  | 
| kind | string | X | √ |  | 
| project_id | string | X | √ |  | 
| denied | json | X | √ |  | 


