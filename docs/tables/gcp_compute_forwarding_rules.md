# Table: gcp_compute_forwarding_rules

## Primary Keys 

```
self_link
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| all_ports | bool | X | √ |  | 
| creation_timestamp | string | X | √ |  | 
| metadata_filters | json | X | √ |  | 
| ports | string_array | X | √ |  | 
| target | string | X | √ |  | 
| project_id | string | X | √ |  | 
| self_link | string | √ | √ |  | 
| ip_protocol | string | X | √ |  | 
| service_name | string | X | √ |  | 
| subnetwork | string | X | √ |  | 
| ip_address | string | X | √ |  | 
| description | string | X | √ |  | 
| no_automate_dns_zone | bool | X | √ |  | 
| backend_service | string | X | √ |  | 
| name | string | X | √ |  | 
| network_tier | string | X | √ |  | 
| psc_connection_id | big_int | X | √ |  | 
| service_directory_registrations | json | X | √ |  | 
| kind | string | X | √ |  | 
| label_fingerprint | string | X | √ |  | 
| network | string | X | √ |  | 
| port_range | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| id | big_int | X | √ |  | 
| labels | json | X | √ |  | 
| load_balancing_scheme | string | X | √ |  | 
| ip_version | string | X | √ |  | 
| is_mirroring_collector | bool | X | √ |  | 
| region | string | X | √ |  | 
| fingerprint | string | X | √ |  | 
| psc_connection_status | string | X | √ |  | 
| service_label | string | X | √ |  | 
| allow_global_access | bool | X | √ |  | 


