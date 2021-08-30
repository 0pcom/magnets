/*funcmap.go*/

package funcmap

import (
	"log"
		inv "github.com/0pcom/magnets/pkg/inv"
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
var FM = template.FuncMap{
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
//func timeFunc(w http.ResponseWriter, r *http.Request) {	//a page with just the time
//	w.Header().Set("Content-Type", "text/html")
//	tp1 := template.Must(template.New("").Funcs(fm).ParseFiles("time.html"))
//	if err := tp1.ExecuteTemplate(w, "time.html",nil); err != nil {	fmt.Printf("error: %s", err) }
//}


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


// // function called from the template to get products by page // //
func productListPage(pagenumber int) []inv.Product {
	var products1 []inv.Product
	fmt.Println(pagenumber)
	lenprod := len(inv.Mproducts)
	if pagenumber == 0 {
		for i := 0; i < lenprod; i++ {
			if inv.Mproducts[i].Enable == true {
				products1 = append(products1, inv.Mproducts[i])
			}
			//pull products from the whole database for the first page!
			randnu := rand.New(rand.NewSource(time.Now().Unix()))
			products2 := make([]inv.Product, len(products1))
			perm := randnu.Perm(len(products1))
			for i, randIndex := range perm {
				products2[i] = products1[randIndex]
			}
		}
	} else {
		//subsequent pages
		for i := ((pagenumber - 1) * 10); len(products1) < 10; i++ {
			if inv.Mproducts[i].Enable == false {break}
			products1 = append(products1, inv.Mproducts[i])
		}
	}
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
cats2 = cats2[:10]
	//limit to 10 results!
  //k := len(cats2)
  //l := len(products2)
	products3 := make([]inv.Product, 10)
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
  //sort.Sort(alphab(inv.Pcats))
//    for i := 0; i < 10; i++ {
//          products3[i] = products2[i]
//    }
	return products3
}
// // function called from the template to get products by page // //
func productIndexPage(pagenumber int) int {
	var products1 []inv.Product
		for i := (pagenumber * 10); len(products1) < 10; i++ {
				if inv.Mproducts[i].Enable == false {break}
				products1 = append(products1, inv.Mproducts[i])
			}
			fmt.Println(len(products1))
return len(products1)
}
// // function called from the template to get products by page // //
func lenprods() int {
return inv.Lenproducts
}
// // function called from the template to get equipments by page // //
func equipmentListPage(pagenumber int) []inv.Equipment {
	//upper := (pagenumber + 1) * 10
	//lower := pagenumber * 10
	var equipments1 []inv.Equipment
	fmt.Println(pagenumber)
	//var a int
	//	a = len(equipments)
	//b := 0
	//if a < 10 {	b = a	} else {	b = 10	}
	//if a < upper {	upper = a	}
	//if pagenumber == 0 {
	for i := 0; len(equipments1) < 10; i++ {
			if inv.Mequipments[i].Enable == false {break}
			equipments1 = append(equipments1, inv.Mequipments[i])
		}
		//for i := 0; i < b; i++ {
		//		if inv.Mequipments[i].Enable == false {break}
		//		equipments1 = append(equipments1, inv.Mequipments[i])
		//	}
		//} else {
		//	for i := lower; i < upper; i++  {
		//			if inv.Mequipments[i].Enable == false {break}
		//			equipments1 = append(equipments1, inv.Mequipments[i])
		//		}
	//	}
	//}	//randomize the order in which products appear on the page
	//var products2
		randnu := rand.New(rand.NewSource(time.Now().Unix()))
		equipments2 := make([]inv.Equipment, len(equipments1))
		perm := randnu.Perm(len(equipments1))
		for i, randIndex := range perm {
			equipments2[i] = equipments1[randIndex]
		}
return equipments2
}
// // function called from the template to get products by category & page // //
func productsCategoryListPage(cat string, pagenumber int) []inv.Product {
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
					if inv.Mproducts[i].Enable == false {break}
					products1 = append(products1, inv.Mproducts[i])
				}
//				fmt.Println(len(products1))
	return len(products1)
	}



	// // function called from template for above endpoint // //
	func findProduct1(part string) []inv.Product {	//, product string
		var ppartno []inv.Product
		for i := range inv.Mproducts {
			if inv.Mproducts[i].PartNo == part {
				ppartno = append(ppartno, inv.Mproducts[i])
				break				// Found!
			}
	}
		return ppartno
	}



	// // function called from template for above endpoint // //
	func findEquipment1(part string) []inv.Equipment {	//, product string
		var epartno []inv.Equipment
		for i := range inv.Mequipments {
			if inv.Mequipments[i].PartNo == part {
				epartno = append(epartno, inv.Mequipments[i])
				break				// Found!
			}
		}
		return epartno
	}
