# Table: gcp_compute_backend_services

## Primary Keys 

```
self_link
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| max_stream_duration | json | X | √ |  | 
| name | string | X | √ |  | 
| network | string | X | √ |  | 
| security_settings | json | X | √ |  | 
| custom_response_headers | string_array | X | √ |  | 
| load_balancing_scheme | string | X | √ |  | 
| locality_lb_policies | json | X | √ |  | 
| security_policy | string | X | √ |  | 
| subsetting | json | X | √ |  | 
| consistent_hash | json | X | √ |  | 
| health_checks | string_array | X | √ |  | 
| id | big_int | X | √ |  | 
| compression_mode | string | X | √ |  | 
| description | string | X | √ |  | 
| iap | json | X | √ |  | 
| kind | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| project_id | string | X | √ |  | 
| self_link | string | √ | √ |  | 
| backends | json | X | √ |  | 
| port | big_int | X | √ |  | 
| custom_request_headers | string_array | X | √ |  | 
| locality_lb_policy | string | X | √ |  | 
| outlier_detection | json | X | √ |  | 
| protocol | string | X | √ |  | 
| service_bindings | string_array | X | √ |  | 
| connection_draining | json | X | √ |  | 
| fingerprint | string | X | √ |  | 
| port_name | string | X | √ |  | 
| enable_cdn | bool | X | √ |  | 
| session_affinity | string | X | √ |  | 
| timeout_sec | big_int | X | √ |  | 
| cdn_policy | json | X | √ |  | 
| circuit_breakers | json | X | √ |  | 
| creation_timestamp | string | X | √ |  | 
| failover_policy | json | X | √ |  | 
| affinity_cookie_ttl_sec | big_int | X | √ |  | 
| connection_tracking_policy | json | X | √ |  | 
| edge_security_policy | string | X | √ |  | 
| log_config | json | X | √ |  | 
| region | string | X | √ |  | 


