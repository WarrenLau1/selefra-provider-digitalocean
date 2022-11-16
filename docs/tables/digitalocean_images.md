# Table: digitalocean_images

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| created_at | string | X | √ |  | 
| id | int | √ | √ |  | 
| name | string | X | √ |  | 
| description | string | X | √ |  | 
| error_message | string | X | √ |  | 
| type | string | X | √ |  | 
| regions | string_array | X | √ |  | 
| size_gigabytes | float | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| slug | string | X | √ |  | 
| public | bool | X | √ |  | 
| tags | string_array | X | √ |  | 
| status | string | X | √ |  | 
| distribution | string | X | √ |  | 
| min_disk_size | int | X | √ |  | 


