// Routing based on the gorilla/mux router

package main

import (
	"fmt"
//	"html/template"
	"log"
	//"github.com/shopspring/decimal"
//	"net/http"
//	"bufio"
//	"fmt"
//	"os"
//	"time"
//	"strconv"
	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/cockroachdb"
)
//var Serve http.Handler

//var partnoinput string
//var partno Product
//var category string
var products []Product

func main() {
	// /* cockroachdb stuff using upper/db database access layer */ //
	fmt.Printf("Initializing cockroachDB connection\n")
	sess, err := cockroachdb.Open(settings)		//establish the session
	if err != nil {
		log.Fatal("cockroachdb.Open: ", err)
	}
	defer sess.Close()

	//test actions on database
//	createTables(sess)
	//deleteAll(sess)
	createProd(sess)

	// Find().All() maps all the records from the products collection.
	productsCol := Products(sess)
	products = []Product{}
	err = productsCol.Find().All(&products)
	if err != nil {
		log.Fatal("productsCol.Find: ", err)
	}
}


func createProd(sess db.Session) {
	productsTable := sess.Collection("products")

	products := []Product{}
	err = productsTable.Find().All(&products)
	if err != nil {
		log.Fatal("productsTable.Find: ", err)
	}

	// Print the queried information.
//	fmt.Printf("Records in the %q collection:\n", productsTable.Name())
//	for i := range products {
//		fmt.Printf("record #%d: %#v\n", i, products[i])
//	}

fmt.Printf("Creating product\n")
p1 := Product{}
p1.Price = 1
p1.Qty = 5000
p1.PartNo = "cap1500uf6p3e"
p1.Value = 1500
p1.ValUnit = "μF"
p1.VoltsRating = 6.3
p1.TempRating = 105.0
p1.TempUnit = "°C"
p1.PackageType = "Radial"
p1.Materials = "Aluminum"
p1.SubCategory = "Electrolytic"
p1.Category = "Capacitor"
p1.Image1 = "cap1500uf6e.jpg"
p1.Name = fmt.Sprintf("%.0f", p1.Value) + p1.ValUnit + " " + fmt.Sprintf("%.1f", p1.VoltsRating) + "V"
p1.Description1 = p1.Name + " " + fmt.Sprintf("%.0f", p1.TempRating) + " " + p1.TempUnit + " " + p1.PackageType + " " + p1.Materials + " " + p1.SubCategory + " " + p1.Category

	err = Products(sess).InsertReturning(&p1)
	if err != nil {
		log.Fatal("sess.Save: ", err)
	}
	//special chars:
	//	μ°
//	fmt.Printf("Creating second test product 'dummy2'\n")
//	product1 = Product{Name: "dummy2", PartNo:"test1", Description1:descp, Price:1.00}
//	err = Products(sess).InsertReturning(&product1)
//	if err != nil {
//		log.Fatal("sess.Save: ", err)
//	}
// Print the queried information.
fmt.Printf("Records in the %q collection:\n", productsTable.Name())
for i := range products {
	fmt.Printf("record #%d: %#v\n", i, products[i])
}
}

// /* database stuff */ //

var err error
//var sitedata basedata.
var settings = cockroachdb.ConnectionURL{
Host:     "localhost",
Database: "product",
User:     "madmin",
Options: map[string]string{
  // Secure node.
  "sslrootcert": "certs/ca.crt",
  "sslkey":      "certs/client.madmin.key",
  "sslcert":     "certs/client.madmin.crt",
},
}


// Products is a handy way to represent a collection.
func Products(sess db.Session) db.Store {
return sess.Collection("products")
}

