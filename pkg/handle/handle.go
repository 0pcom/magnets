// handle.go //

package handle

import (
"fmt"
"net/http"
"html/template"
//"github.com/gorilla/handlers"
"strconv"
funcmap "github.com/0pcom/magnets/pkg/funcmap"
inv "github.com/0pcom/magnets/pkg/inv"
    "log"
  	"github.com/gorilla/mux"
  	"os"
)
var(
title = "we have the technology  "
)

//type Page struct {
//	Title string
//	Partno string
//	Table string //specifies the table; products or equipments
//  Category string //category
//	View string
//  PageNumber int	//pagination
//}

// custom 404 not found page
func Page404() http.Handler {
return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
  wd, err := os.Getwd()
	if err != nil { log.Fatal(err)	}
	tpl1 := template.Must(template.New("").Funcs(funcmap.FM).ParseFiles(wd + "/public/404.html"))
	if err :=	tpl1.ExecuteTemplate(w, "404.html", nil); err != nil {	fmt.Printf("error: %s", err) }
})
}
// custom 404 not found page
func Page404Page (w http.ResponseWriter, r *http.Request) {
  wd, err := os.Getwd()
	if err != nil { log.Fatal(err)	}
	tpl1 := template.Must(template.ParseFiles(wd + "/public/404.html"))
	if err :=	tpl1.ExecuteTemplate(w, "404.html", nil); err != nil {	fmt.Printf("error: %s", err) }
}

//func helpmenu1() {
//	fmt.Printf("Usage: magnets -dDctCyepirh\n")
//	fmt.Printf("\tSuggested Demo: magnets -ctpr\n")
//	flags.PrintDefaults()
//}




// // individual product page ENDPOINT: /post/{slug} // //
func FindProduct(w http.ResponseWriter, r *http.Request) {	//, product string
  var title01 string
	slug := mux.Vars(r)["slug"]
	a := false
	for i := range inv.Mproducts {
		if inv.Mproducts[i].PartNo == slug {
      title01 = inv.Mproducts[i].Name
      if inv.Mproducts[i].Category == "resistor" {
        title01 = title01 + " resistor"
    }
			a = true
			break				// Found!
		}
}
if !a {
	fmt.Fprint(w, "No product found for part number:\n", slug)
} else {
	productp := funcmap.Page{title01, slug, "products", "", "0", 0}
	wd, err := os.Getwd()
	if err != nil { log.Fatal(err)	}
	tpl1 := template.Must(template.New("").Funcs(funcmap.FM).ParseFiles(wd + "/public/post/product/index.html"))
	//tpl1, err := template.Must(template.New("").Funcs(funcmap.FM).ParseFiles(wd + "/public/post/product/index.html"))
	//if err != nil {
	//	fmt.Printf("error: %s", err)
	//	return
	//}
	if err :=	tpl1.ExecuteTemplate(w, "index.html", productp); err != nil {	fmt.Printf("error: %s", err) }
	}
}


// // individual equipment page ENDPOINT: /equipment/post/{slug} // //
func FindEquipment(w http.ResponseWriter, r *http.Request) {	//, product string
	slug := mux.Vars(r)["slug"]
	a := false
  var title01 string
	for i := range inv.Mequipments {
		if inv.Mequipments[i].PartNo == slug {
			a = true
			break				// Found!
		}
}
if !a {
	fmt.Fprint(w, "No equipment found for part number:\n", slug)
} else {
  title01 = slug
	equipmentp := funcmap.Page{title01, slug, "equipments", "", "0", 0}
	wd, err := os.Getwd()
	if err != nil { log.Fatal(err)	}
	tpl1 := template.Must(template.New("").Funcs(funcmap.FM).ParseFiles(wd + "/public/post/product/index.html"))
	if err :=	tpl1.ExecuteTemplate(w, "index.html", equipmentp); err != nil {	fmt.Printf("error: %s", err) }
}
}
// // individual category page ENDPOINT: /cat/{slug}/{id:[0-9]+} // //
func CategoryPage(w http.ResponseWriter, r *http.Request) {
	slug := mux.Vars(r)["slug"]
	vars := mux.Vars(r)
  var title01 string
	id, _ := strconv.Atoi(vars["id"])
	a := false
	for k := range inv.Pcats {
		if inv.Pcats[k].Name == slug {
			a = true
		}
	}
if !a {
	fmt.Fprint(w, "No product category matching\n", slug)
} else {
  title01 = "category: " + slug
	categoryp := funcmap.Page{title01, "", "products", slug, "1", id}
	wd, err := os.Getwd()
	if err != nil {	log.Fatal(err)	}
	tpl1 := template.Must(template.New("").Funcs(funcmap.FM).ParseFiles(wd + "/public/index.html"))
	if err :=	tpl1.ExecuteTemplate(w, "index.html", categoryp); err != nil {	fmt.Printf("error: %s", err) }
}
}
// // individual equipment category page ENDPOINT: /equipment/cat/{slug}/{id:[0-9]+} // //
func CategoryEquipmentPage(w http.ResponseWriter, r *http.Request) {
	slug := mux.Vars(r)["slug"]
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	a := false
	for k := range inv.Ecats {
		if inv.Ecats[k].Name == slug {
			a = true
		}
	}
if !a {
	fmt.Fprint(w, "No equipment category matching\n", slug)
} else {
		categoryp := funcmap.Page{title, "", "equipments", slug, "1", id}
		wd, err := os.Getwd()
		if err != nil {	log.Fatal(err)	}
		tpl1 := template.Must(template.New("").Funcs(funcmap.FM).ParseFiles(wd + "/public/index.html"))
		if err :=	tpl1.ExecuteTemplate(w, "index.html", categoryp); err != nil {	fmt.Printf("error: %s", err) }
	}
}
//}
// // Front Page - main page ENDPOINT: magnetosphere.net/
func FrontPage(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
  slug := mux.Vars(r)["slug"]
	wd, err := os.Getwd()
	if err != nil {	log.Fatal(err) }
  if slug == "p" {
    //title := "<title>we have the technology | magnetosphere.net</title>"
    fp := funcmap.Page{title, "", "products", "", "1", id} //no category specified here
    tp1 := template.Must(template.New("").Funcs(funcmap.FM).ParseFiles(wd + "/public/index.html"))
    if err := tp1.ExecuteTemplate(w, "index.html", fp); err != nil {	fmt.Printf("error: %s", err) }
  }
	if slug != "" && slug != "p" {
    tp1 := template.Must(template.New("").Funcs(funcmap.FM).ParseFiles(wd + "/public/404.html"))
    if err = tp1.ExecuteTemplate(w, "404.html", nil); err != nil {	fmt.Printf("error: %s", err) }
	}
  if slug == "" {
    tp1 := template.Must(template.New("").Funcs(funcmap.FM).ParseFiles(wd + "/public/index.html"))
    if err = tp1.ExecuteTemplate(w, "index.html", funcmap.Page{title, "", "products", "", "1", 0}); err != nil {	fmt.Printf("error: %s", err) }
}
}

