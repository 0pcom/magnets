/*funcmap.go*/

package funcmap

import (
	"log"
		inv "github.com/0pcom/magnets/pkg/inv"
		"strconv"
)
import (
 	"fmt"
	"math"
	"math/rand"
	"html/template"
	"time"
	"os"
)


//these are functions called from the template or webpage
var FM = template.FuncMap{              //   inputs:           returns:              use:
	"fdateMDY": monthDayYear,              //   n/a               time                  time
	"snipcartapikey": snipcartApiKey,      //   n/a               key from env          snipcart
	"multiply": multiply,                  //   num,num           result                snipcart
	"correct": correct,                    //   num               formatted             snipcart
	"convertozgrams": convertozgrams,      //   num               grams from oz         snipcart
	"convertincm": convertincm,            //   num               in from cm            snipcart
	"listCategories": listCategories,      //   table             categories            menu
	"lenprods": lenprods,                  //   table,cat         len(products)         displays # of products
	"productList": productList,            //   table,cat,page#   []products{}          browse database
	"productList1": productList1,          //   table,cat				  []products{}          browse database
	"productIndex": productIndex,          //   table,cat,page#   index#                pagination
	"page": page,                          //   table,cat,page#   page #                previous, next
	"findProduct1": findProduct1,          //   PartNo            product{}             individual product page
}//below are the functions
// // timepage  // //
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

// // return list of categories // //
func listCategories(table string) []inv.Category {
	var toreturn []inv.Category
	if table == "products" {
		toreturn = inv.RetPcats()
	}
	if table == "equipments" {
		toreturn = inv.RetEcats()
	}
	return toreturn
}


//the "current context" of the template ({{.}}) is now an index
func page(directive string, table string, cat string, pageno int) string {
var pagestr string
var pre string
if table == "products" && cat == ""{		pre = "/p/"}
if table == "products" && cat != ""{		pre = fmt.Sprintf("/cat/%s", cat)}
if table == "equipments" && cat == ""{	pre = "/equipment/p/"}
if table == "equipments" && cat != ""{	pre = fmt.Sprintf("/equipment/cat/%s", cat)}

if directive == "first"{		pagestr = fmt.Sprintf("%s%s", pre, strconv.Itoa(0))}
if directive == "firstt"{		pagestr = strconv.Itoa(0)}
if directive == "prev"{			pagestr = fmt.Sprintf("%s%s", pre, strconv.Itoa(pageno -1))}
if directive == "prevv"{		pagestr = strconv.Itoa(pageno - 1)}
if directive == "next"{			pagestr = fmt.Sprintf("%s%s", pre, strconv.Itoa(pageno + 1))}
if directive == "nextt"{		pagestr = strconv.Itoa(pageno + 1)}
if directive == "last"{			pagestr = fmt.Sprintf("%s%s", pre, strconv.Itoa(lastPage(table, cat)))}
if directive == "lastt"{		pagestr = strconv.Itoa(lastPage(table, cat))}
if directive == "current"{	pagestr = fmt.Sprintf("%s%s", pre, strconv.Itoa(pageno))}
if directive == "currentt"{	pagestr = strconv.Itoa(pageno)}
return pagestr
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
	Title string
	Partno string
	Table string //specifies the table; products or equipments
  Category string //category
	View string
  PageNumber int	//pagination
}
type SubPage struct {
	Title string
	Table string
  PartNumber string
}


