# Table: gcp_bigquery_tables

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| num_bytes | big_int | X | √ |  | 
| range_partitioning | json | X | √ |  | 
| table_reference | json | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| location | string | X | √ |  | 
| num_long_term_logical_bytes | big_int | X | √ |  | 
| gcp_bigquery_datasets_selefra_id | string | X | X | fk to gcp_bigquery_datasets.selefra_id | 
| clustering | json | X | √ |  | 
| encryption_configuration | json | X | √ |  | 
| labels | json | X | √ |  | 
| last_modified_time | big_int | X | √ |  | 
| num_rows | big_int | X | √ |  | 
| num_total_logical_bytes | big_int | X | √ |  | 
| schema | json | X | √ |  | 
| self_link | string | X | √ |  | 
| description | string | X | √ |  | 
| external_data_configuration | json | X | √ |  | 
| kind | string | X | √ |  | 
| num_long_term_bytes | big_int | X | √ |  | 
| streaming_buffer | json | X | √ |  | 
| view | json | X | √ |  | 
| id | string | X | √ |  | 
| model | json | X | √ |  | 
| num_total_physical_bytes | big_int | X | √ |  | 
| require_partition_filter | bool | X | √ |  | 
| project_id | string | X | √ |  | 
| default_collation | string | X | √ |  | 
| etag | string | X | √ |  | 
| friendly_name | string | X | √ |  | 
| expiration_time | big_int | X | √ |  | 
| num_active_logical_bytes | big_int | X | √ |  | 
| num_long_term_physical_bytes | big_int | X | √ |  | 
| type | string | X | √ |  | 
| clone_definition | json | X | √ |  | 
| materialized_view | json | X | √ |  | 
| num_partitions | big_int | X | √ |  | 
| creation_time | big_int | X | √ |  | 
| max_staleness | string | X | √ |  | 
| num_physical_bytes | big_int | X | √ |  | 
| num_time_travel_physical_bytes | big_int | X | √ |  | 
| num_active_physical_bytes | big_int | X | √ |  | 
| snapshot_definition | json | X | √ |  | 
| time_partitioning | json | X | √ |  | 


