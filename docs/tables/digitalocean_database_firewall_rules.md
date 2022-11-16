# Table: digitalocean_database_firewall_rules

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| uuid | string | X | √ |  | 
| cluster_uuid | string | X | √ |  | 
| type | string | X | √ |  | 
| value | string | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| digitalocean_databases_selefra_id | string | X | X | fk to digitalocean_databases.selefra_id | 
| selefra_id | string | √ | √ | random id | 


