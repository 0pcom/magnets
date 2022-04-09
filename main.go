// /* */ //
// /* main.go */ //
package main

import (
 	"fmt"
	"log"
	"net/http"
	"html/template"
	"time"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/cockroachdb"
	"math"
	"os/exec"
	"os"
	"strconv"
	"sort"
	"math/rand"
	_ "github.com/0pcom/magnets/statik"
  //"github.com/0pcom/magnets/statik"
	"github.com/rakyll/statik/fs"
  //"encoding/json"
flags "github.com/spf13/pflag"
)

//var baseURL = "https://magnetosphere.net"
var port = 8040
//var createpartno string

	var (
		runapp	bool
		createtables	bool
		deleteproducts	bool
		deleteequipments	bool
		droptable	bool
		droptable1	bool
		testprod	bool
		testequip	bool
		createpartno	string
		create100p	bool
		create100e	bool
		create5000p	bool
		create5000e	bool
		importcsv	bool
		importcsv1	bool
		printinv	bool
		printinv1	bool
		vprintinv	bool
		exportcsv	bool
		exportcsv1	bool
		helpmenu	bool
	)

	func main(){
		//TO DO: implement more string arguments for import export and port
		flags.BoolVarP(&droptable, "droptables", "d", false, "Drop products table")
		flags.BoolVarP(&droptable1, "droptables1", "b", false, "Drop equipments table")
		flags.BoolVarP(&deleteproducts, "deleteall", "D", false, "Delete the products in the products database")
		flags.BoolVarP(&deleteequipments, "deleteall1", "E", false, "Delete the equipment in the equipments database")
		flags.BoolVarP(&createtables, "createtables", "c", false, "Create the tables if they do not exist")
		flags.BoolVarP(&testprod, "testprod", "t", false, "create test product")
		flags.BoolVarP(&testequip, "testequip", "u", false, "create test equipment")
		flags.StringVarP(&createpartno, "createpartno", "C", "", "Create a part by providing the part number")
		flags.BoolVarP(&create5000p, "create5000p", "Y", false, "Create 5000 products with sequential part numbers")
		flags.BoolVarP(&create5000e, "create5000e", "Z", false, "Create 5000 equipments with sequential part numbers")
		flags.BoolVarP(&create100p, "create100p", "y", false, "Create 100 products with sequential part numbers")
		flags.BoolVarP(&create100e, "create100e", "z", false, "Create 100 equipments with sequential part numbers")
		flags.BoolVarP(&exportcsv, "exportcsv", "e", false, "Export products to productsexport01.csv")
		flags.BoolVarP(&exportcsv1, "exportcsv1", "f", false, "Export equipments to equipmentsexport01.csv")
		flags.BoolVarP(&printinv, "printinv", "p", false, "Print the products table to the terminal")
		flags.BoolVarP(&printinv1, "printinv1", "q", false, "Print the equipments table to the terminal")
		//flags.BoolVarP(&printinv2, "printinv2", "q", false, "Print the  inventory - BTCPayserver format")
		flags.BoolVarP(&vprintinv, "vprintinv", "v", false, "More verbose printinventory")
		flags.BoolVarP(&importcsv, "importcsv", "i", false, "Import products csv from http://127.0.0.1:8079/export01.csv")
		flags.BoolVarP(&importcsv1, "importcsv1", "j", false, "Import equipments csv from http://127.0.0.1:8079/export01.csv")
		flags.BoolVarP(&runapp, "run", "r", false, "run the web app")
		flags.BoolVarP(&helpmenu, "help", "h", false, "show this help menu")
		flags.Parse()

var cmdargs int = 0
// /* help menu */ //
if helpmenu {
	helpmenu1()
	os.Exit(0)
}
// /* database connection */ //
	fmt.Printf("Initializing cockroachDB connection\n")
	sess, err := cockroachdb.Open(settings)		//establish the session
	if err != nil {
		log.Fatal("cockroachdb.Open: ", err)
	}
	defer sess.Close()
	// /* drop tables */ //
	if droptable {
			dropProductsTable(sess)
			cmdargs = 1
		}
	// /* drop tables */ //
	if droptable1 {
			dropEquipmentsTable(sess)
			cmdargs = 1
		}
	// /* create tables */ //
	if createtables {
			createProductsTableIfNotExists(sess)
			createEquipmentsTableIfNotExists(sess)
			cmdargs = 1
	}
	// /* delete products */ //
	if deleteproducts {
		defineproducts(sess)
		deleteAllProducts(sess)
		cmdargs = 1
	}
	// /* delete products */ //
	if deleteequipments {
		defineequipments(sess)
		deleteAllEquipments(sess)
		cmdargs = 1
	}
	// /* Create Test Product */ //
	if testprod {
		defineproducts(sess)
			createTestProd(sess)
			cmdargs = 1
	}
	// /* Create Test Product */ //
	if testequip {
		defineequipments(sess)
			createTestEquip(sess)
			cmdargs = 1
	}
	// /* create part */ //
	if createpartno != "" {
		fmt.Println("createpart has value ", createpartno)
		createProduct(sess, createpartno)
		cmdargs = 1
	}
	// /* Create 100 Product */ //
	if create5000p {
		fmt.Println("creating 5000 products")
		create5kP(sess)
		cmdargs = 1
	}
	// /* Create 100 Equipment */ //
	if create5000e {
		fmt.Println("creating 5000 equipments")
		create5kE(sess)
		cmdargs = 1
	}
	if create100p {
		fmt.Println("creating 100 products")
		createHundredP(sess)
		cmdargs = 1
	}
	// /* Create 100 Equipment */ //
	if create100e {
		fmt.Println("creating 100 equipments")
		createHundredE(sess)
		cmdargs = 1
	}
	// /* print products table inventory */ //
	if printinv {
		defineproducts(sess)
		log.Printf("products:")
		for i := range products {
			//fmt.Printf("product #%d: %#v\n", i, products[i])
				fmt.Printf("product[%d]:\n", products[i].Id)
				fmt.Printf("\tpartno:		");	fmt.Printf("%s\n", products[i].PartNo)
				if products[i].Image1 != "" {	fmt.Printf("\tImage1:		");	fmt.Printf("%s\n", products[i].Image1)	}
				if products[i].Name != "" {	fmt.Printf("\tName:		");	fmt.Printf("%s\n", products[i].Name)	}
				fmt.Printf("\tQty:		"); fmt.Printf("%d\n", products[i].Qty)
				fmt.Printf("\tPrice:		"); fmt.Printf("%.2f\n", products[i].Price)
				fmt.Printf("\tEnable:		"); fmt.Printf("%t\n", products[i].Enable)
		}
		cmdargs = 1
	}
	// /* print equipments table inventory*/ //
	if printinv1 {
		defineequipments(sess)
		log.Printf("equipments:")
		for i := range equipments {
			//fmt.Printf("product #%d: %#v\n", i, products[i])
				fmt.Printf("equipment[%d]:\n", equipments[i].Id)
				fmt.Printf("\tpartno:		");	fmt.Printf("%s\n", equipments[i].PartNo)
				if equipments[i].Image1 != "" {	fmt.Printf("\tImage1:		");	fmt.Printf("%s\n", equipments[i].Image1)	}
				if equipments[i].Name != "" {	fmt.Printf("\tName:		");	fmt.Printf("%s\n", equipments[i].Name)	}
				fmt.Printf("\tQty:		"); fmt.Printf("%d\n", equipments[i].Qty)
				fmt.Printf("\tPrice:		"); fmt.Printf("%.2f\n", equipments[i].Price)
				fmt.Printf("\tEnable:		"); fmt.Printf("%t\n", equipments[i].Enable)
		}
		cmdargs = 1
	}
	// /* print inventory so it can be pasted in the btcpayserver raw editor! */ //
	/*
	if printinv2 {
		defineproducts(sess)
		log.Printf("products:")
		for i := range products {
			if products[i].Enable {
				fmt.Printf("%s:\n", products[i].PartNo)
				fmt.Printf("  price: "); fmt.Printf("%.2f\n", products[i].Price)
				fmt.Printf("  title:	");	fmt.Printf("%s\n", products[i].Name)
				fmt.Printf("  description: "); fmt.Printf("%s\n", products[i].Description1)
				fmt.Printf("  image:	");	fmt.Printf("%s/img/%s\n", siteURL, products[i].Image1)
				}
			}
	cmdargs = 1
		}
		*/
	// /* more verbosely print inventory */ //
	if vprintinv {
		defineproducts(sess)
		log.Printf("products:")
		for i := range products {
			fmt.Printf("product #%d: %#v\n", i, products[i])
			cmdargs = 1
				}
			}
			// /* export products table to csv RELIES ON MAKEFILE - COCKROACHDB COMMANDS */ //
			if exportcsv {
				fmt.Println("Exporting csv to productsexport01.csv")
				exportProductsCSV()
				cmdargs = 1
			}
			// /* export equipments table csv RELIES ON MAKEFILE - COCKROACHDB COMMANDS */ //
			if exportcsv1 {
				fmt.Println("Exporting csv to productsexport01.csv")
				exportEquipmentsCSV()
				cmdargs = 1
			}
			// /* import csv to products table */ //
			if importcsv {
					fmt.Println("Import a csv into products table from http://127.0.0.1:8079/productsexport01.csv")
					importProductsCSV(sess)
					cmdargs = 1
			}
			// /* import csv to equipments table */ //
			if importcsv1 {
					fmt.Println("Import a csv into equipments table from http://127.0.0.1:8079/equipmentsexport01.csv")
					importEquipmentsCSV(sess)
					cmdargs = 1
			}
	// /* run the web app */ //
	if runapp {
		statikFS, err := fs.New()
		if err != nil {
			log.Fatal(err)
		}
		defineproducts(sess)
		defineequipments(sess)
		r := mux.NewRouter() //.StrictSlash(true)
    r.NotFoundHandler = page404()
		r.PathPrefix("/img/").Handler(http.StripPrefix("/img/", http.FileServer(http.Dir("./img")))) //images
		//products table - main site original endpoints
		r.Handle("/list", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(listPage))).Methods("GET") //pagination
		r.Handle("/list/", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(listPage))).Methods("GET") //pagination
		r.Handle("/list/{id:[0-9]+}", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(listPage))).Methods("GET") //pagination
		r.Handle("/post/{slug}", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(findProduct))).Methods("GET")	//individual product page
		r.Handle("/cat/{slug}", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(categoryPage))).Methods("GET")	//category
		r.Handle("/cat/{slug}/{id:[0-9]+}", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(categoryPage))).Methods("GET")	//category pagination
		// equipment table endpoints
		r.Handle("/equipment", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(equipmentPage))).Methods("GET")	//equipment
		r.Handle("/equipment/p/{id:[0-9]+}", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(equipmentPage))).Methods("GET")	//equipment pagination
		r.Handle("/equipment/post/{slug}", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(findEquipment))).Methods("GET")	//individual product page
		r.Handle("/equipment/cat/{slug}", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(categoryEquipmentPage))).Methods("GET")	//category
		r.Handle("/equipment/cat/{slug}/{id:[0-9]+}", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(categoryEquipmentPage))).Methods("GET")	//category pagination
		//single pages
		r.Handle("/about", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(aboutPage))).Methods("GET")	//about page
		r.Handle("/blog", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(blogPage))).Methods("GET")	//about page
		r.Handle("/blog/", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(blogPage))).Methods("GET")	//about page
		r.Handle("/blog/{slug}/", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(blogPage))).Methods("GET")	//about page
		r.Handle("/friend", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(friendPage))).Methods("GET")	//friends page
		r.Handle("/friend/", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(friendPage))).Methods("GET")	//friends page
		r.Handle("/policy", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(policyPage))).Methods("GET")	//shipping page
		r.Handle("/policy/", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(policyPage))).Methods("GET")	//shipping page
		//r.Handle("/time", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(timeFunc))).Methods("GET")	//shipping page
		//
    r.Handle("/{slug}", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(frontPage))).Methods("GET") //site root
    r.Handle("/", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(frontPage))).Methods("GET") //site root
    r.Handle("/p/{id:[0-9]+}", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(pagePage))).Methods("GET") //pagination
		r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(statikFS)))	//statik sources
    r.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
      // this handler behaved as I would expect the r.NotFoundHandler to behave..
      w.WriteHeader(501)
      w.Write([]byte(`{"status":501,"message":"501: Not implemented."}`))
    })
		Serve = r
		fmt.Printf("listening on http://127.0.0.1:%d using gorilla router\n", port)
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), Serve))
		cmdargs = 1
	}
