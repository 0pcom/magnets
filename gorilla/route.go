// Routing based on the gorilla/mux router

package gorilla

import (
//	"crypto/sha1"
	"fmt"
	"html/template"
//	"io"
	"log"
	"net/http"
//	"os"
//	"path/filepath"
	//"strconv"
//	"strings"
	//"sort"
	"time"
	"github.com/gorilla/mux"
)

import (
	"github.com/upper/db/v4"
"github.com/upper/db/v4/adapter/cockroachdb"
)
var Serve http.Handler

//var partnoinput string
var partno Product
var category string
var products []Product

func init() {
	// /* cockroachdb stuff using upper/db database access layer */ //
	fmt.Printf("Initializing cockroachDB connection\n")
	sess, err := cockroachdb.Open(settings)		//establish the session
	if err != nil {
		log.Fatal("cockroachdb.Open: ", err)
	}
	defer sess.Close()

	//test actions on database
	//createTables(sess)
	//deleteAll(sess)
	//createTestProd(sess)

	// Find().All() maps all the records from the products collection.
	productsCol := Products(sess)
	products = []Product{}
	err = productsCol.Find().All(&products)
	if err != nil {
		log.Fatal("productsCol.Find: ", err)
	}

	r := mux.NewRouter()
	r.PathPrefix("/img/").Handler(http.StripPrefix("/img/", http.FileServer(http.Dir("./img"))))
	r.HandleFunc("/", frontPage).Methods("GET")
	r.HandleFunc("/about", aboutPage).Methods("GET")
		r.HandleFunc("/time", timeFunc).Methods("GET")
	//	r.HandleFunc("/gallery", galleryFunc).Methods("GET")
		r.HandleFunc("/products", findProducts).Methods("GET")
		r.HandleFunc("/product/{slug}", findProduct).Methods("GET")
		//r.HandleFunc("/gallery/{slug}", ).Methods("GET")
	Serve = r
}

// /* timepage  */ //

func monthDayYear(t time.Time) string {
	return t.Format("Monday January 2, 2006 15:04:05")
}

func timeFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	var fm = template.FuncMap{ "fdateMDY": monthDayYear,	}
	tp1 := template.Must(template.New("").Funcs(fm).ParseFiles("time.gohtml"))
	if err := tp1.ExecuteTemplate(w, "time.gohtml", time.Now()); err != nil {
		log.Fatalln(err)
	}
}

// /* products page */ //
func findProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	tpl0 := template.Must(template.New("").ParseFiles("products.gohtml"))
	tpl0.ExecuteTemplate(w, "products.gohtml", products)	//fmt.Fprint(w, "products\n")
}

// /* individual product page */ //

func findProduct(w http.ResponseWriter, r *http.Request) {	//, product string
	slug := mux.Vars(r)["slug"]

	for i := range products {
		if products[i].PartNo == slug {
			partno = products[i]
			break
				// Found!
		}
}
if partno.Name == "" {
	fmt.Fprint(w, "No product found for partno:\n", slug)
} else {
	w.Header().Set("Content-Type", "text/html")
	tpl1 := template.Must(template.New("").ParseFiles("product.gohtml"))
	//fmt.Fprint(w, "products\n", partno.Name)
	tpl1.ExecuteTemplate(w, "product.gohtml", partno)
}
}


// /* About Page  */ //
func aboutPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	var fm = template.FuncMap{ "fdateMDY": monthDayYear,	}
	tp1 := template.Must(template.New("").Funcs(fm).ParseFiles("about.gohtml"))
	if err := tp1.ExecuteTemplate(w, "about.gohtml", time.Now()); err != nil {
		log.Fatalln(err)
	}
}

// /* Front Page */ //

func frontPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	var fm = template.FuncMap{ "fdateMDY": monthDayYear,	}
	tp1 := template.Must(template.New("").Funcs(fm).ParseFiles("index.gohtml"))
	if err := tp1.ExecuteTemplate(w, "index.gohtml", time.Now()); err != nil {
		log.Fatalln(err)
	}
}

// /*  */ //

func NewRouter() *mux.Router {
router := mux.NewRouter().StrictSlash(true)

// Choose the folder to serve
staticDir := "/static/"

// Create the route
router.
	PathPrefix(staticDir).
	Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("."+staticDir))))

return router
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
/*
//set default values - does not detect existing values
func(prod *Product) setDefaults(){
prod.Image1 = ""
prod.Image2 = ""
prod.Image3 = ""
prod.Thumb = ""
prod.Name = ""
prod.PartNo = ""
prod.MfgPartNo = ""
prod.MfgName = ""
prod.Qty = 0
prod.UnlimitQty = false
prod.Enable = true
prod.Price = 0.00
prod.Msrp = 0.00
prod.Cost = 0.00
//prod.Sold = 0
prod.MinOrder = 0
prod.MaxOrder = 100
prod.Location = ""
prod.Category = ""
prod.SubCategory = ""
prod.Type = ""
prod.PackageType = ""
prod.Technology = ""
prod.Materials = ""
prod.Value = 0.0
prod.ValUnit = ""
prod.VoltsRating = 0.0
prod.AmpsRating = 0.0
prod.WattsRating = 0.0
prod.TempRating = 0.0
prod.TempUnit = ""
prod.Description1 = ""
prod.Description2 = ""
prod.Color1 = ""
prod.Color2 = ""
prod.Sourceinfo = ""
prod.Datasheet = ""
prod.Docs = ""
prod.Reference = ""
prod.Attributes = ""
prod.Condition = ""
prod.Note = ""
prod.Warning = ""
prod.Length = 0.0
prod.Width = 0.0
prod.Height = 0.0
prod.WeightLb = 0.0
prod.WeightOz = 0.0
prod.MetaTitle = ""
prod.MetaDesc = ""
prod.MetaKeywords = ""
//todo: add extra control fields
}

*/