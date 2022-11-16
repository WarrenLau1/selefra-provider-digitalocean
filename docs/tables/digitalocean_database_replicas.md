# Table: digitalocean_database_replicas

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| created_at | timestamp | X | √ |  | 
| private_network_uuid | string | X | √ |  | 
| digitalocean_databases_selefra_id | string | X | X | fk to digitalocean_databases.selefra_id | 
| status | string | X | √ |  | 
| connection | json | X | √ |  | 
| private_connection | json | X | √ |  | 
| region | string | X | √ |  | 
| tags | string_array | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| name | string | X | √ |  | 