if cmdargs == 0 { helpmenu1() }
}


// custom 404 not found page
func page404() http.Handler {
return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
  wd, err := os.Getwd()
	if err != nil { log.Fatal(err)	}
	tpl1 := template.Must(template.New("").Funcs(fm).ParseFiles(wd + "/public/404.html"))
	if err :=	tpl1.ExecuteTemplate(w, "404.html", nil); err != nil {	fmt.Printf("error: %s", err) }
})
}
// custom 404 not found page
func page404Page (w http.ResponseWriter, r *http.Request) {
  wd, err := os.Getwd()
	if err != nil { log.Fatal(err)	}
	tpl1 := template.Must(template.ParseFiles(wd + "/public/404.html"))
	if err :=	tpl1.ExecuteTemplate(w, "404.html", nil); err != nil {	fmt.Printf("error: %s", err) }
}

func helpmenu1() {
	fmt.Printf("Usage: magnets -dDctCyepirh\n")
	fmt.Printf("\tSuggested Demo: magnets -ctpr\n")
	flags.PrintDefaults()
}

// /* defines products and categories from sess / fetch from database  */ //
func defineproducts(sess db.Session) {
	//define products
	productsCol := Products(sess)
	products = []Product{}
	err = productsCol.Find().All(&products) 	// Find().All() maps all the records from the products collection.
	if err != nil {
		log.Fatal("productsCol.Find: ", err)
	}
	//count products
	for i := 0; i < len(products); i++ {
		if products[i].Enable == true {
			lenproducts = lenproducts + 1
	}
}
fmt.Printf("%d Products\n", lenproducts)

	//define categories
	var cat1 []string
	for i:= range products{
		if products[i].Category != "" {
			cat1 = append(cat1, products[i].Category)
	}
}
distProdCats(cat1)
//countProdCats(cat1)
}
// /* defines products and categories from sess / fetch from database  */ //
func defineequipments(sess db.Session) {
	//define equipments
	equipmentsCol := Equipments(sess)
	equipments = []Equipment{}
	err = equipmentsCol.Find().All(&equipments) 	// Find().All() maps all the records from the products collection.
	if err != nil {
		log.Fatal("equipmentsCol.Find: ", err)
	}
	//define categories
	var cat1 []string
	for i:= range equipments{
		if equipments[i].Category != "" {
	cat1 = append(cat1, equipments[i].Category)
	}
}
distEquipCats(cat1)
}

