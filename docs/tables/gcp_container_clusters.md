# Table: gcp_container_clusters

## Primary Keys 

```
self_link
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| node_pools | json | X | √ |  | 
| database_encryption | json | X | √ |  | 
| notification_config | json | X | √ |  | 
| current_node_count | big_int | X | √ |  | 
| location | string | X | √ |  | 
| name | string | X | √ |  | 
| logging_service | string | X | √ |  | 
| addons_config | json | X | √ |  | 
| resource_labels | json | X | √ |  | 
| confidential_nodes | json | X | √ |  | 
| endpoint | string | X | √ |  | 
| current_master_version | string | X | √ |  | 
| status_message | string | X | √ |  | 
| node_pool_auto_config | json | X | √ |  | 
| description | string | X | √ |  | 
| autoscaling | json | X | √ |  | 
| create_time | string | X | √ |  | 
| node_pool_defaults | json | X | √ |  | 
| maintenance_policy | json | X | √ |  | 
| authenticator_groups_config | json | X | √ |  | 
| monitoring_service | string | X | √ |  | 
| vertical_pod_autoscaling | json | X | √ |  | 
| project_id | string | X | √ |  | 
| ip_allocation_policy | json | X | √ |  | 
| resource_usage_export_config | json | X | √ |  | 
| shielded_nodes | json | X | √ |  | 
| instance_group_urls | string_array | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| locations | string_array | X | √ |  | 
| network_config | json | X | √ |  | 
| conditions | json | X | √ |  | 
| enable_kubernetes_alpha | bool | X | √ |  | 
| master_auth | json | X | √ |  | 
| subnetwork | string | X | √ |  | 
| network_policy | json | X | √ |  | 
| status | big_int | X | √ |  | 
| services_ipv4_cidr | string | X | √ |  | 
| enable_tpu | bool | X | √ |  | 
| initial_node_count | big_int | X | √ |  | 
| cluster_ipv4_cidr | string | X | √ |  | 
| label_fingerprint | string | X | √ |  | 
| legacy_abac | json | X | √ |  | 
| identity_service_config | json | X | √ |  | 
| expire_time | string | X | √ |  | 
| id | string | X | √ |  | 
| monitoring_config | json | X | √ |  | 
| self_link | string | √ | √ |  | 
| release_channel | json | X | √ |  | 
| initial_cluster_version | string | X | √ |  | 
| network | string | X | √ |  | 
| autopilot | json | X | √ |  | 
| tpu_ipv4_cidr_block | string | X | √ |  | 
| mesh_certificates | json | X | √ |  | 
| zone | string | X | √ |  | 
| private_cluster_config | json | X | √ |  | 
| binary_authorization | json | X | √ |  | 
| default_max_pods_constraint | json | X | √ |  | 
| current_node_version | string | X | √ |  | 
| node_config | json | X | √ |  | 
| node_ipv4_cidr_size | big_int | X | √ |  | 
| logging_config | json | X | √ |  | 
| workload_identity_config | json | X | √ |  | 
| master_authorized_networks_config | json | X | √ |  | 


