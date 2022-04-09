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
    "github.com/ikeikeikeike/go-sitemap-generator/v2/stm"

)
var(
title = "we have the technology  "
)

type WrapperStruct struct {
  Dir string
  Site string
}

// custom 404 not found page
func (ws WrapperStruct) Page404handlr() http.Handler {
return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(404)
  wd, err := os.Getwd()
	if err != nil { log.Fatal(err)	}
	tpl1 := template.Must(template.New("").Funcs(funcmap.FM).ParseFiles(wd + "/" + ws.Dir + "/404.html"))
	if err :=	tpl1.ExecuteTemplate(w, "404.html", nil); err != nil {	fmt.Printf("error: %s", err) }
})
}

// sitemap.xml
func (ws WrapperStruct) SiteMap(w http.ResponseWriter, r *http.Request) {
sm1 := stm.NewSitemap(1)
sm1.SetDefaultHost(ws.Site)
sm1.Create()
sm1.Add(stm.URL{{"loc", "https://" + ws.Site + "/"}, {"changefreq", "daily"}})
sm1.Add(stm.URL{{"loc", "https://" + ws.Site + "/list"}, {"changefreq", "daily"}})
sm1.Add(stm.URL{{"loc", "https://" + ws.Site + "/about"}, {"changefreq", "daily"}})
sm1.Add(stm.URL{{"loc", "https://" + ws.Site + "/blog"}, {"changefreq", "daily"}})
sm1.Add(stm.URL{{"loc", "https://" + ws.Site + "/friend"}, {"changefreq", "daily"}})
sm1.Add(stm.URL{{"loc", "https://" + ws.Site + "/policy"}, {"changefreq", "daily"}})
sm1.Add(stm.URL{{"loc", "https://" + ws.Site + "/post"}, {"changefreq", "daily"}})
sm1.Add(stm.URL{{"loc", "https://" + ws.Site + "/cat"}, {"changefreq", "daily"}})
sm1.Add(stm.URL{{"loc", "https://" + ws.Site + "/p"}, {"changefreq", "daily"}})

  w.Write(sm1.XMLContent())
  return
}


func (ws WrapperStruct) Page404handlfunc(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(404)
  wd, err := os.Getwd()
	if err != nil { log.Fatal(err)	}
	tpl1 := template.Must(template.New("").Funcs(funcmap.FM).ParseFiles(wd + "/" + ws.Dir + "/404.html"))
	if err :=	tpl1.ExecuteTemplate(w, "404.html", nil); err != nil {	fmt.Printf("error: %s", err) }
}

// // individual product page ENDPOINT: /post/{slug} // //
func (ws WrapperStruct) FindProduct(w http.ResponseWriter, r *http.Request) {	//, product string
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
  w.WriteHeader(404)
  wd, err := os.Getwd()
  if err != nil { log.Fatal(err)	}
  tpl1 := template.Must(template.New("").Funcs(funcmap.FM).ParseFiles(wd + "/" + ws.Dir + "/404.html"))
  if err :=	tpl1.ExecuteTemplate(w, "404.html", nil); err != nil {	fmt.Printf("error: %s", err) }
	//fmt.Fprint(w, "No product found for part number:\n", slug)
} else {
	productp := funcmap.Page{title01, slug, "products", "", "", "0", 0, ws.Site}
	wd, err := os.Getwd()
	if err != nil { log.Fatal(err)	}
	tpl1 := template.Must(template.New("").Funcs(funcmap.FM).ParseFiles(wd + "/" + ws.Dir + "/post/product/index.html"))
	//tpl1, err := template.Must(template.New("").Funcs(funcmap.FM).ParseFiles(wd + "/public/post/product/index.html"))
	//if err != nil {
	//	fmt.Printf("error: %s", err)
	//	return
	//}
	if err :=	tpl1.ExecuteTemplate(w, "index.html", productp); err != nil {	fmt.Printf("error: %s", err) }
	}
}