var pcats []Category
func distProdCats(cat1 []string){
    for i:= range cat1{
        if pcats == nil || len(pcats)==0{ pcats = append(pcats, Category{Name: cat1[i], Count: 1}) } else {
            founded:=false
            for j:= range pcats{
							if pcats[j].Name == cat1[i] {
								founded=true
								pcats[j].Count += 1
								}
							}
						if !founded{
							pcats = append(pcats, Category{Name: cat1[i], Count: 1})
							}
        }
    }
		sort.Sort(alphab(pcats))
}


type alphab []Category

func (cat alphab) Len() int { return len(cat) }
func (cat alphab) Less(i, j int) bool { return cat[i].Name < cat[j].Name }
func (cat alphab) Swap(i, j int) { cat[i], cat[j] = cat[j], cat[i] }

var ecats []Category
func distEquipCats(cat1 []string){
	for i:= range cat1{
			if ecats == nil || len(ecats)==0{ ecats = append(ecats, Category{Name: cat1[i], Count: 1}) } else {
					founded:=false
					for j:= range ecats{
						if ecats[j].Name == cat1[i] {
							founded=true
							ecats[j].Count += 1
							}
						}
					if !founded{
						ecats = append(ecats, Category{Name: cat1[i], Count: 1})
						}
			}
	}
	sort.Sort(alphab(ecats))
}

