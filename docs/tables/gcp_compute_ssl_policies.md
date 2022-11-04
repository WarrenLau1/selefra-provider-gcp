# Table: gcp_compute_ssl_policies

## Primary Keys 

```
self_link
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| self_link | string | √ | √ |  | 
| custom_features | string_array | X | √ |  | 
| description | string | X | √ |  | 
| enabled_features | string_array | X | √ |  | 
| profile | string | X | √ |  | 
| warnings | json | X | √ |  | 
| project_id | string | X | √ |  | 
| fingerprint | string | X | √ |  | 
| kind | string | X | √ |  | 
| name | string | X | √ |  | 
| creation_timestamp | string | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| id | big_int | X | √ |  | 
| min_tls_version | string | X | √ |  | 


