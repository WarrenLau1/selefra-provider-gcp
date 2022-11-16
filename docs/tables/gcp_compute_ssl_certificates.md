# Table: gcp_compute_ssl_certificates

## Primary Keys 

```
self_link
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| subject_alternative_names | string_array | X | √ |  | 
| expire_time | string | X | √ |  | 
| name | string | X | √ |  | 
| private_key | string | X | √ |  | 
| region | string | X | √ |  | 
| type | string | X | √ |  | 
| self_link | string | √ | √ |  | 
| certificate | string | X | √ |  | 
| description | string | X | √ |  | 
| kind | string | X | √ |  | 
| managed | json | X | √ |  | 
| self_managed | json | X | √ |  | 
| project_id | string | X | √ |  | 
| id | big_int | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| creation_timestamp | string | X | √ |  | 