// // individual category page ENDPOINT: /cat/{slug}/{id:[0-9]+} // //
func (ws WrapperStruct) CategoryPage(w http.ResponseWriter, r *http.Request) {
  cat := mux.Vars(r)["cat"]
  subcat := mux.Vars(r)["subcat"]
	vars := mux.Vars(r)
  var title01 string
	id, _ := strconv.Atoi(vars["id"])
	a := false
  b := false
  for k := range inv.Pcats {
		if inv.Pcats[k].Name == cat {
			a = true
      if a {
        for j := range inv.Pcats[k].Subcategories {
          if inv.Pcats[k].Subcategories[j].Name == subcat {
            b = true
          }
        }
      }
    }
  }
  if subcat != ""{
    if !b{
      w.WriteHeader(404)
      wd, err := os.Getwd()
      if err != nil { log.Fatal(err)	}
      tpl1 := template.Must(template.New("").Funcs(funcmap.FM).ParseFiles(wd + "/" + ws.Dir + "/404.html"))
      if err :=	tpl1.ExecuteTemplate(w, "404.html", nil); err != nil {	fmt.Printf("error: %s", err) }
      //fmt.Fprint(w, "No subcategory matching\n", subcat)
    } else {
      title01 = "category: " + subcat
      categoryp := funcmap.Page{title01, "", "products", cat, subcat, "1", id, ws.Site}
      wd, err := os.Getwd()
      if err != nil {	log.Fatal(err)	}
      tpl1 := template.Must(template.New("").Funcs(funcmap.FM).ParseFiles(wd + "/" + ws.Dir + "/index.html"))
      if err :=	tpl1.ExecuteTemplate(w, "index.html", categoryp); err != nil {	fmt.Printf("error: %s", err) }
    }
  } else {
    if !a {
      w.WriteHeader(404)
      wd, err := os.Getwd()
      if err != nil { log.Fatal(err)	}
      tpl1 := template.Must(template.New("").Funcs(funcmap.FM).ParseFiles(wd + "/" + ws.Dir + "/404.html"))
      if err :=	tpl1.ExecuteTemplate(w, "404.html", nil); err != nil {	fmt.Printf("error: %s", err) }
      //fmt.Fprint(w, "No product category matching\n", cat)
    } else {
      title01 = "category: " + cat
      categoryp := funcmap.Page{title01, "", "products", cat, subcat, "1", id, ws.Site}
      wd, err := os.Getwd()
      if err != nil {	log.Fatal(err)	}
      tpl1 := template.Must(template.New("").Funcs(funcmap.FM).ParseFiles(wd + "/" + ws.Dir + "/index.html"))
      if err :=	tpl1.ExecuteTemplate(w, "index.html", categoryp); err != nil {	fmt.Printf("error: %s", err) }
    }
  }
}

// // individual category page ENDPOINT: /cat/{slug}/{id:[0-9]+} // //
func (ws WrapperStruct) CategoryPage1(w http.ResponseWriter, r *http.Request) {
  cat := mux.Vars(r)["cat"]
  subcat := mux.Vars(r)["subcat"]
	vars := mux.Vars(r)
  var title01 string
	id, _ := strconv.Atoi(vars["id"])
	a := false
  b := false
  for k := range inv.Pcats {
		if inv.Pcats[k].Name == cat {
			a = true
      if a {
        for j := range inv.Pcats[k].Subcategories {
          if inv.Pcats[k].Subcategories[j].Name == subcat {
            b = true
          }
        }
      }
    }
  }
  if subcat != ""{
    if !b{
      fmt.Fprint(w, "No subcategory matching\n", subcat)
    } else {
      title01 = "category: " + subcat
      categoryp := funcmap.Page{title01, "", "products", cat, subcat, "1", id, ws.Site}
      wd, err := os.Getwd()
      if err != nil {	log.Fatal(err)	}
      tpl1 := template.Must(template.New("").Funcs(funcmap.FM).ParseFiles(wd + "/" + ws.Dir + "/index.html"))
      if err :=	tpl1.ExecuteTemplate(w, "index.html", categoryp); err != nil {	fmt.Printf("error: %s", err) }
    }
  } else {
    if !a {
      fmt.Fprint(w, "No product category matching\n", cat)
    } else {
      title01 = "category: " + cat
      categoryp := funcmap.Page{title01, "", "products", cat, subcat, "1", id, ws.Site}
      wd, err := os.Getwd()
      if err != nil {	log.Fatal(err)	}
      tpl1 := template.Must(template.New("").Funcs(funcmap.FM).ParseFiles(wd + "/" + ws.Dir + "/index.html"))
      if err :=	tpl1.ExecuteTemplate(w, "index.html", categoryp); err != nil {	fmt.Printf("error: %s", err) }
    }
  }
}


