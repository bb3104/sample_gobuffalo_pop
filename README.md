# sample_gobuffalo_pop

## migrate
https://github.com/gobuffalo/pop

https://libraries.io/github/gobuffalo/pop
```
soda g config -t mysql
```
Created database.yml
```
soda generate model user id:int name:text
soda migrate
```
Roolback
```
soda migrate down
```
