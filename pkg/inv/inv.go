/*inv.go*/

package inv

import(
	"log"
	"fmt"
	"sort"
	"github.com/upper/db/v4"
	//"github.com/upper/db/v4/adapter/cockroachdb"

	//mdb "github.com/0pcom/magnets/pkg/db"
	//user "github.com/0pcom/magnets/pkg/user"
	"strconv"

)

var err error
var Mproducts []Product
var ppartno Product
var epartno Equipment
var category string
var Lenproducts int
var Mequipments []Equipment
var lenequipments int

// /* from upper/db tutorial -OR- cockroachdb documentation for upper/db */ //
func Products(sess db.Session) db.Store {	// Products is a handy way to represent a collection.
return sess.Collection("products")
}
func Equipments(sess db.Session) db.Store {	// Equipments is a handy way to represent a collection.
return sess.Collection("equipments")
}


func (a *Product) Store(sess db.Session) db.Store {// Collection is required in order to create a relation between the Product struct and the "products" table.
return Products(sess)
}
//func (a *Equipment) Store(sess db.Session) db.Store {// Collection is required in order to create a relation between the Equipment struct and the "equipments" table.
//return Equipments(sess)
//}

//func RetSess() (sess db.Session){
//settings := user.FetchSettings()
//sess = mdb.Connect(settings)
//return sess
//}


func DropProductsTable(sess db.Session) error {
	fmt.Printf("Dropping 'products' table\n")
	_, err := sess.SQL().Exec(`
		DROP TABLE product.products
		`)
	if err != nil {
		return err
	}
	return nil
}

func DropEquipmentsTable(sess db.Session) error {
	fmt.Printf("Dropping 'equipments' table\n")
	_, err := sess.SQL().Exec(`
		DROP TABLE product.equipments
		`)
	if err != nil {
		return err
	}
	return nil
}


// /* defines Mproducts and categories from sess / fetch from database  */ //
func DefineProducts(sess db.Session) {
	//define Mproducts
	MproductsCol := Products(sess)
	Mproducts = []Product{}
	err = MproductsCol.Find("enable", true).All(&Mproducts) 	// Find().All() maps all the records from the Mproducts collection.
	if err != nil {
		log.Fatal("productsCol.Find: ", err)
	}
	//count Mproducts
//	for i := 0; i < len(Mproducts); i++ {
//		if Mproducts[i].Enable == true {
//			Lenproducts = Lenproducts + 1
//	}
//}

Lenproducts = len(Mproducts)
			fmt.Printf("%d Products\n", Lenproducts)

	//define categories
	var cat1 []string
	for i:= range Mproducts{
		if Mproducts[i].Category != "" {
			cat1 = append(cat1, Mproducts[i].Category)
	}
}
DistProdCats(cat1)
//countProdCats(cat1)
}
// /* defines Mproducts and categories from sess / fetch from database  */ //
func DefineEquipments(sess db.Session) {
	//define Mequipments
	MequipmentsCol := Equipments(sess)
	Mequipments = []Product{}
	err = MequipmentsCol.Find().All(&Mequipments) 	// Find().All() maps all the records from the Mproducts collection.
	if err != nil {
		log.Fatal("equipmentsCol.Find: ", err)
	}
	//define categories
	var cat1 []string
	for i:= range Mequipments{
		if Mequipments[i].Category != "" {
	cat1 = append(cat1, Mequipments[i].Category)
	}
}
DistEquipCats(cat1)
}

var Pcats []Category
func RetPcats() ([]Category) {
return Pcats
}


func DistProdCats(cat1 []string){
    for i:= range cat1{
        if Pcats == nil || len(Pcats)==0{ Pcats = append(Pcats, Category{Name: cat1[i], Count: 1}) } else {
            founded:=false
            for j:= range Pcats{
							if Pcats[j].Name == cat1[i] {
								founded=true
								Pcats[j].Count += 1
								}
							}
						if !founded{
							Pcats = append(Pcats, Category{Name: cat1[i], Count: 1})
							}
        }
    }
		sort.Sort(alphab(Pcats))
}


type alphab []Category

