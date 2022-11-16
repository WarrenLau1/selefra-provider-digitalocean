# Table: digitalocean_registry_repositories

## Primary Keys 

```
name
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| digitalocean_registries_selefra_id | string | X | X | fk to digitalocean_registries.selefra_id | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| name | string | √ | √ |  | 
| registry_name | string | X | √ |  | 
| latest_tag | json | X | √ |  | 
| tag_count | int | X | √ |  | 