//package gorilla	// Routing based on the gorilla/mux router
var Serve http.Handler
var ppartno Product
var epartno Equipment
var category string
var products []Product
var lenproducts int
var equipments []Equipment
var lenequipments int
//these are functions called from the template or webpage
var fm = template.FuncMap{
	"fdateMDY": monthDayYear,
	"snipcartapikey": snipcartApiKey,
	"multiply": multiply,
	"correct": correct,
	"convertozgrams": convertozgrams,
	"convertincm": convertincm,
	"lenprods": lenprods,
	"productListPage": productListPage,
	"listPage1": listPage1,
	"productIndexPage": productIndexPage,
	"indexPage1": indexPage1,
	"equipmentListPage": equipmentListPage,
	"nextPage": nextPage,
	"prevPage": prevPage,
	"listCategories": listCategories,
	"productsCategoryListPage": productsCategoryListPage,
	"equipmentsCategoryListPage": equipmentsCategoryListPage,
	"findProduct1": findProduct1,
	"findEquipment1": findEquipment1,
	//"pagination": pagination,
}//below are the functions
// /* timepage  */ //
//func monthDayYear(t time.Time) string {
func monthDayYear() string {
	return time.Now().Format("Monday January 2, 2006 15:04:05")
}
func multiply(a float64, b float64) float64 { //generic multiply
return a * b
}
func correct(a float64) string {	//correct the number of decimals
return fmt.Sprintf("%.2f", a)
}
func convertozgrams(a float64) float64 {	//convert ounces to grams for snipcart
	return math.Round((a*28.35)*100)/100
}
func convertincm(a float64) float64 {	//convert inches to centimeters for snipcart
	return math.Round((a*2.54)*100)/100
}
//func timeFunc(w http.ResponseWriter, r *http.Request) {	//a page with just the time
//	w.Header().Set("Content-Type", "text/html")
//	tp1 := template.Must(template.New("").Funcs(fm).ParseFiles("time.html"))
//	if err := tp1.ExecuteTemplate(w, "time.html",nil); err != nil {	fmt.Printf("error: %s", err) }
//}
// /* return list of categories */ //
func listCategories(table string) []Category {
	var toreturn []Category
	if table == "products" {
		toreturn = pcats
	}
	if table == "equipments" {
		toreturn = ecats
	}
	return toreturn
}
func nextPage(a int) int {//the "current context" of the template ({{.}}) is now an index
	return a + 1
}
func prevPage(a int) int {//check for nonzero argument is done on the template side
	return a - 1
}
func nextProd(a int) int {//the "current context" of the template ({{.}}) is now an index
	return a + 1
}
func prevProd(a int) int {//check for nonzero argument is done on the template side
	return a - 1
}

var snipcartapikey string = os.Getenv("SNIPCARTAPIKEY")
func snipcartApiKey() string {
	if snipcartapikey == "" {
		log.Fatal("error no api key found")
	}
	return snipcartapikey
}
//these structs are passed to the template by the handler Funcs
//like an index, providing context to the template
//the template then calls the mapped functions, passing the values that were passed to it
//returned from there is the data which is rendered on the page
type Page struct {
	Table string //specifies the table; products or equipments
  Category string //category
  PageNumber int	//pagination
}
type SubPage struct {
	Table string
  PartNumber string
}
// /* individual product page ENDPOINT: /post/{slug} */ //
func findProduct(w http.ResponseWriter, r *http.Request) {	//, product string
	slug := mux.Vars(r)["slug"]
	a := false
	for i := range products {
		if products[i].PartNo == slug {
			a = true
			break				// Found!
		}
}
if !a {
	fmt.Fprint(w, "No product found for part number:\n", slug)
} else {
	productp := SubPage{"products", slug}
	wd, err := os.Getwd()
	if err != nil { log.Fatal(err)	}
	tpl1 := template.Must(template.New("").Funcs(fm).ParseFiles(wd + "/public/post/product/index.html"))
	//tpl1, err := template.Must(template.New("").Funcs(fm).ParseFiles(wd + "/public/post/product/index.html"))
	//if err != nil {
	//	fmt.Printf("error: %s", err)
	//	return
	//}
	if err :=	tpl1.ExecuteTemplate(w, "index.html", productp); err != nil {	fmt.Printf("error: %s", err) }
	}
}

