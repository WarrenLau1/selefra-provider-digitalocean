# Table: digitalocean_firewalls

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| tags | string_array | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| status | string | X | √ |  | 
| id | string | √ | √ |  | 
| name | string | X | √ |  | 
| inbound_rules | json | X | √ |  | 
| outbound_rules | json | X | √ |  | 
| created_at | string | X | √ |  | 
| pending_changes | json | X | √ |  | 
| droplet_ids | int_array | X | √ |  | 


