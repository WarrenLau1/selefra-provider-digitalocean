# Table: digitalocean_domain_records

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| type | string | X | √ |  | 
| name | string | X | √ |  | 
| data | string | X | √ |  | 
| priority | int | X | √ |  | 
| port | int | X | √ |  | 
| tag | string | X | √ |  | 
| id | int | √ | √ |  | 
| ttl | int | X | √ |  | 
| weight | int | X | √ |  | 
| flags | int | X | √ |  | 
| digitalocean_domains_selefra_id | string | X | X | fk to digitalocean_domains.selefra_id | 


