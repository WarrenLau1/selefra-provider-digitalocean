# Table: digitalocean_databases

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| engine | string | X | √ |  | 
| version | string | X | √ |  | 
| num_nodes | int | X | √ |  | 
| private_network_uuid | string | X | √ |  | 
| tags | string_array | X | √ |  | 
| name | string | X | √ |  | 
| maintenance_window | json | X | √ |  | 
| project_id | string | X | √ |  | 
| size | string | X | √ |  | 
| private_connection | json | X | √ |  | 
| db_names | string_array | X | √ |  | 
| region | string | X | √ |  | 
| status | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| id | string | √ | √ |  | 
| users | json | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| connection | json | X | √ |  | 


