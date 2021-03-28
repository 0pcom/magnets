# magnets

A golang web app for ecommerce inventory management

Uses [cockroachdb](https://www.cockroachlabs.com/docs/v20.2/build-a-go-app-with-cockroachdb-upperdb) for database and [upper/db](https://tour.upper.io/queries/01) as the database access layer.

Tested on Archlinux

## Prerequisite

```
yay -S cockroachdb go
```
it's recommended to use cockroachdb-bin

## Database initialization

in a terminal:
```
make secure
```

cockroach node is started

in a new terminal or tab:
```
make db-secure
```

## Sync dependencies
```
go mod init
go mod vendor -v
```

## Run the application
```
go run main.go
```

server starts on 127.0.0.1:8040/


## Production

requires caddy

```
yay -S caddy
```

reverse proxy to port 80 from 8040

```
sudo caddy reverse-proxy --from magnetosphere.net --to localhost:8040

```

## Add nodes / cluster setup
Nodes are an access point to the database
https://www.cockroachlabs.com/docs/stable/deploy-cockroachdb-on-premises.html#recommendations
