# Table: digitalocean_space_cors

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| max_age_seconds | int | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| digitalocean_spaces_selefra_id | string | X | X | fk to digitalocean_spaces.selefra_id | 
| id | string | √ | √ |  | 
| allowed_methods | string_array | X | √ |  | 
| allowed_origins | string_array | X | √ |  | 
| allowed_headers | string_array | X | √ |  | 
| expose_headers | string_array | X | √ |  | 


