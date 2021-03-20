# magnets

A golang web app for ecommerce inventory management

Uses [cockroachdb](https://www.cockroachlabs.com/docs/v20.2/build-a-go-app-with-cockroachdb-upperdb) for database and [upper/db](https://tour.upper.io/queries/01) as the database access layer.

Tested on Archlinux

## Prerequisite

```
yay -S cockroachdb-git golang
```

## Run

in a terminal:
```
make secure
```

cockroach node is started 

in a new terminal:
```
make db-secure
go run main.go
```

server starts on 127.0.0.1:8084/