// // function called from the template to get products by page // //
func productList(table string, category string, view string, pagenumber int) []inv.Product {
	products1 := inv.Mproducts
	products3 := make([]inv.Product, 0)
	//lenprod := len(inv.Mproducts)
	if view == "2" {
				products3 = products1
	}
	if view == "1" && pagenumber == 0 && category == "" {
		//}	//randomize the order in which products appear on the page
		randnu := rand.New(rand.NewSource(time.Now().Unix()))
		products2 := make([]inv.Product, len(products1))
		perm := randnu.Perm(len(products1))
		for i, randIndex := range perm {
			products2[i] = products1[randIndex]
		}
		//	//randomize the categories products are selected from
		randnu1 := rand.New(rand.NewSource(time.Now().Unix()))
		cats2 := make([]inv.Category, len(inv.Pcats))
		perm1 := randnu1.Perm(len(inv.Pcats))
		for i, randIndex := range perm1 {
			cats2[i] = inv.Pcats[randIndex]
		}
		for i:= range products2 {
			if products3 == nil || len(products3)==0{
				products3 = append(products3, products2[i])
				} else {
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
	//}
	}
	if view == "1" && pagenumber != 0 {
	//subsequent pages
	v := ((pagenumber - 1) * 10)
	if v < (inv.Lenproducts - 10) {
		for i := v; i < (v + 10) ; i++ {
			products3 = append(products3, inv.Mproducts[i])
		}
	}
}
if view == "1" && category != "" {
		products3 = productCat(table, category, pagenumber)
	}

return products3
}
// // function called from the template to get products by page // //
func productList1(table string, category string) []inv.Product {
	ret := make([]inv.Product, 0)
	if table == "products" {
		ret = inv.Mproducts
	}
	if table == "equipments" {
		ret = inv.Mequipments
	}
return ret
}
// // function called from the template to get products by page // //
func equipmentList(pagenumber int) []inv.Product {
	products1 := inv.Mproducts
	products3 := make([]inv.Product, 0)
	//lenprod := len(inv.Mproducts)
	if pagenumber == 0 {
		//}	//randomize the order in which products appear on the page
		randnu := rand.New(rand.NewSource(time.Now().Unix()))
		products2 := make([]inv.Product, len(products1))
		perm := randnu.Perm(len(products1))
		for i, randIndex := range perm {
			products2[i] = products1[randIndex]
		}
		//	//randomize the categories products are selected from
		randnu1 := rand.New(rand.NewSource(time.Now().Unix()))
		cats2 := make([]inv.Category, len(inv.Pcats))
		perm1 := randnu1.Perm(len(inv.Pcats))
		for i, randIndex := range perm1 {
			cats2[i] = inv.Pcats[randIndex]
		}
		for i:= range products2 {
			if products3 == nil || len(products3)==0{
				products3 = append(products3, products2[i])
				} else {
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
	//}
	} else {
		//subsequent pages
		v := ((pagenumber - 1) * 10)
		if v < (inv.Lenproducts - 10) {
			for i := v; i < (v + 10) ; i++ {
				products3 = append(products3, inv.Mproducts[i])
			}
		}
	}
return products3
}

// // index has a role in the display of pagination options // //
func productIndex(table string, cat string, pageno int) int {
	var toreturn int
	v := (pageno * 10)
	if table == "products" && cat == ""{		if v < (len(inv.Mproducts) - 10) {toreturn = (len(inv.Mproducts) - v)} else { toreturn = 0 }}
	if table == "products" && cat != ""{		toreturn = productCatIndex(table, cat, pageno)}
	if table == "equipments" && cat == ""{	if v < (len(inv.Mproducts) - 10) {toreturn = (len(inv.Mproducts) - v)} else { toreturn = 0 }}
	if table == "equipments" && cat != ""{	toreturn = productCatIndex(table, cat, pageno)}
	return toreturn
}

// // function called from the template to get products by page // //
func productCatIndex(table string, cat string, pageno int) int {
	var catcount int
	var toreturn int
	if table == "products" {
		for i := range inv.Pcats {
			a := false
			if inv.Pcats[i].Name == cat {
				catcount = inv.Pcats[i].Count
				a = true
			}
			if a == true {break}
		}
	}
	if table == "equipments"{
		for i := range inv.Ecats {
			a := false
			if inv.Ecats[i].Name == cat {
				catcount = inv.Ecats[i].Count
				a = true
			}
			if a == true {break}
		}

	}
	v := (pageno * 10)
	if v < (catcount - 10) {toreturn = (catcount - v)} else { toreturn = 0 }
return toreturn
}
// // function called from the template to get products by page // //
func lenprods() int {
return inv.Lenproducts
}
// // function called from the template to get equipments by page // //
//func equipmentList(pagenumber int) []inv.Equipment {
//	var equipments1 []inv.Equipment
//	for i := 0; len(equipments1) < 10; i++ {
//			equipments1 = append(equipments1, inv.Mequipments[i])
//		}
//		randnu := rand.New(rand.NewSource(time.Now().Unix()))
//		equipments2 := make([]inv.Equipment, len(equipments1))
//		perm := randnu.Perm(len(equipments1))
//		for i, randIndex := range perm {
//			equipments2[i] = equipments1[randIndex]
//		}
//return equipments2
//}
// // function called from the template to get products by category & page // //
func productCat(table string, cat string, pagenumber int) []inv.Product {
	var categoryProducts []inv.Product
	var categoryProducts1 []inv.Product
	for i := range inv.Mproducts { if inv.Mproducts[i].Category == cat {	categoryProducts = append(categoryProducts, inv.Mproducts[i]) } }
	upper := (pagenumber + 1) * 10
	lower := pagenumber * 10
	fmt.Println(pagenumber)
	a := len(categoryProducts)
	b := 0
	if a < 10 {	b = a	} else {	b = 10	}
	if a < upper {	upper = a	}
	if pagenumber == 0 {
		for i := 0; i < b; i++ {
			categoryProducts1 = append(categoryProducts1, categoryProducts[i])
		}
		} else {
			for i := lower; i < upper; i++  {
					categoryProducts1 = append(categoryProducts1, categoryProducts[i])
		}
	}	//randomize the order in which products appear on the page
  randnu := rand.New(rand.NewSource(time.Now().Unix()))
		categoryProducts2 := make([]inv.Product, len(categoryProducts1))
		perm := randnu.Perm(len(categoryProducts1))
		for i, randIndex := range perm { categoryProducts2[i] = categoryProducts1[randIndex] }
		return categoryProducts2
}
// // function called from the template to get products by category & page // //
func equipmentsCategoryListPage(cat string, pagenumber int) []inv.Equipment {
	var categoryEquipments []inv.Equipment
	var categoryEquipments1 []inv.Equipment
	for i := range inv.Mequipments {	if inv.Mequipments[i].Category == cat { categoryEquipments = append(categoryEquipments, inv.Mequipments[i]) }	}
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
		categoryEquipments2 := make([]inv.Equipment, len(categoryEquipments1))
		perm := randnu.Perm(len(categoryEquipments1))
		for i, randIndex := range perm { categoryEquipments2[i] = categoryEquipments1[randIndex] }
		return categoryEquipments2
	}
	// // function called from the template to get products  for the text-only listings page // //
	func listPage1(pagenumber int) []inv.Product {

		//if pagenumber == 0 {
			return inv.Mproducts
		//		}
	}
	// // function called from the template to get index number for products by page // //
	func indexPage1(pagenumber int) int {

		var products1 []inv.Product
			for i := 0; i < len(products1); i++ {
					products1 = append(products1, inv.Mproducts[i])
				}
//				fmt.Println(len(products1))

	return len(products1)
	}



	// // returns the database entry given the part number // //
	func findProduct1(table string, partno string) []inv.Product {	//, product string
		var ppartno []inv.Product
		if table == "products" {
		for i := range inv.Mproducts {
			if inv.Mproducts[i].PartNo == partno {
				ppartno = append(ppartno, inv.Mproducts[i])
				break				// Found!
			}
		}
	}
	if table == "equipments" {
	for i := range inv.Mequipments {
		if inv.Mequipments[i].PartNo == partno {
			ppartno = append(ppartno, inv.Mequipments[i])
			break				// Found!
		}
	}
}
		return ppartno
	}



	// // function called from template for above endpoint // //
	func findEquipment1(part string) inv.Equipment {	//, product string
		var epartno inv.Equipment
		for i := range inv.Mequipments {
			if inv.Mequipments[i].PartNo == part {
				epartno = inv.Mequipments[i]
				break				// Found!
			}
		}
		return epartno
	}

	// // returns the last page with products // //
	func lastPage(table string, cat string) int {
		var toreturn int
		if table == "products" {
			if cat == "" {
				toreturn = int(math.Floor(float64(len(inv.Mproducts) / 10)))
				} else {
					a := false
					for i := range inv.Pcats {
						if inv.Pcats[i].Name == cat {
							a = true
							toreturn = int(math.Floor(float64(inv.Pcats[i].Count) / 10))
						}
						if a == true {break}
					}
				}
			}
		if table == "equipments" {
			if cat == "" {
				toreturn = int(math.Floor(float64(len(inv.Mequipments) / 10)))
				} else {
					a := false
					for i := range inv.Ecats {
						if inv.Ecats[i].Name == cat {
							a = true
							toreturn = int(math.Floor(float64(inv.Ecats[i].Count) / 10))
						}
						if a == true {break}
					}
				}
			}
		return toreturn
	}