// /* function called from template for above endpoint */ //
func findProduct1(part string) []Product {	//, product string
	var ppartno []Product
	for i := range products {
		if products[i].PartNo == part {
			ppartno = append(ppartno, products[i])
			break				// Found!
		}
}
	return ppartno
}
// /* individual equipment page ENDPOINT: /equipment/post/{slug} */ //
func findEquipment(w http.ResponseWriter, r *http.Request) {	//, product string
	slug := mux.Vars(r)["slug"]
	a := false
	for i := range equipments {
		if equipments[i].PartNo == slug {
			a = true
			break				// Found!
		}
}
if !a {
	fmt.Fprint(w, "No equipment found for part number:\n", slug)
} else {
	equipmentp := SubPage{"equipments", slug}
	wd, err := os.Getwd()
	if err != nil { log.Fatal(err)	}
	tpl1 := template.Must(template.New("").Funcs(fm).ParseFiles(wd + "/public/post/product/index.html"))
	if err :=	tpl1.ExecuteTemplate(w, "index.html", equipmentp); err != nil {	fmt.Printf("error: %s", err) }
}
}
// /* function called from template for above endpoint */ //
func findEquipment1(part string) []Equipment {	//, product string
	var epartno []Equipment
	for i := range equipments {
		if equipments[i].PartNo == part {
			epartno = append(epartno, equipments[i])
			break				// Found!
		}
	}
	return epartno
}
// /* individual category page ENDPOINT: /cat/{slug}/{id:[0-9]+} */ //
func categoryPage(w http.ResponseWriter, r *http.Request) {
	slug := mux.Vars(r)["slug"]
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	a := false
	for k := range pcats {
		if pcats[k].Name == slug {
			a = true
		}
	}
if !a {
	fmt.Fprint(w, "No product category matching\n", slug)
} else {
	categoryp := Page{"products", slug, id}
	wd, err := os.Getwd()
	if err != nil {	log.Fatal(err)	}
	tpl1 := template.Must(template.New("").Funcs(fm).ParseFiles(wd + "/public/index.html"))
	if err :=	tpl1.ExecuteTemplate(w, "index.html", categoryp); err != nil {	fmt.Printf("error: %s", err) }
}
}
// /* individual equipment category page ENDPOINT: /equipment/cat/{slug}/{id:[0-9]+} */ //
func categoryEquipmentPage(w http.ResponseWriter, r *http.Request) {
	slug := mux.Vars(r)["slug"]
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	a := false
	for k := range ecats {
		if ecats[k].Name == slug {
			a = true
		}
	}
if !a {
	fmt.Fprint(w, "No equipment category matching\n", slug)
} else {
		categoryp := Page{"equipments", slug, id}
		wd, err := os.Getwd()
		if err != nil {	log.Fatal(err)	}
		tpl1 := template.Must(template.New("").Funcs(fm).ParseFiles(wd + "/public/index.html"))
		if err :=	tpl1.ExecuteTemplate(w, "index.html", categoryp); err != nil {	fmt.Printf("error: %s", err) }
	}
}
//}
// /* Front Page - main page ENDPOINT: magnetosphere.net/
func frontPage(w http.ResponseWriter, r *http.Request) {
  slug := mux.Vars(r)["slug"]
	wd, err := os.Getwd()
	if err != nil {	log.Fatal(err) }
	if slug == "" {
	tp1 := template.Must(template.New("").Funcs(fm).ParseFiles(wd + "/public/index.html"))
	if err = tp1.ExecuteTemplate(w, "index.html", Page{"products", "", 0}); err != nil {	fmt.Printf("error: %s", err) }
	} else {
	tp1 := template.Must(template.New("").Funcs(fm).ParseFiles(wd + "/public/404.html"))
	if err = tp1.ExecuteTemplate(w, "404.html", nil); err != nil {	fmt.Printf("error: %s", err) }
}
}
// /* page Page - /p/ page ENDPOINT: magnetosphere.net/p/{id:[0-9]+} */ //
func pagePage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
  wd, err := os.Getwd()
  if err != nil { log.Fatal(err)	}
  fp := Page{"products", "", id} //no category specified here
	tp1 := template.Must(template.New("").Funcs(fm).ParseFiles(wd + "/public/index.html"))
	if err := tp1.ExecuteTemplate(w, "index.html", fp); err != nil {	fmt.Printf("error: %s", err) }
}
// /* List Page - text-only listings main page ENDPOINT: magnetosphere.net/list OR: magnetosphere.net/list/{id:[0-9]+} */ //
func listPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
wd, err := os.Getwd()
if err != nil { log.Fatal(err)	}
fp := Page{"products", "", id} //no category specified here
	tp1 := template.Must(template.New("").Funcs(fm).ParseFiles(wd + "/public/list/index.html"))
	if err := tp1.ExecuteTemplate(w, "index.html", fp); err != nil {	fmt.Printf("error: %s", err) }
}
// /* Equipment Page - main equipment page ENDPOINT magnetosphere.net/equipment OR magnetosphere.net/equipment/p/{id:[0-9]+} */ //
func equipmentPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
wd, err := os.Getwd()
if err != nil { log.Fatal(err)	}
fp := Page{"equipments", "", id} //no category specified here
	tp1 := template.Must(template.New("").Funcs(fm).ParseFiles(wd + "/public/index.html"))
	if err := tp1.ExecuteTemplate(w, "index.html", fp); err != nil {	fmt.Printf("error: %s", err) }
}
// /* function called from the template to get products by page */ //
func productListPage(pagenumber int) []Product {
	var products1 []Product
	fmt.Println(pagenumber)
	lenprod := len(products)
	if pagenumber == 0 {
		for i := 0; i < lenprod; i++ {
			if products[i].Enable == true {
				products1 = append(products1, products[i])
			}
			//pull products from the whole database for the first page!
			randnu := rand.New(rand.NewSource(time.Now().Unix()))
			products2 := make([]Product, len(products1))
			perm := randnu.Perm(len(products1))
			for i, randIndex := range perm {
				products2[i] = products1[randIndex]
			}
		}
	} else {
		//subsequent pages
		for i := ((pagenumber - 1) * 10); len(products1) < 10; i++ {
			if products[i].Enable == false {break}
			products1 = append(products1, products[i])
		}
	}
	//}	//randomize the order in which products appear on the page
	randnu := rand.New(rand.NewSource(time.Now().Unix()))
	products2 := make([]Product, len(products1))
	perm := randnu.Perm(len(products1))
	for i, randIndex := range perm {
		products2[i] = products1[randIndex]
	}
  //	//randomize the categories products are selected from
	randnu1 := rand.New(rand.NewSource(time.Now().Unix()))
	cats2 := make([]Category, len(pcats))
	perm1 := randnu1.Perm(len(pcats))
	for i, randIndex := range perm1 {
		cats2[i] = pcats[randIndex]
	}
cats2 = cats2[:10]
	//limit to 10 results!
  //k := len(cats2)
  //l := len(products2)
	products3 := make([]Product, 10)
  //m := len(products3)
  for i:= range products2 {
      if products3 == nil || len(products3)==0{ products3 = append(products3, products2[i]) } else {
          founded:=false
          for j:= range products3{
            if products3[j].Category == products2[i].Category {
              founded=true
              }
            }
          if !founded{
            products3 = append(products3, products2[i])
            }
      }
  }
  //sort.Sort(alphab(pcats))
//    for i := 0; i < 10; i++ {
//          products3[i] = products2[i]
//    }
	return products3
}
// /* function called from the template to get products by page */ //
func productIndexPage(pagenumber int) int {
	var products1 []Product
		for i := (pagenumber * 10); len(products1) < 10; i++ {
				if products[i].Enable == false {break}
				products1 = append(products1, products[i])
			}
			fmt.Println(len(products1))
return len(products1)
}
// /* function called from the template to get products by page */ //
func lenprods() int {
return lenproducts
}
// /* function called from the template to get equipments by page */ //
func equipmentListPage(pagenumber int) []Equipment {
	//upper := (pagenumber + 1) * 10
	//lower := pagenumber * 10
	var equipments1 []Equipment
	fmt.Println(pagenumber)
	//var a int
	//	a = len(equipments)
	//b := 0
	//if a < 10 {	b = a	} else {	b = 10	}
	//if a < upper {	upper = a	}
	//if pagenumber == 0 {
	for i := 0; len(equipments1) < 10; i++ {
			if equipments[i].Enable == false {break}
			equipments1 = append(equipments1, equipments[i])
		}
		//for i := 0; i < b; i++ {
		//		if equipments[i].Enable == false {break}
		//		equipments1 = append(equipments1, equipments[i])
		//	}
		//} else {
		//	for i := lower; i < upper; i++  {
		//			if equipments[i].Enable == false {break}
		//			equipments1 = append(equipments1, equipments[i])
		//		}
	//	}
	//}	//randomize the order in which products appear on the page
	//var products2
		randnu := rand.New(rand.NewSource(time.Now().Unix()))
		equipments2 := make([]Equipment, len(equipments1))
		perm := randnu.Perm(len(equipments1))
		for i, randIndex := range perm {
			equipments2[i] = equipments1[randIndex]
		}
return equipments2
}
// /* function called from the template to get products by category & page */ //
func productsCategoryListPage(cat string, pagenumber int) []Product {
	var categoryProducts []Product
	var categoryProducts1 []Product
	for i := range products { if products[i].Category == cat {	categoryProducts = append(categoryProducts, products[i]) } }
	upper := (pagenumber + 1) * 10
	lower := pagenumber * 10
	fmt.Println(pagenumber)
	a := len(categoryProducts)
	b := 0
	if a < 10 {	b = a	} else {	b = 10	}
	if a < upper {	upper = a	}
	if pagenumber == 0 {
		for i := 0; i < b; i++ {
			if categoryProducts[i].Enable == false {break}
			categoryProducts1 = append(categoryProducts1, categoryProducts[i])
		}
		} else {
			for i := lower; i < upper; i++  {
					if categoryProducts[i].Enable == false {break}
					categoryProducts1 = append(categoryProducts1, categoryProducts[i])
		}
	}	//randomize the order in which products appear on the page
  randnu := rand.New(rand.NewSource(time.Now().Unix()))
		categoryProducts2 := make([]Product, len(categoryProducts1))
		perm := randnu.Perm(len(categoryProducts1))
		for i, randIndex := range perm { categoryProducts2[i] = categoryProducts1[randIndex] }
		return categoryProducts2
}
// /* function called from the template to get products by category & page */ //
func equipmentsCategoryListPage(cat string, pagenumber int) []Equipment {
	var categoryEquipments []Equipment
	var categoryEquipments1 []Equipment
	for i := range equipments {	if equipments[i].Category == cat { categoryEquipments = append(categoryEquipments, equipments[i]) }	}
	upper := (pagenumber + 1) * 10
	lower := pagenumber * 10
	fmt.Println(pagenumber)
	a := len(categoryEquipments)
	b := 0
	if a < 10 {	b = a	} else {	b = 10	}
	if a < upper {	upper = a	}
	if pagenumber == 0 {
		for i := 0; i < b; i++ {
			if categoryEquipments[i].Enable == false {break}
			categoryEquipments1 = append(categoryEquipments1, categoryEquipments[i])
		}
		} else {
			for i := lower; i < upper; i++  {
					if categoryEquipments[i].Enable == false {break}
					categoryEquipments1 = append(categoryEquipments1, categoryEquipments[i])
					}	//randomize the order in which products appear on the page
				}
  randnu := rand.New(rand.NewSource(time.Now().Unix()))
		categoryEquipments2 := make([]Equipment, len(categoryEquipments1))
		perm := randnu.Perm(len(categoryEquipments1))
		for i, randIndex := range perm { categoryEquipments2[i] = categoryEquipments1[randIndex] }
		return categoryEquipments2
	}
	// /* function called from the template to get products  for the text-only listings page */ //
	func listPage1(pagenumber int) []Product {

		//if pagenumber == 0 {
			return products
		//		}
	}
	// /* function called from the template to get index number for products by page */ //
	func indexPage1(pagenumber int) int {
		var products1 []Product
			for i := 0; i < len(products1); i++ {
					if products[i].Enable == false {break}
					products1 = append(products1, products[i])
				}
//				fmt.Println(len(products1))
	return len(products1)
	}

