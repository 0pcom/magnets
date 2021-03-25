// Routing based on the gorilla/mux router

package main

import (
	"fmt"
//	"html/template"
	"log"
//	"net/http"
//	"bufio"
//	"fmt"
//	"os"
//	"time"
	"strconv"
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
	createTables(sess)
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


var _id uint64
var _name string
var _partno string
var _desc1 string
var _desc1p *string
var _desc2 string
var _desc2p *string
var _img string
var _imgp *string
var _qty int64
//var _qtyp *int64
var _col1 string
var _col1p *string
var _col2 string
var _col2p *string
var _mpn string
var _mpnp *string
var _mfn string
var _mfnp *string
var _enab bool
var _enabp *bool
var _price float64
var _pricep *float64
var _msrp float64
var _msrpp *float64
var _cost float64
var _costp *float64
var _mino int64
var _minop *int64
var _maxo int64
var _maxop *int64
var _loc string
var _locp *string
var _cat string
var _catp *string
var _scat string
var _scatp *string
var _typ string
var _typp *string
var _ptyp string
var _ptypp *string
var _tech string
var _techp *string
var _mat string
var _matp *string
var _val float64
var _valp *float64
var _valu string
var _valup *string
var _vol float64
var _volp *float64
var _amp float64
var _ampp *float64
var _wat float64
var _watp *float64
var _tmp float64
var _tmpp *float64
var _tmpu string
var _tmpup *string
var _src string
var _srcp *string
var _dat string
var _datp *string
var _doc string
var _docp *string
var _ref string
var _refp *string
var _att string
var _attp *string
var _con string
var _conp *string
var _not string
var _notp *string
var _wrn string
var _wrnp *string
var _d1 float64
var _d1p *float64
var _d2 float64
var _d2p *float64
var _d3 float64
var _d3p *float64
var _wtlb float64
var _wtlbp *float64
var _wtoz float64
var _wtozp *float64
var _metat string
var _metatp *string
var _metad string
var _metadp *string
var _metak string
var _metakp *string



func createProd(sess db.Session) {
	productsTable := sess.Collection("products")
	// Use Find to create a result set (db.Result).
	res := productsTable.Find()
	//count the products to get the ID of the next one
	//_index, err := res.Count()
	//add one to index for ID
	fmt.Printf("Creatinte product\n")

	_price = 1
	_qty = 5000
	_qty1 = strconv.Itoa(_qty)
	_name = strconv.Ftoa(_val, 'g', -1) + _valu + " " + strconv.Ftoa(_vol, 'g', -1) + "V "
	_partno = "cap1500uf6e"
	_val = 1500
	_valu = "μF"
	_vol = 6.3
	_tmp = 105
	_tmpu = "C"
	//_typ = ""
	_ptyp = "Radial"
	_mat = "Aluminum"
	_scat = "Electrolytic"
	_cat = "Capacitor"
	_desc1 = 	strconv.Ftoa(_val, 'g', -1) + _valu + " " + strconv.Ftoa(_vol, 'g', -1) + "V " + strconv.Ftoa(_tmp, 'g', -1) + "°" + _tmpu + " " + _ptyp + _mat + _scat + _cat
	_img = "cap1500uf6e.jpg"

	fmt.Printf("Price: $ %.2f\n", _price)
	fmt.Printf("quantity: $%d\n", _qty)
	fmt.Printf("Product Name: %s\n", _name)
	fmt.Printf("PartNo: %s\n", _partno)
	fmt.Printf("Value: $ %.2f\n", _val)
	fmt.Printf("Value Unit: %s\n", _valu)
	fmt.Printf("Voltage: $ %.2f\n", _vol)
	fmt.Printf("Temperature: $ %.2f\n", _tmp)
	fmt.Printf("Temperature Unit: %s\n", _tmpu)
	fmt.Printf("Package Type: %s\n", _ptyp)
	fmt.Printf("Materials: %s\n", _mat)
	fmt.Printf("subcategory: %s\n", _scat)
	fmt.Printf("category: %s\n", _cat)
	fmt.Printf("description: %s\n", _desc1)
  fmt.Printf("image path: %s\n", _img)


	_valp = &_val
	_valup = &_valu
	_volp = &_vol
	_tmpp = &_tmp
	_tmpup = &_tmpu
	_ptypp = &_ptyp
	_matp = &_mat
	_scatp = &_scat
	_catp = &_cat
	_scatp = &_scat
	_desc1p = &_desc1
	_imgp = &_img


/*
	fmt.Printf("Product Name: ")
  _name, _ := reader.ReadString('\n')
  fmt.Printf("Product Name: %s\n", _name)
	fmt.Printf("PartNo: ")
  _partno, _ := reader.ReadString('\n')
  fmt.Printf("PartNo: %s\n", _partno)
	fmt.Printf("category: ")
	_cat, _ := reader.ReadString('\n')
	_catp = &_cat
	fmt.Printf("category: %s\n", _cat)
	fmt.Printf("price: $")
	_, err = fmt.Scanf("%f", &_price)
	if err != nil {
		fmt.Println(err)
	}
	//  fmt.Printf("You have entered : %f \n", _price)
	//	_p1, _ := reader.ReadString('\n')
	//	_price, _ := strconv.ParseFloat(_p1, 64)
	fmt.Printf("Price: $ %.2f\n", _price)
	//fmt.Printf("price %f", _price)
	fmt.Printf("quantity: ")
	_, err = fmt.Scanf("%d", &_qty)
	if err != nil {
		fmt.Println(err)
	}
	//	_qty1, _ := reader.ReadString('\n')
	//	_qty, _ := strconv.ParseInt(_qty1, 0, 64)

	fmt.Printf("quantity: $%d\n", _qty)
	fmt.Printf("description: ")
  _desc1, _ := reader.ReadString('\n')
	_desc1p = &_desc1
  fmt.Printf("description: %s\n", _desc1)
	fmt.Printf("image path: ")
  _img, _ := reader.ReadString('\n')
	_imgp = &_img
  fmt.Printf("image path: %s\n", _img)
*/
	product1 := Product{Name: _name, PartNo: _partno, Price: _price, Description1: _desc1p, Image1: _imgp, Qty: _qty, Description2: _desc2p, Color1: _col1p, Color2: _col2p, MfgPartNo: _mpnp, MfgName: _mfnp, Enable: _enabp, Msrp: _msrpp, Cost: _costp, MinOrder: _minop, MaxOrder: _maxop, Location: _locp, Category: _catp, Type: _typp, PackageType: _ptypp, Technology: _techp, Value: _valp, ValUnit: _valup, VoltsRating: _volp, AmpsRating: _ampp, WattsRating: _watp, Sourceinfo: _srcp, Datasheet: _datp, Docs: _docp, Reference: _refp, Attributes: _attp, Condition: _conp, Note: _notp, Warning: _wrnp, Length: _d1p, Width: _d2p, Height: _d3p, WeightLb: _wtlbp, WeightOz: _wtozp, MetaTitle: _metatp,  MetaDesc: _metadp, MetaKeywords: _metakp, TempRating: _tmpp, TempUnit: _tmpup}
	err = Products(sess).InsertReturning(&product1)
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
Image1 *string `db:"image1,omitempty" json:"image1,omitempty"`
Image2 *string `db:"image2,omitempty" json:"image2,omitempty"`
Image3 *string `db:"image3,omitempty" json:"image3,omitempty"`
Thumb *string `db:"thumb,omitempty" json:"thumb,omitempty"`
Name	string `db:"name" json:"name"`
PartNo string `db:"partno" json:"partno"`
MfgPartNo *string `db:"mfgpartno,omitempty" json:"mfgpartno,omitempty"`
MfgName *string `db:"mfgname,omitempty" json:"mfgname,omitempty"`
Qty	int64  `db:"quantity" json:"quantity"`
UnlimitQty *bool `db:"unlimitqty,omitempty" json:"unlimitqty,omitempty"`
Enable  *bool `db:"enable,omitempty" json:"enable,omitempty"`
Price float64  `db:"price" json:"price"`
Msrp *float64  `db:"msrp,omitempty" json:"msrp,omitempty"`
Cost *float64  `db:"cost,omitempty" json:"cost,omitempty"`
//Sold 	int64  `db:"soldqty,omitempty" json:"soldqty,omitempty"`
MinOrder *int64 `db:"minorder" json:"minorder"`
MaxOrder *int64 `db:"maxorder,omitempty" json:"maxorder,omitempty"`
Location *string `db:"location" json:"location"`
Category *string `db:"category" json:"category"`
SubCategory *string `db:"subcategory" json:"subcategory"`
Type *string  `db:"type,omitempty" json:"type,omitempty"`
PackageType *string  `db:"packagetype,omitempty" json:"packagetype,omitempty"`
Technology *string  `db:"technology,omitempty" json:"technology,omitempty"`
Materials *string  `db:"materials,omitempty" json:"materials,omitempty"`
Value *float64 `db:"value,omitempty" json:"value,omitempty"`
ValUnit *string  `db:"valunit,omitempty" json:"valunit,omitempty"`
VoltsRating *float64 `db:"voltsrating,omitempty" json:"voltsrating,omitempty"`
AmpsRating *float64 `db:"ampsrating,omitempty" json:"ampsrating,omitempty"`
WattsRating *float64 `db:"wattsrating,omitempty" json:"temprating,omitempty"`
TempRating *float64 `db:"temprating,omitempty" json:"temprating,omitempty"`
TempUnit *string `db:"tempunit,omitempty" json:"tempunit,omitempty"`
Description1 *string `db:"description1,omitempty" json:"description1,omitempty"`
Description2 *string `db:"description2,omitempty" json:"description2,omitempty"`
Color1 *string `db:"color1,omitempty" json:"color1,omitempty"`
Color2 *string `db:"color2,omitempty" json:"color2,omitempty"`
Sourceinfo *string `db:"sourceinfo,omitempty" json:"sourceinfo,omitempty"`
Datasheet *string `db:"datasheet,omitempty" json:"datasheet,omitempty"`
Docs *string `db:"docs,omitempty" json:"docs,omitempty"`
Reference *string `db:"reference,omitempty" json:"reference,omitempty"`
Attributes *string `db:"attributes,omitempty" json:"attributes,omitempty"`
Condition *string `db:"condition,omitempty" json:"condition,omitempty"`
Note *string `db:"note,omit    category STRING,
empty" json:"note,omitempty"`
Warning  *string `db:"warning,omitempty" json:"warning,omitempty"`
Length *float64 `db:"length,omitempty,omitempty" json:"length,omitempty"`
Width *float64 `db:"width,omitempty" json:"width,omitempty"`
Height *float64 `db:"height,omitempty" json:"height,omitempty"`
WeightLb *float64 `db:"weightlb,omitempty" json:"weightlb,omitempty"`
WeightOz *float64 `db:"weightoz,omitempty" json:"weightoz,omitempty"`
MetaTitle *string `db:"metatitle,omitempty" json:"metatitle,omitempty"`
MetaDesc *string `db:"metadesc,omitempty" json:"metadesc,omitempty"`
MetaKeywords *string `db:"metakeywords,omitempty" json:"metakeywords,omitempty"`
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
    ID SERIAL PRIMARY KEY,
    image1 STRING,
    image2 STRING,
    image3 STRING,
    thumb STRING,
    name STRING,
    partno STRING,
    mfgpartno STRING,
    mfgname STRING,
    quantity INT,
    unlimitqty BOOL,
    enable BOOL,
    price FLOAT,
    msrp FLOAT,
    cost FLOAT,
    minorder INT,
    maxorder INT,
    location STRING,
		category STRING,
		subcategory STRING,
    type STRING,
    packagetype STRING,
    technology STRING,
		materials STRING,
    value FLOAT,
    valunit STRING,
    voltsrating FLOAT,
    ampsrating FLOAT,
		wattsrating FLOAT,
		temprating FLOAT,
		tempunit STRING,
    description1 STRING,
    description2 STRING,
		color1 STRING,
		color2 STRING,
    sourceinfo STRING,
    datasheet STRING,
    docs STRING,
    reference STRING,
    attributes STRING,
    condition STRING,
    note STRING,
    warning STRING,
    length FLOAT,
    width FLOAT,
    height FLOAT,
    weightlb FLOAT,
    weightoz FLOAT,
    metatitle STRING,
    metadesc STRING,
    metakeywords STRING
  )
  `)

if err != nil {
  return err
}
return nil
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
