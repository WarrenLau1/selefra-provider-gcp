# Table: gcp_compute_disks

## Primary Keys 

```
self_link
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| label_fingerprint | string | X | √ |  | 
| kind | string | X | √ |  | 
| provisioned_iops | big_int | X | √ |  | 
| source_disk | string | X | √ |  | 
| source_image_id | string | X | √ |  | 
| status | string | X | √ |  | 
| last_detach_timestamp | string | X | √ |  | 
| satisfies_pzs | bool | X | √ |  | 
| source_snapshot | string | X | √ |  | 
| size_gb | big_int | X | √ |  | 
| users | string_array | X | √ |  | 
| self_link | string | √ | √ |  | 
| location_hint | string | X | √ |  | 
| name | string | X | √ |  | 
| options | string | X | √ |  | 
| region | string | X | √ |  | 
| resource_policies | string_array | X | √ |  | 
| physical_block_size_bytes | big_int | X | √ |  | 
| replica_zones | string_array | X | √ |  | 
| source_snapshot_encryption_key | json | X | √ |  | 
| source_snapshot_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| description | string | X | √ |  | 
| guest_os_features | json | X | √ |  | 
| id | big_int | X | √ |  | 
| source_storage_object | string | X | √ |  | 
| project_id | string | X | √ |  | 
| disk_encryption_key | json | X | √ |  | 
| licenses | string_array | X | √ |  | 
| params | json | X | √ |  | 
| source_image | string | X | √ |  | 
| zone | string | X | √ |  | 
| source_image_encryption_key | json | X | √ |  | 
| type | string | X | √ |  | 
| architecture | string | X | √ |  | 
| creation_timestamp | string | X | √ |  | 
| labels | json | X | √ |  | 
| last_attach_timestamp | string | X | √ |  | 
| license_codes | int_array | X | √ |  | 
| source_disk_id | string | X | √ |  | 


