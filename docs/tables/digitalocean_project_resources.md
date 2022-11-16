# Table: digitalocean_project_resources

## Primary Keys 

```
urn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| links | json | X | √ |  | 
| status | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| digitalocean_projects_selefra_id | string | X | X | fk to digitalocean_projects.selefra_id | 
| urn | string | √ | √ |  | 
| assigned_at | string | X | √ |  | 


