# Table: digitalocean_storage_volumes

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| id | string | √ | √ |  | 
| region | json | X | √ |  | 
| name | string | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| tags | string_array | X | √ |  | 
| droplet_ids | int_array | X | √ |  | 
| size_gigabytes | int | X | √ |  | 
| description | string | X | √ |  | 
| filesystem_type | string | X | √ |  | 
| filesystem_label | string | X | √ |  | 


