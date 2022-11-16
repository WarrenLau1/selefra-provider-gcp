# Table: gcp_dns_policies

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| enable_inbound_forwarding | bool | X | √ |  | 
| enable_logging | bool | X | √ |  | 
| name | string | X | √ |  | 
| networks | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| project_id | string | X | √ |  | 
| id | big_int | √ | √ |  | 
| alternative_name_server_config | json | X | √ |  | 
| description | string | X | √ |  | 
| kind | string | X | √ |  | 


