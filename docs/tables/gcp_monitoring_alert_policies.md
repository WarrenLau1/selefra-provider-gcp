# Table: gcp_monitoring_alert_policies

## Primary Keys 

```
name
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| mutation_record | json | X | √ |  | 
| conditions | json | X | √ |  | 
| notification_channels | string_array | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| project_id | string | X | √ |  | 
| display_name | string | X | √ |  | 
| enabled | json | X | √ |  | 
| validity | json | X | √ |  | 
| documentation | json | X | √ |  | 
| user_labels | json | X | √ |  | 
| creation_record | json | X | √ |  | 
| alert_strategy | json | X | √ |  | 
| name | string | √ | √ |  | 
| combiner | big_int | X | √ |  | 


