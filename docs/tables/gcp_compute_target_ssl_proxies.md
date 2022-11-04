# Table: gcp_compute_target_ssl_proxies

## Primary Keys 

```
self_link
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| project_id | string | X | √ |  | 
| certificate_map | string | X | √ |  | 
| id | big_int | X | √ |  | 
| kind | string | X | √ |  | 
| ssl_certificates | string_array | X | √ |  | 
| self_link | string | √ | √ |  | 
| creation_timestamp | string | X | √ |  | 
| description | string | X | √ |  | 
| name | string | X | √ |  | 
| proxy_header | string | X | √ |  | 
| service | string | X | √ |  | 
| ssl_policy | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


