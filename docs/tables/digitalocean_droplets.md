# Table: digitalocean_droplets

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| tags | string_array | X | √ |  | 
| vcpus | int | X | √ |  | 
| size_slug | string | X | √ |  | 
| kernel | json | X | √ |  | 
| status | string | X | √ |  | 
| networks | json | X | √ |  | 
| vpc_uuid | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| volume_ids | string_array | X | √ |  | 
| id | int | √ | √ |  | 
| size | json | X | √ |  | 
| region | json | X | √ |  | 
| next_backup_window | json | X | √ |  | 
| locked | bool | X | √ |  | 
| memory | int | X | √ |  | 
| disk | int | X | √ |  | 
| image | json | X | √ |  | 
| features | string_array | X | √ |  | 
| created_at | string | X | √ |  | 
| backup_ids | int_array | X | √ |  | 
| snapshot_ids | int_array | X | √ |  | 
| name | string | X | √ |  | 


