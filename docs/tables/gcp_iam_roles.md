# Table: gcp_iam_roles

## Primary Keys 

```
project_id, name
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| deleted | bool | X | √ |  | 
| stage | string | X | √ |  | 
| project_id | string | X | √ |  | 
| name | string | X | √ |  | 
| included_permissions | string_array | X | √ |  | 
| title | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| description | string | X | √ |  | 
| etag | string | X | √ |  | 


