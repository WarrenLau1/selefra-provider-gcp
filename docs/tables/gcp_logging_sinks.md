# Table: gcp_logging_sinks

## Primary Keys 

```
name
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| include_children | bool | X | √ |  | 
| create_time | timestamp | X | √ |  | 
| project_id | string | X | √ |  | 
| description | string | X | √ |  | 
| exclusions | json | X | √ |  | 
| disabled | bool | X | √ |  | 
| output_version_format | big_int | X | √ |  | 
| writer_identity | string | X | √ |  | 
| update_time | timestamp | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| name | string | √ | √ |  | 
| destination | string | X | √ |  | 
| filter | string | X | √ |  | 


