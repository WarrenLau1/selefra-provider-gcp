# Table: gcp_kms_crypto_keys

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| labels | json | X | √ |  | 
| primary | json | X | √ |  | 
| purpose | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| gcp_kms_keyrings_selefra_id | string | X | X | fk to gcp_kms_keyrings.selefra_id | 
| project_id | string | X | √ |  | 
| create_time | string | X | √ |  | 
| import_only | bool | X | √ |  | 
| name | string | X | √ |  | 
| rotation_period | string | X | √ |  | 
| version_template | json | X | √ |  | 
| crypto_key_backend | string | X | √ |  | 
| next_rotation_time | string | X | √ |  | 
| destroy_scheduled_duration | string | X | √ |  | 


