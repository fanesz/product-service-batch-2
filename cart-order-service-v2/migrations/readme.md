### Run Migration
```
go run migrations/migration.go migrations/sql "host=192.168.1.5 port=5432 user=postgres dbname=shopeefun_orders password=qwerty23 sslmode=disable" up
```

### Down Migration
```
go run migrations/migration.go migrations/sql "host=192.168.1.5 port=5432 user=postgres dbname=shopeefun_orders password=fatannajuda sslmode=disable" down
```

### Create new SQL
```
go run migrations/migration.go migrations/sql "host=192.168.1.5 port=5432 user=postgres dbname=shopeefun_orders sslmode=disable" create add_orders_table sql
```