// /* Blog Pages - main page ENDPOINT: magnetosphere.net/ OR: magnetosphere.net/p/{id:[0-9]+} */ //
func blogPage(w http.ResponseWriter, r *http.Request) {
	slug := mux.Vars(r)["slug"]
	wd, err := os.Getwd()
	if err != nil {	log.Fatal(err) }
	if slug == "" {
	tp1 := template.Must(template.New("").Funcs(fm).ParseFiles(wd + "/public/blog/index.html"))
	if err = tp1.ExecuteTemplate(w, "index.html", nil); err != nil {	fmt.Printf("error: %s", err) }
	} else {
	tp1 := template.Must(template.New("").Funcs(fm).ParseFiles(wd + "/public/blog/" + slug + "/index.html"))
	if err = tp1.ExecuteTemplate(w, "index.html", nil); err != nil {	fmt.Printf("error: %s", err) }
}
}
// /* Single Pages  */ //
// /* About Page  */ //
func aboutPage(w http.ResponseWriter, r *http.Request) {

	wd, err := os.Getwd()
	if err != nil {	log.Fatal(err) }
	tp1 := template.Must(template.New("").Funcs(fm).ParseFiles(wd + "/public/about/index.html"))
	if err := tp1.ExecuteTemplate(w, "index.html", nil); err != nil {	fmt.Printf("error: %s", err) }
}
// /* friends Page  */ //
func friendPage(w http.ResponseWriter, r *http.Request) {
	wd, err := os.Getwd()
	if err != nil {	log.Fatal(err)	}
	tp1 := template.Must(template.New("").Funcs(fm).ParseFiles(wd + "/public/friend/index.html"))
	if err := tp1.ExecuteTemplate(w, "index.html", nil); err != nil {	fmt.Printf("error: %s", err) }
}
// /* Shipping / orders policy Page  */ //
func policyPage(w http.ResponseWriter, r *http.Request) {
	wd, err := os.Getwd()
	if err != nil {	log.Fatal(err)	}
	tp1 := template.Must(template.New("").Funcs(fm).ParseFiles(wd + "/public/policy/index.html"))
	if err := tp1.ExecuteTemplate(w, "index.html", nil); err != nil {	fmt.Printf("error: %s", err) }
}
// /*  */ //
func NewRouter() *mux.Router {
router := mux.NewRouter().StrictSlash(true)
staticDir := "/static/"	// Choose the folder to serve
router.	// Create the route
	PathPrefix(staticDir).
	Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("."+staticDir))))
