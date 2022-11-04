# Table: gcp_storage_buckets

## Primary Keys 

```
name
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| name | string | √ | √ |  | 
| acl | json | X | √ |  | 
| website | json | X | √ |  | 
| etag | string | X | √ |  | 
| project_id | string | X | √ |  | 
| bucket_policy_only | json | X | √ |  | 
| created | timestamp | X | √ |  | 
| lifecycle | json | X | √ |  | 
| retention_policy | json | X | √ |  | 
| project_number | big_int | X | √ |  | 
| uniform_bucket_level_access | json | X | √ |  | 
| default_object_acl | json | X | √ |  | 
| labels | json | X | √ |  | 
| custom_placement_config | json | X | √ |  | 
| versioning_enabled | bool | X | √ |  | 
| storage_class | string | X | √ |  | 
| cors | json | X | √ |  | 
| default_event_based_hold | bool | X | √ |  | 
| predefined_default_object_acl | string | X | √ |  | 
| meta_generation | big_int | X | √ |  | 
| logging | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| location | string | X | √ |  | 
| requester_pays | bool | X | √ |  | 
| encryption | json | X | √ |  | 
| public_access_prevention | big_int | X | √ |  | 
| predefined_acl | string | X | √ |  | 
| location_type | string | X | √ |  | 
| rpo | big_int | X | √ |  | 


