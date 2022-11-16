# Table: gcp_serviceusage_services

## Primary Keys 

```
name
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| name | string | √ | √ |  | 
| parent | string | X | √ |  | 
| config | json | X | √ |  | 
| state | big_int | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| project_id | string | X | √ |  | 