// Product is used to represent a single record in the "products" table.
type Product struct {
ID	uint64	`db:"ID,omitempty" json:"ID,omitempty"`
Image1 string `db:"image1,omitempty" json:"image1,omitempty"`
Image2 string `db:"image2,omitempty" json:"image2,omitempty"`
Image3 string `db:"image3,omitempty" json:"image3,omitempty"`
Thumb string `db:"thumb,omitempty" json:"thumb,omitempty"`
Name	string `db:"name" json:"name"`
PartNo string `db:"partno" json:"partno"`
MfgPartNo string `db:"mfgpartno,omitempty" json:"mfgpartno,omitempty"`
MfgName string `db:"mfgname,omitempty" json:"mfgname,omitempty"`
Qty	int64  `db:"quantity" json:"quantity"`
UnlimitQty bool `db:"unlimitqty,omitempty" json:"unlimitqty,omitempty"`
Enable  bool `db:"enable,omitempty" json:"enable,omitempty"`
Price float64  `db:"price" json:"price"`
Msrp float64  `db:"msrp,omitempty" json:"msrp,omitempty"`
Cost float64  `db:"cost,omitempty" json:"cost,omitempty"`
//Sold 	int64  `db:"soldqty,omitempty" json:"soldqty,omitempty"`
MinOrder int64 `db:"minorder" json:"minorder"`
MaxOrder int64 `db:"maxorder,omitempty" json:"maxorder,omitempty"`
Location string `db:"location" json:"location"`
Category string `db:"category" json:"category"`
SubCategory string `db:"subcategory" json:"subcategory"`
Type string  `db:"type,omitempty" json:"type,omitempty"`
PackageType string  `db:"packagetype,omitempty" json:"packagetype,omitempty"`
Technology string  `db:"technology,omitempty" json:"technology,omitempty"`
Materials string  `db:"materials,omitempty" json:"materials,omitempty"`
Value float64 `db:"value,omitempty" json:"value,omitempty"`
ValUnit string  `db:"valunit,omitempty" json:"valunit,omitempty"`
Tolerance float64 `db:"tolerance,omitempty" json:"tolerance,omitempty"`
VoltsRating float64 `db:"voltsrating,omitempty" json:"voltsrating,omitempty"`
AmpsRating float64 `db:"ampsrating,omitempty" json:"ampsrating,omitempty"`
WattsRating float64 `db:"wattsrating,omitempty" json:"temprating,omitempty"`
TempRating float64 `db:"temprating,omitempty" json:"temprating,omitempty"`
TempUnit string `db:"tempunit,omitempty" json:"tempunit,omitempty"`
Description1 string `db:"description1,omitempty" json:"description1,omitempty"`
Description2 string `db:"description2,omitempty" json:"description2,omitempty"`
Color1 string `db:"color1,omitempty" json:"color1,omitempty"`
Color2 string `db:"color2,omitempty" json:"color2,omitempty"`
Sourceinfo string `db:"sourceinfo,omitempty" json:"sourceinfo,omitempty"`
Datasheet string `db:"datasheet,omitempty" json:"datasheet,omitempty"`
Docs string `db:"docs,omitempty" json:"docs,omitempty"`
Reference string `db:"reference,omitempty" json:"reference,omitempty"`
Attributes string `db:"attributes,omitempty" json:"attributes,omitempty"`
Condition string `db:"condition,omitempty" json:"condition,omitempty"`
Note string `db:"note,omitempty" json:"note,omitempty"`
Warning  string `db:"warning,omitempty" json:"warning,omitempty"`
Length float64 `db:"length,omitempty,omitempty" json:"length,omitempty"`
Width float64 `db:"width,omitempty" json:"width,omitempty"`
Height float64 `db:"height,omitempty" json:"height,omitempty"`
WeightLb float64 `db:"weightlb,omitempty" json:"weightlb,omitempty"`
WeightOz float64 `db:"weightoz,omitempty" json:"weightoz,omitempty"`
MetaTitle string `db:"metatitle,omitempty" json:"metatitle,omitempty"`
MetaDesc string `db:"metadesc,omitempty" json:"metadesc,omitempty"`
MetaKeywords string `db:"metakeywords,omitempty" json:"metakeywords,omitempty"`
//todo: add extra control fields
}



//var Collection1 string
// Collection is required in order to create a relation between the Product
// struct and the "products" table.
func (a *Product) Store(sess db.Session) db.Store {
return Products(sess)
}

///*

