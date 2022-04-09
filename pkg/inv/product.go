/*product.go*/
package inv

import (
"fmt"
  "github.com/upper/db/v4"

)

type Product struct {	// Product is used to represent a single record in the "products" table.
Id	int64	`db:"id,omitempty" json:"id,omitempty"`
Image1 string `db:"image1,omitempty" json:"image1,omitempty"`
Image2 string `db:"image2,omitempty" json:"image2,omitempty"`
Image3 string `db:"image3,omitempty" json:"image3,omitempty"`
Thumb string `db:"thumb,omitempty" json:"thumb,omitempty"`
PartNo string `db:"partno" json:"partno"`
Name	string `db:"name" json:"name"`
Enable  bool `db:"enable,omitempty" json:"enable,omitempty"`
Price float64  `db:"price" json:"price"`
Qty	int64  `db:"quantity" json:"quantity"`
MinOrder int64 `db:"minorder,omitempty" json:"minorder,omitempty"`
MaxOrder int64 `db:"maxorder,omitempty" json:"maxorder,omitempty"`
Shippable bool `db:"shippable,omitempty" json:"shippable,omitempty"`
UnlimitQty bool `db:"unlimitqty,omitempty" json:"unlimitqty,omitempty"`
DefaultQty 	int64  `db:"defaultquantity,omitempty" json:"defaultquantity,omitempty"`
StepQty 	int64  `db:"stepquantity,omitempty" json:"stepquantity,omitempty"`
MfgPartNo string `db:"mfgpartno,omitempty" json:"mfgpartno,omitempty"`
MfgName string `db:"mfgname,omitempty" json:"mfgname,omitempty"`
Category string `db:"category,omitempty" json:"category,omitempty"`
SubCategory string `db:"subcategory,omitempty" json:"subcategory,omitempty"`
Location string `db:"location,omitempty" json:"location,omitempty"`
Msrp float64  `db:"msrp,omitempty" json:"msrp,omitempty"`
Cost float64  `db:"cost,omitempty" json:"cost,omitempty"`
Type string  `db:"type,omitempty" json:"type,omitempty"`
PackageType string  `db:"packagetype,omitempty" json:"packagetype,omitempty"`
Technology string  `db:"technology,omitempty" json:"technology,omitempty"`
Materials string  `db:"materials,omitempty" json:"materials,omitempty"`
Value float64 `db:"value,omitempty" json:"value,omitempty"`
ValUnit string  `db:"valunit,omitempty" json:"valunit,omitempty"`
Resistance float64 `db:"resistance,omitempty" json:"resistance,omitempty"`
ResUnit string  `db:"resunit,omitempty" json:"resunit,omitempty"`
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
Year	 	int64  `db:"year,omitempty" json:"year,omitempty"`
Condition string `db:"condition,omitempty" json:"condition,omitempty"`
Note string `db:"note,omitempty" json:"note,omitempty"`
Warning  string `db:"warning,omitempty" json:"warning,omitempty"`
Length float64 `db:"length,omitempty,omitempty" json:"length,omitempty"`
CableLength float64 `db:"length,omitempty,omitempty" json:"length,omitempty"`
Width float64 `db:"width,omitempty" json:"width,omitempty"`
Height float64 `db:"height,omitempty" json:"height,omitempty"`
WeightLb float64 `db:"weightlb,omitempty" json:"weightlb,omitempty"`
WeightOz float64 `db:"weightoz,omitempty" json:"weightoz,omitempty"`
MetaTitle string `db:"metatitle,omitempty" json:"metatitle,omitempty"`
MetaDesc string `db:"metadesc,omitempty" json:"metadesc,omitempty"`
MetaKeywords string `db:"metakeywords,omitempty" json:"metakeywords,omitempty"`
}

type Category struct {
	Name string
	Count int
}

type SubCategory struct {
	Name string
	Count int
  Subcategories []Category
}


