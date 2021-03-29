// /* */ //

package main

import (
	"fmt"
	"log"
	"net/http"
	"html/template"
	"time"
	"github.com/gorilla/mux"
	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/cockroachdb"
)

var products []Product
func main() {
	fmt.Printf("Initializing cockroachDB connection\n")
	sess, err := cockroachdb.Open(settings)		//establish the session
	if err != nil {
		log.Fatal("cockroachdb.Open: ", err)
	}
	defer sess.Close()
	importCSV(sess)
	productsCol := Products(sess)
	products = []Product{}
	err = productsCol.Find().All(&products) 	// Find().All() maps all the records from the products collection.
	if err != nil {
		log.Fatal("productsCol.Find: ", err)
	}
	log.Printf("products:")
	for i := range products {
		fmt.Printf("product #%d: %#v\n", i, products[i])
//			fmt.Printf("\tproducts[%d]: %d\n", products[i].ID, products[i].PartNo)
	}
}

// /* database stuff */ //
var err error
var settings = cockroachdb.ConnectionURL{
Host:     "localhost",
Database: "product",
User:     "madmin",
Options: map[string]string{ // Secure node.
  "sslrootcert": "certs/ca.crt",
  "sslkey":      "certs/client.madmin.key",
  "sslcert":     "certs/client.madmin.crt",
},
}

func Products(sess db.Session) db.Store {	// Products is a handy way to represent a collection.
return sess.Collection("products")
}

type Product struct {	// Product is used to represent a single record in the "products" table.
Id	int64	`db:"id,omitempty" json:"id,omitempty"`
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
}	//todo: add extra control fields

func (a *Product) Store(sess db.Session) db.Store {// Collection is required in order to create a relation between the Product struct and the "products" table.
return Products(sess)
}
func importCSV(sess db.Session) error {
	fmt.Printf("Importing CSV from http://127.0.0.1:8079/export01.csv\n")
	_, err := sess.SQL().Exec(`
		IMPORT INTO product.products
		CSV DATA (
			'http://127.0.0.1:8079/export01.csv'
			);
			WITH skip = '1';
			`)
	if err != nil {
		return err
	}
	return nil
}
