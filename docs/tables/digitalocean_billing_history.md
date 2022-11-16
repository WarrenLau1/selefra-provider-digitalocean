# Table: digitalocean_billing_history

## Primary Keys 

```
invoice_id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| invoice_uuid | string | X | √ |  | 
| date | timestamp | X | √ |  | 
| type | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| invoice_id | string | √ | √ |  | 
| description | string | X | √ |  | 
| amount | string | X | √ |  | 


