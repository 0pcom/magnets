# magnets

A minimalist dynamic golang web application for e-commerce - Work In Progress 2021

view live at:
[magnetosphere.net](https://magnetosphere.net)
a minimal theme variation is also running at:
[antiqueengineer.com](https://antiqueengineer.com)

(Necessity is the mother of invention)


## Table of Contents

<!-- MarkdownTOC levels="1,2,3,4,5" autolink="true" bracket="round" -->
- [Concept](#concept)
- [Prerequisite](#prerequisite)
- [CockroachDB Single Node](#cockroachdb-single-node)
- [CockroachDB Cluster Setup](#cockroachdb-cluster-setup)
  - [Starting the cockroachDB cluster](#starting-the-cockroachdb-cluster)
- [Database Setup](#database-setup)
- [Sync Golang Dependencies](#sync-golang-dependencies)
- [Build the Frontend with Hugo](#build-the-frontend-with-hugo)
- [Build the statik resource](#build-the-statik-resource)
- [Application usage](#application-usage)
- [Example Run](#example-run)
- [Production](#production)
- [Previous Documentation](#previous-documentation)

<!-- /MarkdownTOC -->


## Concept

Hugo is used to generate the html templates and page resources for the go web application.

**The escaped functions (which need to appear in the final template files generated by hugo) are defined in [config.toml](/config.toml)**

These are used in [partial templates](/layouts/partials/product.html)

The process for using hugo to generate templates for golang applications in this particular way __is not, to my knowledge, well documented elsewhere.__

The advantage is either to avoid editing the generated templates in some way to insert the functions used by the web application; or to avoid re-creating what hugo already does quite well - in terms of the vertical assembly of a web site in individual html files- from head to footer, using a templating system.

Thus, hugo is used here as a __templating language,__ or in other words, **this is a way to use [hugo](gohugo.io) with a database.**

The database used here is [cockroachdb](https://www.cockroachlabs.com/docs/v20.2/build-a-go-app-with-cockroachdb-upperdb), with [upper/db](https://tour.upper.io/queries/01) as the database access layer.

[Snipcart.com](https://snipcart.com) provides the shopping cart and payment processing.

## Prerequisite
(a.k.a. makedepends)

Tested on Archlinux

Install dependencies:
```
yay -S cockroachdb go hugo
```
it's recommended to use `cockroachdb-bin` for faster testing and deployment

Additionally, to build the files that are embedded in the generated binary you will need:
```
go get github.com/rakyll/statik
sudo ln -s ~/go/bin/statik /usr/bin/statik
```

(Alternatively, you can add GOBIN to your PATH)

clone this repo:
```
git clone https://github.com/0pcom/magnets
cd magnets
```
## CockroachDB Single Node

The first step is creating the certs used to establish the connection between the go application and cocroachdb. `make single-node` starts the cockroachdb node

in a terminal:
```
make clean0 certs0 single-node
```

A cockroach node is started. **Proceed to [database setup](#database-setup)**

## CockroachDB Cluster Setup

(This is not a required step)

Nodes act as access points to the database. Nodes can be started as-needed to give an access point (in this example) within the local network to the database. **Follow along with the [upstream documentation of this process](https://www.cockroachlabs.com/docs/stable/deploy-cockroachdb-on-premises.html). You will need to sync the clocks first!**

**(note - you will need to change the example addresses and aliases in the Makefile for your cluster)**

```
make certs
```

The certificates are generated, and compressed into an archive. These must be copied to the corresponding node before continuing. Refer to the linked documentation above for a description of this process.

In this example, it is assume that this repository is cloned to the GOPATH on the nodes, and that, for instance, `certs1.tar.gz` is extracted into the cloned repository folder and is renamed `certs` from `certs1`.

### Starting the CockroachDB cluster

(this is not a required step; retained for internal reference and for those who are ambitious.)

on each node, beginning with the primary instance
```
make start1
```

the local instance
```
make start0
```

the third node
```
make start2
```

the fourth node
```
make start3
```
refer to the documentation of cockroachdb for troubleshooting.

## Database Setup

in a new terminal or tab:
```
make db-secure
```

## Sync Go Dependencies

Sync the needed golang dependencies for running the web application
```
go mod init
go mod vendor -v
```

## Build the Frontend with Hugo

The hugo template exists and should be modified or customized if desired. A modification to the hugo.386 theme is employed here.

As previously mentioned, hugo is generating __a template for the golang web application.__

Any hugo theme can be used. Patience and a delicate touch are required when making theme-based changes.

Any partial template in [layouts/partials](/layouts/partials) should be usable with another [hugo theme](https://themes.gohugo.io/), given that the corresponding entries in [config.toml](/config.toml) which are used by the partial are added to the config file

Build the front end (or rebuild to integrate changes)
```
hugo
```

build the frontend for the other domain

```
hugo --baseurl magnetosphereelectronicsurplus.com --destination ./public1
```


The generated html sources are now in the `public` directory.

## Build the statik resource

It is necessary to do something like this in order that any parts of the rendered template are not otherwise directly defined with a [route](/pkg/route/route.go) (such as various page resources) are available to each page of the web application. Routes must be defined for any file (produced by hugo) which contains template functions, or else the template functions will not render and will appear in plaintext.

(requires the statik binary)
```
statik src=./public
go generate
```

Note that the /img directory is not built into the binary but the files included in it (any referenced in the database) still appear in the web app; as the folder is hosted.


## Application usage

```
$ go run cmd/magnets/magnets.go --help
magnetosphere.net website implementation

Usage:
  magnets [flags]
  magnets [command]

Available Commands:
  help        Help about any command
  inv         sub-commands for inventory management - interact with the local cockroachdb instance
  run         run the web application

Flags:
  -h, --help   help for magnets

Use "magnets [command] --help" for more information about a command.

```

help for run subcommand:
```
$ go run cmd/magnets/magnets.go run  --help
run the web application

Usage:
  magnets run [flags]

Flags:
  -h, --help       help for run
  -p, --port int   port to serve on (default 8040)
```
help for inv subcommands
```
$ go run cmd/magnets/magnets.go inv  --help
sub-commands for inventory management - interact with the local cockroachdb instance

Usage:
  magnets inv [command]

Available Commands:
  create      add parts to the database
  delete      delete products and drop tables from the database
  impexp      csv import and export operations
  print       
print parts from the database in the terminal

Flags:
  -h, --help   help for inv

Use "magnets inv [command] --help" for more information about a command.

$ go run cmd/magnets/magnets.go inv create  --help
add parts to the database

Usage:
  magnets inv create [flags]

Flags:
  -n, --createpartno string   Create a part by providing the part number
  -z, --createseriese int     Create equipments with sequential part numbers
  -y, --createseriesp int     Create products with sequential part numbers
  -a, --createtables          Create the tables if they do not exist
  -h, --help                  help for create
  -d, --testequip             create test equipment
  -b, --testprod              create test product

$ go run cmd/magnets/magnets.go inv delete  --help
delete products and drop tables from the database

Usage:
  magnets inv delete [flags]

Flags:
  -E, --deletee   Delete the equipment in the equipments database
  -D, --deletep   Delete the products in the products database
  -e, --drope     Drop equipments table
  -d, --dropp     Drop products table
  -h, --help      help for delete

$ go run cmd/magnets/magnets.go inv impexp --help
csv import and export operations

Usage:
  magnets inv impexp [flags]

Flags:
  -f, --exporte   Export equipments to equipmentsexport01.csv
  -e, --exportp   Export products to productsexport01.csv
  -h, --help      help for impexp
  -j, --importe   Import equipments csv from http://127.0.0.1:8079/equipmentsexport01.csv
  -i, --importp   Import products csv from http://127.0.0.1:8079/productsexport01.csv

	$ go run cmd/magnets/magnets.go inv print --help

print parts from the database in the terminal

Usage:
  magnets inv print [flags]

Flags:
  -h, --help        help for print
  -p, --printinv    Print the inv.Mproducts table to the terminal
  -q, --printinv1   Print the inv.Mequipments table to the terminal
  -v, --vprintinv   More verbose printinventory

```

## Example Run

Create the tables and a sequence of parts. Then, export this for editing:

```
go run cmd/magnets/magnets.go inv create -a
go run cmd/magnets/magnets.go inv create -z 100
go run cmd/magnets/magnets.go inv impexp -e
```

The file is written to the current directory.

(the following process is pending improvement)

To import this file after it has been updated, create a folder called `test` and place the updated csv file in it.

**Serve the file on port 8079** Use your preferred application for this.
example with darkhttpd:

```
cd test
darkhttpd .
```

**FIRST REMOVE THE EXISTING ENTRIES BEFORE IMPORTING PRODUCTS:**

```
go run cmd/magnets/magnets.go inv delete -D
go run cmd/magnets/magnets.go inv impexp -i
```

Print a basic view of the inventory to the terminal:

```
$ go run cmd/magnets/magnets.go inv print -p
Initializing cockroachDB connection
5 Products
2021/08/31 16:15:06 products:
product[684981598321901569]:
	partno:		1
	Qty:		0
	Price:		0.00
	Enable:		false
product[684981598363811841]:
	partno:		2
	Qty:		0
	Price:		0.00
	Enable:		false
product[684981598407917569]:
	partno:		3
	Qty:		0
	Price:		0.00
	Enable:		false
product[684981598452023297]:
	partno:		4
	Qty:		0
	Price:		0.00
	Enable:		false
product[684981598495670273]:
	partno:		5
	Qty:		0
	Price:		0.00
	Enable:		false
...
```

Run the web application:

```
go run cmd/magnets/magnets.go run -p 8040
```


## Production

requires caddy

```
yay -S caddy
```

reverse proxy to port 80 from 8040

```
sudo caddy reverse-proxy --from magnetosphere.net --to localhost:8040
```


# PREVIOUS DOCUMENTATION

## Run the application

test the database connection and view the help menu with no flags:
```
go run main.go
```
output:
```
$ go run main.go
Initializing cockroachDB connection
Usage: magnets -dDctCyepirh
Suggested Demo: magnets -ctpr
-y, --create100             Create 100 parts with sequential part numbers
-C, --createpartno string   Create a part by providing the part number
-c, --createtables          Create the tables if they do not exist
-D, --deleteall             Delete the products in the products database
-d, --droptables            Drop tables
-e, --exportcsv             Export a csv to export01.csv
-h, --help                  show this help menu
-i, --importcsv             Import csv from http://127.0.0.1:8079/export01.csv
-p, --printinventory        Print the inventory to the terminal
-r, --run                   run the web app
-t, --testprod              create test product
-v, --vprintinventory       More verbose printinventory

```

Please note this is the order the flags are executed: -dDctCyepirh

Try the demo:
```
magnets -ctpr
```
* the product table is created in the products database
* test products are created
* the inventory is printed to the screen
* the web application is started (on 127.0.0.1:8040/)

## Managing Inventory

Inventory is managed through importing and exporting csv files. Let us first create a list of blank products, with only the ID and part number

```
go run main.go -y
```

The current limitation is that the inventory can be imported only into the existing __empty__ table. So the table must be cleared first before the csv can be imported. The following example demonstrates this workflow:

```
go run main.go -e
```
The inventory is exported (this is a lazy call directly to the Makefile; `make export`)

Open the CSV, and **be certain to format the first column as text**

Add some values. Make a folder called test
```
mkdir test
```
**Save the file into the directory created above with it's existing name**
Switch to that directory and start a http server
```
cd test
darkhttpd . --port 8079
```

Then, execute the import: (by first removing the existing products as indicated earlier)
```
go run main.go -Di
```

You may get a error message. You may get an error `slow query` which actually masks other errors as multiple errors are not printed. Rerun the transaction until either there is no error or until it is determined that there is error in the input.
