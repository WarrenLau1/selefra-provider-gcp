# Table: gcp_logging_metrics

## Primary Keys 

```
name
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| label_extractors | json | X | √ |  | 
| create_time | timestamp | X | √ |  | 
| update_time | timestamp | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| name | string | √ | √ |  | 
| description | string | X | √ |  | 
| value_extractor | string | X | √ |  | 
| metric_descriptor | json | X | √ |  | 
| bucket_options | json | X | √ |  | 
| version | big_int | X | √ |  | 
| project_id | string | X | √ |  | 
| filter | string | X | √ |  | 
| disabled | bool | X | √ |  | 


