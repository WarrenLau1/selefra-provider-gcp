# Table: gcp_iam_service_account_keys

## Primary Keys 

```
service_account_unique_id, name
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| key_algorithm | string | X | √ |  | 
| public_key_data | string | X | √ |  | 
| valid_after_time | string | X | √ |  | 
| valid_before_time | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| project_id | string | X | √ |  | 
| service_account_unique_id | string | X | √ |  | 
| disabled | bool | X | √ |  | 
| key_origin | string | X | √ |  | 
| key_type | string | X | √ |  | 
| name | string | X | √ |  | 
| private_key_type | string | X | √ |  | 
| gcp_iam_service_accounts_selefra_id | string | X | X | fk to gcp_iam_service_accounts.selefra_id | 


