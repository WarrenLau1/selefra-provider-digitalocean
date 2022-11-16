# Table: digitalocean_monitoring_alert_policies

## Primary Keys 

```
uuid
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| type | string | X | √ |  | 
| compare | string | X | √ |  | 
| alerts | json | X | √ |  | 
| enabled | bool | X | √ |  | 
| tags | string_array | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| uuid | string | √ | √ |  | 
| description | string | X | √ |  | 
| value | float | X | √ |  | 
| window | string | X | √ |  | 
| entities | string_array | X | √ |  | 


