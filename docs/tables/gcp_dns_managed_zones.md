# Table: gcp_dns_managed_zones

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| visibility | string | X | √ |  | 
| id | big_int | √ | √ |  | 
| kind | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| forwarding_config | json | X | √ |  | 
| reverse_lookup_config | json | X | √ |  | 
| project_id | string | X | √ |  | 
| cloud_logging_config | json | X | √ |  | 
| creation_time | string | X | √ |  | 
| description | string | X | √ |  | 
| dns_name | string | X | √ |  | 
| dnssec_config | json | X | √ |  | 
| service_directory_config | json | X | √ |  | 
| labels | json | X | √ |  | 
| name | string | X | √ |  | 
| name_server_set | string | X | √ |  | 
| name_servers | string_array | X | √ |  | 
| peering_config | json | X | √ |  | 
| private_visibility_config | json | X | √ |  | 