// createTables creates all the tables that are neccessary to run this example.
func createTables(sess db.Session) error {
fmt.Printf("Creating 'products' table\n")
_, err := sess.SQL().Exec(`
  CREATE TABLE IF NOT EXISTS products (
    ID SERIAL PRIMARY KEY UNIQUE NOT NULL,
    image1 STRING NULL DEFAULT '',
    image2 STRING NULL DEFAULT '',
    image3 STRING NULL DEFAULT '',
    thumb STRING NULL DEFAULT '',
    name STRING NOT NULL,
    partno STRING UNIQUE NOT NULL,
    mfgpartno STRING NULL DEFAULT '',
    mfgname STRING NULL DEFAULT '',
    quantity INT DEFAULT 0,
    unlimitqty BOOL DEFAULT FALSE,
    enable BOOL DEFAULT TRUE,
    price FLOAT DEFAULT 0.00,
    msrp FLOAT DEFAULT 0.00,
    cost FLOAT DEFAULT 0.00,
    minorder INT DEFAULT 1,
    maxorder INT DEFAULT 100,
    location STRING NULL DEFAULT '',
		category STRING NULL DEFAULT '',
		subcategory STRING NULL DEFAULT '',
    type STRING NULL DEFAULT '',
    packagetype STRING NULL DEFAULT '',
    technology STRING NULL DEFAULT '',
		materials STRING NULL DEFAULT '',
    value FLOAT DEFAULT 0.0,
    valunit STRING NULL DEFAULT '',
		tolerance DECIMAL(3,2) DEFAULT 0.00,
    voltsrating FLOAT DEFAULT 0.0,
    ampsrating FLOAT DEFAULT 0.0,
		wattsrating FLOAT DEFAULT 0.0,
		temprating FLOAT DEFAULT 0.0,
		tempunit STRING NULL DEFAULT '',
    description1 STRING NULL DEFAULT '',
    description2 STRING NULL DEFAULT '',
		color1 STRING NULL DEFAULT '',
		color2 STRING NULL DEFAULT '',
    sourceinfo STRING NULL DEFAULT '',
    datasheet STRING NULL DEFAULT '',
    docs STRING NULL DEFAULT '',
    reference STRING NULL DEFAULT '',
    attributes STRING NULL DEFAULT '',
    condition STRING NULL DEFAULT '',
    note STRING NULL DEFAULT '',
    warning STRING NULL DEFAULT '',
    length FLOAT DEFAULT 0.0,
    width FLOAT DEFAULT 0.0,
    height FLOAT DEFAULT 0.0,
    weightlb FLOAT DEFAULT 0.0,
    weightoz FLOAT DEFAULT 0.0,
    metatitle STRING NULL DEFAULT '',
    metadesc STRING NULL DEFAULT '',
    metakeywords STRING NULL DEFAULT ''
  )
  `)

if err != nil {
  return err
}
return nil
}

///*	database test stuff	*///
///*
func deleteAll(sess db.Session) {
fmt.Printf("Clearing tables\n")
//clear tables ; testing
err := Products(sess).Truncate()
if err != nil {
  log.Fatal("Truncate: ", err)
}
}

func createTestProd(sess db.Session) {
fmt.Printf("Creating test product 'dummy'\n")
_desc := "test entry to database"
_img := "test.jpg"
//product1 := Product{}
//product1.setDefaults()
//fmt.Println(product1)
product1 := Product{Name: "dummy", PartNo:"test", Description1: _desc, Price:1.00, Image1: _img, Qty: 10}
err := Products(sess).InsertReturning(&product1)
if err != nil {
    log.Fatal("sess.Save: ", err)
}
fmt.Printf("Creating second test product 'dummy2'\n")
//product1.setDefaults()
product1 = Product{Name: "dummy2", PartNo:"test1", Description1: _desc, Price: 1.00, Qty: 100}
err = Products(sess).InsertReturning(&product1)
if err != nil {
    log.Fatal("sess.Save: ", err)
}

}

//*/
///*	database test stuff	*///
/*
func deleteAll(sess db.Session) {
fmt.Printf("Clearing tables\n")
//clear tables ; testing
err := Products(sess).Truncate()
if err != nil {
  log.Fatal("Truncate: ", err)
}
}
*/