func (cat alphab) Len() int { return len(cat) }
func (cat alphab) Less(i, j int) bool { return cat[i].Name < cat[j].Name }
func (cat alphab) Swap(i, j int) { cat[i], cat[j] = cat[j], cat[i] }

var Ecats []Category
func RetEcats() ([]Category) {
return Pcats
}
func DistEquipCats(cat1 []string){
	for i:= range cat1{
			if Ecats == nil || len(Ecats)==0{ Ecats = append(Ecats, Category{Name: cat1[i], Count: 1}) } else {
					founded:=false
					for j:= range Ecats{
						if Ecats[j].Name == cat1[i] {
							founded=true
							Ecats[j].Count += 1
							}
						}
					if !founded{
						Ecats = append(Ecats, Category{Name: cat1[i], Count: 1})
						}
			}
	}
	sort.Sort(alphab(Ecats))
}


///*	database test stuff	*///
func DeleteAllProducts(sess db.Session) {
fmt.Printf("Clearing products table\n")
//clear tables ; testing
err := Products(sess).Truncate()
if err != nil {
  log.Fatal("Truncate: ", err)
}
}
func DeleteAllEquipments(sess db.Session) {
fmt.Printf("Clearing equipments table\n")
//clear tables ; testing
err := Equipments(sess).Truncate()
if err != nil {
  log.Fatal("Truncate: ", err)
}
}
//creates a single test product
func CreateTestProd(sess db.Session) {
fmt.Printf("Creating test product 'dummy'\n")
_desc := "test entry to database"
_img := "test.jpg"
product1 := Product{Name: "dummy", PartNo:"test", Description1: _desc, Price:1.00, Image1: _img, Qty: 10}
err := Products(sess).InsertReturning(&product1)
if err != nil {
    log.Fatal("sess.Save: ", err)
}
fmt.Printf("Creating second test product 'dummy2'\n")
product1 = Product{Name: "dummy2", PartNo:"test1", Description1: _desc, Price: 1.00, Qty: 100}
err = Products(sess).InsertReturning(&product1)
if err != nil {
    log.Fatal("sess.Save: ", err)
}
}
func CreateTestEquip(sess db.Session) {
fmt.Printf("Creating test equipment 'dummy'\n")
_desc := "test entry to database"
_img := "test.jpg"
equipment1 := Equipment{Name: "dummy", PartNo:"test", Description1: _desc, Price:100.00, Image1: _img, Qty: 1}
err := Equipments(sess).InsertReturning(&equipment1)
if err != nil {
    log.Fatal("sess.Save: ", err)
}
fmt.Printf("Creating second test equipment 'dummy2'\n")
equipment1 = Equipment{Name: "dummy2", PartNo:"test1", Description1: _desc, Price:100.00, Image1: _img, Qty: 1}
err = Equipments(sess).InsertReturning(&equipment1)
if err != nil {
    log.Fatal("sess.Save: ", err)
}
}

func CreateProduct(sess db.Session, Createpartno string) {
fmt.Printf("Creating product with part number: %s'\n", Createpartno)
product1 := Product{PartNo: Createpartno, Enable: false}
err := Products(sess).InsertReturning(&product1)
if err != nil {
    log.Fatal("sess.Save: ", err)
}
}
//create a series of products with sequential part numbers
func CreateSeries(sess db.Session, table string, series int) {
	if table == "products" {
for i := 0; i < series; i++ {
	    p := strconv.Itoa(i)
			fmt.Printf("Creating product part number: %s\n", p)
			p1 := Product{PartNo: p, Enable: false}
			err := Products(sess).InsertReturning(&p1)
			if err != nil {
			    log.Fatal("sess.Save: ", err)
			}
}
}
if table == "equipments" {
	for i := 0; i < series; i++ {
		    p := strconv.Itoa(i)
				fmt.Printf("Creating product part number: %s\n", p)
				p1 := Product{PartNo: p, Enable: false}
				err := Equipments(sess).InsertReturning(&p1)
				if err != nil {
				    log.Fatal("sess.Save: ", err)
				}
	}
}
}
