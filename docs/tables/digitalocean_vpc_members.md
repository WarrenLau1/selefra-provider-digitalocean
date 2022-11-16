# Table: digitalocean_vpc_members

## Primary Keys 

```
urn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| urn | string | √ | √ |  | 
| name | string | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| digitalocean_vpcs_selefra_id | string | X | X | fk to digitalocean_vpcs.selefra_id | 
| selefra_id | string | √ | √ | primary keys value md5 | 


