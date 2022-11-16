# Table: digitalocean_accounts

## Primary Keys 

```
uuid
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| team | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| uuid | string | √ | √ |  | 
| reserved_ip_limit | int | X | √ |  | 
| volume_limit | int | X | √ |  | 
| email | string | X | √ |  | 
| email_verified | bool | X | √ |  | 
| droplet_limit | int | X | √ |  | 
| floating_ip_limit | int | X | √ |  | 
| status | string | X | √ |  | 
| status_message | string | X | √ |  | 


