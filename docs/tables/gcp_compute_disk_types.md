# Table: gcp_compute_disk_types

## Primary Keys 

```
self_link
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| kind | string | X | √ |  | 
| name | string | X | √ |  | 
| region | string | X | √ |  | 
| zone | string | X | √ |  | 
| project_id | string | X | √ |  | 
| self_link | string | √ | √ |  | 
| creation_timestamp | string | X | √ |  | 
| deprecated | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| default_disk_size_gb | big_int | X | √ |  | 
| description | string | X | √ |  | 
| id | big_int | X | √ |  | 
| valid_disk_size | string | X | √ |  | 


