# Table: digitalocean_cdns

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| endpoint | string | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| ttl | int | X | √ |  | 
| certificate_id | string | X | √ |  | 
| custom_domain | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| id | string | √ | √ |  | 
| origin | string | X | √ |  | 


