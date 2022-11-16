# Table: digitalocean_droplet_neighbors

## Primary Keys 

```
neighbor_id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| neighbor_id | int | √ | √ |  | 
| droplet_id | int | X | √ |  | 
| digitalocean_droplets_selefra_id | string | X | X | fk to digitalocean_droplets.selefra_id | 
| selefra_id | string | √ | √ | primary keys value md5 | 


