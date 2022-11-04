# Table: gcp_compute_images

## Primary Keys 

```
self_link
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| architecture | string | X | √ |  | 
| source_disk_encryption_key | json | X | √ |  | 
| source_type | string | X | √ |  | 
| status | string | X | √ |  | 
| labels | json | X | √ |  | 
| source_disk | string | X | √ |  | 
| source_snapshot_id | string | X | √ |  | 
| project_id | string | X | √ |  | 
| satisfies_pzs | bool | X | √ |  | 
| storage_locations | string_array | X | √ |  | 
| deprecated | json | X | √ |  | 
| family | string | X | √ |  | 
| id | big_int | X | √ |  | 
| name | string | X | √ |  | 
| source_disk_id | string | X | √ |  | 
| source_snapshot_encryption_key | json | X | √ |  | 
| label_fingerprint | string | X | √ |  | 
| source_image_encryption_key | json | X | √ |  | 
| source_image_id | string | X | √ |  | 
| source_snapshot | string | X | √ |  | 
| shielded_instance_initial_state | json | X | √ |  | 
| archive_size_bytes | big_int | X | √ |  | 
| description | string | X | √ |  | 
| disk_size_gb | big_int | X | √ |  | 
| raw_disk | json | X | √ |  | 
| kind | string | X | √ |  | 
| license_codes | int_array | X | √ |  | 
| licenses | string_array | X | √ |  | 
| source_image | string | X | √ |  | 
| self_link | string | √ | √ |  | 
| creation_timestamp | string | X | √ |  | 
| guest_os_features | json | X | √ |  | 
| image_encryption_key | json | X | √ |  | 