return router
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
// /* from upper/db tutorial -OR- cockroachdb documentation for upper/db */ //
func Products(sess db.Session) db.Store {	// Products is a handy way to represent a collection.
return sess.Collection("products")
}
func Equipments(sess db.Session) db.Store {	// Equipments is a handy way to represent a collection.
return sess.Collection("equipments")
}

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

func (a *Product) Store(sess db.Session) db.Store {// Collection is required in order to create a relation between the Product struct and the "products" table.
return Products(sess)
}
func (a *Equipment) Store(sess db.Session) db.Store {// Collection is required in order to create a relation between the Equipment struct and the "equipments" table.
return Equipments(sess)
}

func dropProductsTable(sess db.Session) error {
	fmt.Printf("Dropping 'products' table\n")
	_, err := sess.SQL().Exec(`
		DROP TABLE product.products
		`)
	if err != nil {
		return err
	}
	return nil
}
func dropEquipmentsTable(sess db.Session) error {
	fmt.Printf("Dropping 'equipments' table\n")
	_, err := sess.SQL().Exec(`
		DROP TABLE product.equipments
		`)
	if err != nil {
		return err
	}
	return nil
}
//todo: improve importing
func importProductsCSV(sess db.Session) error {
	fmt.Printf("Importing CSV from http://127.0.0.1:8079/export01.csv\n")
	_, err := sess.SQL().Exec(`IMPORT INTO product.products CSV DATA ('http://127.0.0.1:8079/productsexport01.csv')	WITH skip = '1';`)
	if err != nil {
		return err
	}
	return nil
}
func importEquipmentsCSV(sess db.Session) error {
	fmt.Printf("Importing CSV from http://127.0.0.1:8079/equipmentsexport01.csv\n")
	_, err := sess.SQL().Exec(`IMPORT INTO product.equipments CSV DATA ('http://127.0.0.1:8079/equipmentsexport01.csv')	WITH skip = '1';`)
	if err != nil {
		return err
	}
	return nil
}
//correct way from bash shell:
//cockroach sql --certs-dir=certs -e "SELECT * from product.products;" --format=csv > export01.csv
func exportProductsCSV() { //the extremely lazy way
	fmt.Printf("Exporting products table to csv\n")
output, err := exec.Command("make", "export-products").CombinedOutput()
if err != nil {
  os.Stderr.WriteString(err.Error())
}
fmt.Println(string(output))
}
func exportEquipmentsCSV() { //the extremely lazy way
	fmt.Printf("Exporting equipments table to csv\n")
output, err := exec.Command("make", "export-equipments").CombinedOutput()
if err != nil {
  os.Stderr.WriteString(err.Error())
}
fmt.Println(string(output))
}
//id SERIAL PRIMARY KEY,
//id INT8 PRIMARY KEY DEFAULT unique_rowid(),
func createProductsTableIfNotExists(sess db.Session) error {	// createTables creates all the tables that are neccessary to run this example.
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
	func createEquipmentsTableIfNotExists(sess db.Session) error {	// createTables creates all the tables that are neccessary to run this example.
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

///*	database test stuff	*///
func deleteAllProducts(sess db.Session) {
fmt.Printf("Clearing products table\n")
//clear tables ; testing
err := Products(sess).Truncate()
if err != nil {
  log.Fatal("Truncate: ", err)
}
}
func deleteAllEquipments(sess db.Session) {
fmt.Printf("Clearing equipments table\n")
//clear tables ; testing
err := Equipments(sess).Truncate()
if err != nil {
  log.Fatal("Truncate: ", err)
}
}
//creates a single test product
func createTestProd(sess db.Session) {
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
func createTestEquip(sess db.Session) {
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

func createProduct(sess db.Session, partno string) {
fmt.Printf("Creating product with part number: %s'\n", createpartno)
product1 := Product{PartNo: createpartno, Enable: false}
err := Products(sess).InsertReturning(&product1)
if err != nil {
    log.Fatal("sess.Save: ", err)
}
}
//creates 5000 products
func create5kP(sess db.Session) {
for i := 0; i < 5000; i++ {
	    p := strconv.Itoa(i)
			fmt.Printf("Creating product part number: %s\n", p)
			p1 := Product{PartNo: p, Enable: false}
			err := Products(sess).InsertReturning(&p1)
			if err != nil {
			    log.Fatal("sess.Save: ", err)
			}
}
}//creates 5000 equipments
func create5kE(sess db.Session) {
for i := 0; i < 5000; i++ {
	    p := strconv.Itoa(i)
			fmt.Printf("Creating equipment part number: %s\n", p)
			p1 := Equipment{PartNo: p, Enable: false}
			err := Equipments(sess).InsertReturning(&p1)
			if err != nil {
			    log.Fatal("sess.Save: ", err)
			}
}
}//create 100 products
func createHundredP(sess db.Session) {
for i := 0; i < 100; i++ {
	    p := strconv.Itoa(i)
			fmt.Printf("Creating product part number: %s\n", p)
			p1 := Product{PartNo: p, Enable: false}
			err := Products(sess).InsertReturning(&p1)
			if err != nil {
			    log.Fatal("sess.Save: ", err)
			}
}
}//create 100 equipments
func createHundredE(sess db.Session) {
for i := 0; i < 100; i++ {
	    p := strconv.Itoa(i)
			fmt.Printf("Creating equipment part number: %s\n", p)
			p1 := Equipment{PartNo: p, Enable: false}
			err := Equipments(sess).InsertReturning(&p1)
			if err != nil {
			    log.Fatal("sess.Save: ", err)
			}
}
}
