# Table: gcp_domains_registrations

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| create_time | timestamp | X | √ |  | 
| expire_time | timestamp | X | √ |  | 
| issues | int_array | X | √ |  | 
| pending_contact_settings | json | X | √ |  | 
| labels | json | X | √ |  | 
| contact_settings | json | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| domain_name | string | X | √ |  | 
| state | big_int | X | √ |  | 
| dns_settings | json | X | √ |  | 
| supported_privacy | int_array | X | √ |  | 
| project_id | string | X | √ |  | 
| name | string | X | √ |  | 
| management_settings | json | X | √ |  | 


