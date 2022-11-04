# Table: gcp_storage_bucket_policies

## Primary Keys 

```
bucket_name
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| project_id | string | X | √ |  | 
| bucket_name | string | √ | √ |  | 
| bindings | json | X | √ |  | 
| gcp_storage_buckets_selefra_id | string | X | X | fk to gcp_storage_buckets.selefra_id | 
| selefra_id | string | √ | √ | primary keys value md5 | 


