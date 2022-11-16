# Table: digitalocean_snapshots

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| size_gigabytes | float | X | √ |  | 
| resource_id | string | X | √ |  | 
| resource_type | string | X | √ |  | 
| regions | string_array | X | √ |  | 
| min_disk_size | int | X | √ |  | 
| created_at | string | X | √ |  | 
| tags | string_array | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| id | string | √ | √ |  | 
| name | string | X | √ |  | 