// // List Page - text-only listings main page ENDPOINT: magnetosphere.net/list OR: magnetosphere.net/list/{id:[0-9]+} // //
func ListPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
wd, err := os.Getwd()
if err != nil { log.Fatal(err)	}
fp := funcmap.Page{title, "", "products", "", "2", id} //no category specified here
	tp1 := template.Must(template.New("").Funcs(funcmap.FM).ParseFiles(wd + "/public/index.html"))
	if err := tp1.ExecuteTemplate(w, "index.html", fp); err != nil {	fmt.Printf("error: %s", err) }
}
// // Blog Pages - main page ENDPOINT: magnetosphere.net/ OR: magnetosphere.net/p/{id:[0-9]+} // //
func BlogPage(w http.ResponseWriter, r *http.Request) {
  var title01 string
  title01 = "blog"
  fp := funcmap.Page{title01, "blog", "", "", "0", 0}
	slug := mux.Vars(r)["slug"]
	wd, err := os.Getwd()
	if err != nil {	log.Fatal(err) }
	if slug == "" {
	tp1 := template.Must(template.New("").Funcs(funcmap.FM).ParseFiles(wd + "/public/blog/index.html"))
	if err = tp1.ExecuteTemplate(w, "index.html", fp); err != nil {	fmt.Printf("error: %s", err) }
	} else {
	tp1 := template.Must(template.New("").Funcs(funcmap.FM).ParseFiles(wd + "/public/blog/" + slug + "/index.html"))
	if err = tp1.ExecuteTemplate(w, "index.html", nil); err != nil {	fmt.Printf("error: %s", err) }
}
}
// // Single Pages  // //
// // About Page  // //
func AboutPage(w http.ResponseWriter, r *http.Request) {
  var title01 string
  title01 = "about"
  fp := funcmap.Page{title01, "about", "", "", "0", 0}
	wd, err := os.Getwd()
	if err != nil {	log.Fatal(err) }
	tp1 := template.Must(template.New("").Funcs(funcmap.FM).ParseFiles(wd + "/public/about/index.html"))
	if err := tp1.ExecuteTemplate(w, "index.html", fp); err != nil {	fmt.Printf("error: %s", err) }
}
// // friends Page  // //
func FriendPage(w http.ResponseWriter, r *http.Request) {
  var title01 string
  title01 = "friend"
  fp := funcmap.Page{title01, "friend", "", "", "0", 0}
	wd, err := os.Getwd()
	if err != nil {	log.Fatal(err)	}
	tp1 := template.Must(template.New("").Funcs(funcmap.FM).ParseFiles(wd + "/public/friend/index.html"))
	if err := tp1.ExecuteTemplate(w, "index.html", fp); err != nil {	fmt.Printf("error: %s", err) }
}
// // Shipping / orders policy Page  // //
func PolicyPage(w http.ResponseWriter, r *http.Request) {
  var title01 string
  title01 = "policy"
  fp := funcmap.Page{title01, "policy", "", "", "0", 0}
	wd, err := os.Getwd()
	if err != nil {	log.Fatal(err)	}
	tp1 := template.Must(template.New("").Funcs(funcmap.FM).ParseFiles(wd + "/public/policy/index.html"))
	if err := tp1.ExecuteTemplate(w, "index.html", fp); err != nil {	fmt.Printf("error: %s", err) }
}
