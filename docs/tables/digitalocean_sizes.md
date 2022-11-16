# Table: digitalocean_sizes

## Primary Keys 

```
slug
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| slug | string | √ | √ |  | 
| price_monthly | float | X | √ |  | 
| regions | string_array | X | √ |  | 
| available | bool | X | √ |  | 
| description | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| memory | int | X | √ |  | 
| vcpus | int | X | √ |  | 
| disk | int | X | √ |  | 
| price_hourly | float | X | √ |  | 
| transfer | float | X | √ |  | 


