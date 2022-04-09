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

// /* defines Mproducts and categories from sess / fetch from database  */ //
func DefineProducts(sess db.Session) {
	//define Mproducts
	MproductsCol := Products(sess)
	Mproducts = []Product{}
	err = MproductsCol.Find("enable", true).All(&Mproducts) 	// Find().All() maps all the records from the Mproducts collection.
	if err != nil {
		log.Fatal("productsCol.Find: ", err)
	}

	Lenproducts = len(Mproducts)
	fmt.Printf("%d Products\n", Lenproducts)
	//now for the hard part
	//define categories
	//var cat1 []string
	var cat2 []Category
	//range the products list and parse the unique categories
	//also track the subcategories
	for i:= range Mproducts{
		//valid category
		if Mproducts[i].Category != "" {
			//fill in first one if empty
			if Pcats == nil || len(Pcats)==0{
				if Mproducts[i].SubCategory != "" {
					//this represents the subcategory
					cat2 = append(cat2, Category{ Mproducts[i].SubCategory, 1})
				}
				//put subcategory (cat2) at the end of the category and append to pcats
				Pcats = append(Pcats, SubCategory{Mproducts[i].Category, 1, cat2})
				//reset the variable
				cat2 = []Category{}
		} else {
				//Pcats not empty
				founded:=false
				//iterate over the Categories
				for j:= range Pcats{
					//category exists
					if Pcats[j].Name == Mproducts[i].Category {
						//variable used later
						founded=true
						//increment count
						Pcats[j].Count += 1
						//check that subcategory is not blank
						if Mproducts[i].SubCategory != "" {
							//check that there is an array
							if Pcats[j].Subcategories == nil || len(Pcats[j].Subcategories) == 0 {
								//add first element if none are found
								Pcats[j].Subcategories = append(Pcats[j].Subcategories, Category{Name: Mproducts[i].SubCategory, Count: 1})
							} else {
							founded1:=false

							for k := range Pcats[j].Subcategories {

								if Pcats[j].Subcategories[k].Name == Mproducts[i].SubCategory {
									founded1=true
									Pcats[j].Subcategories[k].Count += 1
								}
							}
								if !founded1{
									Pcats[j].Subcategories = append(Pcats[j].Subcategories, Category{ Mproducts[i].SubCategory, 1})
								}
							}
						}
					}
				}
				if !founded{
					if Mproducts[i].SubCategory != "" {
						//this represents the subcategory
						cat2 = append(cat2, Category{ Mproducts[i].SubCategory, 1})
					}
					Pcats = append(Pcats, SubCategory{Mproducts[i].Category, 1, cat2})
					//Pcats = append(Pcats, SubCategory{Name: Mproducts[i].Category, Count: 1, Subcategories: Category{Name: Mproducts[i].SubCategory, Count: 1}})
					cat2 = []Category{}
				}
			}
		}
		sort.Sort(alphab(Pcats))
	}
}
//DistProdCats(cat1)
//countProdCats(cat1)
//}

// /* defines Mproducts and categories from sess / fetch from database  */ //
/*
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
*/
var Pcats []SubCategory
func RetPcats() ([]SubCategory) {
return Pcats
}
var Psubcats []Category
func RetPsubcats(category string) ([]Category) {
Psubcats = []Category{}
	for i:= range Pcats {
		if Pcats[i].Name == category {
			Psubcats = Pcats[i].Subcategories

	}
}

return Psubcats
}

/*
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
*/

type alphab []SubCategory

func (cat alphab) Len() int { return len(cat) }
func (cat alphab) Less(i, j int) bool { return cat[i].Name < cat[j].Name }
func (cat alphab) Swap(i, j int) { cat[i], cat[j] = cat[j], cat[i] }

var Ecats []SubCategory
func RetEcats() ([]SubCategory) {
return Pcats
}
var Esubcats SubCategory
func RetEsubcats(category string) ([]Category) {
return Psubcats
}

/*
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
*/

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
