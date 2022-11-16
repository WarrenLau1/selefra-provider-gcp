# Table: gcp_compute_interconnects

## Primary Keys 

```
self_link
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| description | string | X | √ |  | 
| expected_outages | json | X | √ |  | 
| link_type | string | X | √ |  | 
| name | string | X | √ |  | 
| operational_status | string | X | √ |  | 
| admin_enabled | bool | X | √ |  | 
| interconnect_attachments | string_array | X | √ |  | 
| interconnect_type | string | X | √ |  | 
| location | string | X | √ |  | 
| peer_ip_address | string | X | √ |  | 
| satisfies_pzs | bool | X | √ |  | 
| google_ip_address | string | X | √ |  | 
| creation_timestamp | string | X | √ |  | 
| google_reference_id | string | X | √ |  | 
| id | big_int | X | √ |  | 
| kind | string | X | √ |  | 
| noc_contact_email | string | X | √ |  | 
| provisioned_link_count | big_int | X | √ |  | 
| requested_link_count | big_int | X | √ |  | 
| project_id | string | X | √ |  | 
| circuit_infos | json | X | √ |  | 
| customer_name | string | X | √ |  | 
| state | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| self_link | string | √ | √ |  | 