// // Front Page - main page ENDPOINT: magnetosphere.net/
func (ws WrapperStruct) FrontPage(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
  slug := mux.Vars(r)["slug"]
	wd, err := os.Getwd()
	if err != nil {	log.Fatal(err) }
  if slug == "p" {
    //title := "<title>we have the technology | magnetosphere.net</title>"
    fp := funcmap.Page{title, "", "products", "", "", "1", id, ws.Site} //no category specified here
      tp1 := template.Must(template.New("").Funcs(funcmap.FM).ParseFiles(wd + "/" + ws.Dir + "/index.html"))
      if err := tp1.ExecuteTemplate(w, "index.html", fp); err != nil {	fmt.Printf("error: %s", err) }
  }
	if slug != "" && slug != "p" {
    tp1 := template.Must(template.New("").Funcs(funcmap.FM).ParseFiles(wd + "/" + ws.Dir +"/404.html"))
    if err = tp1.ExecuteTemplate(w, "404.html", nil); err != nil {	fmt.Printf("error: %s", err) }
	}
  if slug == "" {
    tp1 := template.Must(template.New("").Funcs(funcmap.FM).ParseFiles(wd + "/" + ws.Dir + "/index.html"))
    if err = tp1.ExecuteTemplate(w, "index.html", funcmap.Page{title, "", "products", "", "", "1", 0, ws.Site}); err != nil {	fmt.Printf("error: %s", err) }
}
}

// // List Page - text-only listings main page ENDPOINT: magnetosphere.net/list OR: magnetosphere.net/list/{id:[0-9]+} // //
func (ws WrapperStruct) ListPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
wd, err := os.Getwd()
if err != nil { log.Fatal(err)	}
fp := funcmap.Page{title, "", "products", "", "", "2", id, ws.Site} //no category specified here
	tp1 := template.Must(template.New("").Funcs(funcmap.FM).ParseFiles(wd + "/" + ws.Dir + "/index.html"))
	if err := tp1.ExecuteTemplate(w, "index.html", fp); err != nil {	fmt.Printf("error: %s", err) }
}
// // Blog Pages - main page ENDPOINT: magnetosphere.net/ OR: magnetosphere.net/p/{id:[0-9]+} // //
func (ws WrapperStruct) BlogPage(w http.ResponseWriter, r *http.Request) {
  var title01 string
  title01 = "blog index"
  fp := funcmap.Page{title01, "blog", "", "", "", "3", 0, ws.Site}
	slug := mux.Vars(r)["slug"]
	wd, err := os.Getwd()
	if err != nil {	log.Fatal(err) }
	if slug == "" {
      tp1 := template.Must(template.New("").Funcs(funcmap.FM).ParseFiles(wd + "/" + ws.Dir + "/blog/index.html"))
      if err = tp1.ExecuteTemplate(w, "index.html", fp); err != nil {	fmt.Printf("error: %s", err) }
	} else {
      tp1 := template.Must(template.New("").Funcs(funcmap.FM).ParseFiles(wd + "/" + ws.Dir + "/blog/" + slug + "/index.html"))
      if err = tp1.ExecuteTemplate(w, "index.html", fp); err != nil {	fmt.Printf("error: %s", err) }
}
}
// // Single Pages  // //
// // About Page  // //
func (ws WrapperStruct) AboutPage(w http.ResponseWriter, r *http.Request) {
  var title01 string
  title01 = "about the site"
  fp := funcmap.Page{title01, "about", "", "", "", "3", 0, ws.Site}
	wd, err := os.Getwd()
	if err != nil {	log.Fatal(err) }
    tp1 := template.Must(template.New("").Funcs(funcmap.FM).ParseFiles(wd + "/" + ws.Dir + "/about/index.html"))
    if err := tp1.ExecuteTemplate(w, "index.html", fp); err != nil {	fmt.Printf("error: %s", err) }
}
// // friends Page  // //
func (ws WrapperStruct) FriendPage(w http.ResponseWriter, r *http.Request) {
  var title01 string
  title01 = "our friends"
  fp := funcmap.Page{title01, "friend", "", "", "", "1", 0, ws.Site}
	wd, err := os.Getwd()
	if err != nil {	log.Fatal(err)	}
	tp1 := template.Must(template.New("").Funcs(funcmap.FM).ParseFiles(wd + "/" + ws.Dir + "/friend/index.html"))
	if err := tp1.ExecuteTemplate(w, "index.html", fp); err != nil {	fmt.Printf("error: %s", err) }
}
// // Shipping / orders policy Page  // //
func (ws WrapperStruct) PolicyPage(w http.ResponseWriter, r *http.Request) {
  var title01 string
  title01 = "shipping and order policies"
  fp := funcmap.Page{title01, "policy", "", "", "", "3", 0, ws.Site}
	wd, err := os.Getwd()
	if err != nil {	log.Fatal(err)	}
	tp1 := template.Must(template.New("").Funcs(funcmap.FM).ParseFiles(wd + "/"+ ws.Dir + "/policy/index.html"))
	if err := tp1.ExecuteTemplate(w, "index.html", fp); err != nil {	fmt.Printf("error: %s", err) }
}