type Equipment = Product
/*
type Equipment struct {	// Equipment is used to represent a single record in the "equipments" table.
	Id	int64	`db:"id,omitempty" json:"id,omitempty"`
	Image1 string `db:"image1,omitempty" json:"image1,omitempty"`
	Image2 string `db:"image2,omitempty" json:"image2,omitempty"`
	Image3 string `db:"image3,omitempty" json:"image3,omitempty"`
	Thumb string `db:"thumb,omitempty" json:"thumb,omitempty"`
	PartNo string `db:"partno" json:"partno"`
	Name	string `db:"name" json:"name"`
	Enable  bool `db:"enable,omitempty" json:"enable,omitempty"`
	Price float64  `db:"price" json:"price"`
	Qty	int64  `db:"quantity" json:"quantity"`
	MinOrder int64 `db:"minorder,omitempty" json:"minorder,omitempty"`
	MaxOrder int64 `db:"maxorder,omitempty" json:"maxorder,omitempty"`
	Shippable bool `db:"shippable,omitempty" json:"shippable,omitempty"`
	UnlimitQty bool `db:"unlimitqty,omitempty" json:"unlimitqty,omitempty"`
	DefaultQty 	int64  `db:"defaultquantity,omitempty" json:"defaultquantity,omitempty"`
	StepQty 	int64  `db:"stepquantity,omitempty" json:"stepquantity,omitempty"`
	MfgPartNo string `db:"mfgpartno,omitempty" json:"mfgpartno,omitempty"`
	MfgName string `db:"mfgname,omitempty" json:"mfgname,omitempty"`
	Category string `db:"category,omitempty" json:"category,omitempty"`
	SubCategory string `db:"subcategory,omitempty" json:"subcategory,omitempty"`
	Location string `db:"location,omitempty" json:"location,omitempty"`
	Msrp float64  `db:"msrp,omitempty" json:"msrp,omitempty"`
	Cost float64  `db:"cost,omitempty" json:"cost,omitempty"`
	Type string  `db:"type,omitempty" json:"type,omitempty"`
	PackageType string  `db:"packagetype,omitempty" json:"packagetype,omitempty"`
	Technology string  `db:"technology,omitempty" json:"technology,omitempty"`
	Materials string  `db:"materials,omitempty" json:"materials,omitempty"`
	Value float64 `db:"value,omitempty" json:"value,omitempty"`
	ValUnit string  `db:"valunit,omitempty" json:"valunit,omitempty"`
	Resistance float64 `db:"resistance,omitempty" json:"resistance,omitempty"`
	ResUnit string  `db:"resunit,omitempty" json:"resunit,omitempty"`
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
	Year	 	int64  `db:"year,omitempty" json:"year,omitempty"`
	Condition string `db:"condition,omitempty" json:"condition,omitempty"`
	Note string `db:"note,omitempty" json:"note,omitempty"`
	Warning  string `db:"warning,omitempty" json:"warning,omitempty"`
	CableLength float64 `db:"length,omitempty,omitempty" json:"length,omitempty"`
	Length float64 `db:"length,omitempty,omitempty" json:"length,omitempty"`
	Width float64 `db:"width,omitempty" json:"width,omitempty"`
	Height float64 `db:"height,omitempty" json:"height,omitempty"`
	WeightLb float64 `db:"weightlb,omitempty" json:"weightlb,omitempty"`
	WeightOz float64 `db:"weightoz,omitempty" json:"weightoz,omitempty"`
	MetaTitle string `db:"metatitle,omitempty" json:"metatitle,omitempty"`
	MetaDesc string `db:"metadesc,omitempty" json:"metadesc,omitempty"`
	MetaKeywords string `db:"metakeywords,omitempty" json:"metakeywords,omitempty"`
}

*/
//id SERIAL PRIMARY KEY,
//id INT8 PRIMARY KEY DEFAULT unique_rowid(),
func CreateProductsTableIfNotExists(sess db.Session) error {	// createTables creates all the tables that are neccessary to run this example.
	fmt.Printf("Creating 'products' table\n")
	_, err := sess.SQL().Exec(`
	  CREATE TABLE IF NOT EXISTS products (
	    id SERIAL PRIMARY KEY,
			image1 STRING NULL DEFAULT '',
	    image2 STRING NULL DEFAULT '',
	    image3 STRING NULL DEFAULT '',
	    thumb STRING NULL DEFAULT '',
			partno STRING UNIQUE NOT NULL,
	    name STRING NOT NULL,
			enable BOOL DEFAULT FALSE,
			price FLOAT DEFAULT 0.00,
			quantity INT DEFAULT 0,
			shippable BOOL DEFAULT TRUE,
			minorder INT DEFAULT 1,
			maxorder INT DEFAULT 100,
			unlimitqty BOOL DEFAULT FALSE,
			defaultquantity INT DEFAULT 1,
			stepquantity INT DEFAULT 1,
			mfgpartno STRING NULL DEFAULT '',
			mfgname STRING NULL DEFAULT '',
			category STRING NULL DEFAULT '',
			subcategory STRING NULL DEFAULT '',
			location STRING NULL DEFAULT '',
	    msrp FLOAT DEFAULT 0.00,
	    cost FLOAT DEFAULT 0.00,
	    type STRING NULL DEFAULT '',
	    packagetype STRING NULL DEFAULT '',
	    technology STRING NULL DEFAULT '',
			materials STRING NULL DEFAULT '',
			value FLOAT DEFAULT 0.0,
	    valunit STRING NULL DEFAULT '',
			resistance FLOAT DEFAULT 0.0,
	    resunit STRING NULL DEFAULT 'Ω',
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
			year INT DEFAULT 0,
	    condition STRING NULL DEFAULT '',
	    note STRING NULL DEFAULT '',
	    warning STRING NULL DEFAULT '',
			cablelength FLOAT DEFAULT 0.0,
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
	func CreateEquipmentsTableIfNotExists(sess db.Session) error {	// createTables creates all the tables that are neccessary to run this example.
	fmt.Printf("Creating 'equipments' table\n")
	_, err := sess.SQL().Exec(`
	  CREATE TABLE IF NOT EXISTS equipments (
			id SERIAL PRIMARY KEY,
			image1 STRING NULL DEFAULT '',
	    image2 STRING NULL DEFAULT '',
	    image3 STRING NULL DEFAULT '',
	    thumb STRING NULL DEFAULT '',
			partno STRING UNIQUE NOT NULL,
	    name STRING NOT NULL,
			enable BOOL DEFAULT FALSE,
			price FLOAT DEFAULT 0.00,
			quantity INT DEFAULT 0,
			shippable BOOL DEFAULT TRUE,
			minorder INT DEFAULT 1,
			maxorder INT DEFAULT 100,
			unlimitqty BOOL DEFAULT FALSE,
			defaultquantity INT DEFAULT 1,
			stepquantity INT DEFAULT 1,
			mfgpartno STRING NULL DEFAULT '',
			mfgname STRING NULL DEFAULT '',
			category STRING NULL DEFAULT '',
			subcategory STRING NULL DEFAULT '',
			location STRING NULL DEFAULT '',
	    msrp FLOAT DEFAULT 0.00,
	    cost FLOAT DEFAULT 0.00,
	    type STRING NULL DEFAULT '',
	    packagetype STRING NULL DEFAULT '',
	    technology STRING NULL DEFAULT '',
			materials STRING NULL DEFAULT '',
			value FLOAT DEFAULT 0.0,
	    valunit STRING NULL DEFAULT '',
			resistance FLOAT DEFAULT 0.0,
	    resunit STRING NULL DEFAULT 'Ω',
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
			year INT DEFAULT 0,
	    condition STRING NULL DEFAULT '',
	    note STRING NULL DEFAULT '',
	    warning STRING NULL DEFAULT '',
			cablelength FLOAT DEFAULT 0.0,
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
