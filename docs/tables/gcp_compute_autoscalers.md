# Table: gcp_compute_autoscalers

## Primary Keys 

```
self_link
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| kind | string | X | √ |  | 
| scaling_schedule_status | json | X | √ |  | 
| creation_timestamp | string | X | √ |  | 
| id | big_int | X | √ |  | 
| recommended_size | big_int | X | √ |  | 
| region | string | X | √ |  | 
| status_details | json | X | √ |  | 
| zone | string | X | √ |  | 
| project_id | string | X | √ |  | 
| self_link | string | √ | √ |  | 
| autoscaling_policy | json | X | √ |  | 
| description | string | X | √ |  | 
| name | string | X | √ |  | 
| status | string | X | √ |  | 
| target | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


