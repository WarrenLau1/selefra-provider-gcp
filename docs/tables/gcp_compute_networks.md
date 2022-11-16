# Table: gcp_compute_networks

## Primary Keys 

```
self_link
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| auto_create_subnetworks | bool | X | √ |  | 
| gateway_ipv4 | string | X | √ |  | 
| kind | string | X | √ |  | 
| enable_ula_internal_ipv6 | bool | X | √ |  | 
| self_link | string | √ | √ |  | 
| ipv4_range | string | X | √ |  | 
| creation_timestamp | string | X | √ |  | 
| internal_ipv6_range | string | X | √ |  | 
| mtu | big_int | X | √ |  | 
| subnetworks | string_array | X | √ |  | 
| project_id | string | X | √ |  | 
| description | string | X | √ |  | 
| id | big_int | X | √ |  | 
| peerings | json | X | √ |  | 
| routing_config | json | X | √ |  | 
| self_link_with_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| firewall_policy | string | X | √ |  | 
| name | string | X | √ |  | 
| network_firewall_policy_enforcement_order | string | X | √ |  | 


