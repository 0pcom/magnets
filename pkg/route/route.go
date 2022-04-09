/*route.go*/
package route
import (
"log"
"fmt"
"github.com/gorilla/mux"
"github.com/gorilla/handlers"
"github.com/0pcom/magnets/pkg/handle"
_ "github.com/0pcom/magnets/statik"
//"github.com/0pcom/magnets/statik"
"github.com/rakyll/statik/fs"
"net/http"
"os"
)

//statikFS
//hugo
//hugo --baseURL="https://magnetosphereelectronicsurplus.com" -d public1 --config="config1.toml"
//statik -f -src=./public && go generate
//statik -f -src=./public1 && go generate

func Server(webPort int, webPort1 int){
  finish := make(chan bool)
  statikFS1, err := fs.New()
  if err != nil {
    log.Fatal(err)
  }
  handle1 := handle.WrapperStruct{Dir: "public", Site: "magnetosphere.net"}
  statikFS2, err := fs.New()
  if err != nil {
    log.Fatal(err)
  }
  handle2 := handle.WrapperStruct{Dir: "public1", Site: "magnetosphereelectronicsurplus.com"}
  r1 := newRouter(statikFS1, handle1)
  r2 := newRouter(statikFS2, handle2)
  Serve := r1
  Serve1 := r2
  fmt.Printf("listening on http://127.0.0.1:%d and http://127.0.0.1:%d using gorilla router\n", webPort, webPort1)
go func() {   log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", webPort), Serve))  }()
go func() {  log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", webPort1), Serve1)) }()
        //select{} // block forever to prevent exiting
<-finish
//
}

func newRouter(statikFS http.FileSystem, handle handle.WrapperStruct) *mux.Router {
  r := mux.NewRouter()
  r.NotFoundHandler = handle.Page404handlr()
  r.PathPrefix("/img/").Handler(http.StripPrefix("/img/", http.FileServer(http.Dir("./img")))) //images
  r.Handle("/sitemap.xml", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(handle.SiteMap))).Methods("GET") //pagination
  //products table - main site original endpoints
  r.Handle("/list", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(handle.ListPage))).Methods("GET") //pagination
  r.Handle("/list/", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(handle.ListPage))).Methods("GET") //pagination
  r.Handle("/list/{id:[0-9]+}", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(handle.ListPage))).Methods("GET") //pagination
  r.Handle("/post", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(handle.Page404handlfunc))).Methods("GET")	//individual product page
  r.Handle("/post/", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(handle.Page404handlfunc))).Methods("GET")	//individual product page
  r.Handle("/post/{slug}", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(handle.FindProduct))).Methods("GET")	//individual product page
  r.Handle("/cat/{cat}", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(handle.CategoryPage))).Methods("GET")	//category
  r.Handle("/cat/{cat}/{subcat}", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(handle.CategoryPage))).Methods("GET")	//subcategory
  r.Handle("/cat/{cat}/p/{id:[0-9]+}", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(handle.CategoryPage))).Methods("GET")	//category pagination
  r.Handle("/cat/{cat}/{subcat}/p/{id:[0-9]+}", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(handle.CategoryPage))).Methods("GET")	//subcategory pagination
  //single pages
  r.Handle("/about", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(handle.AboutPage))).Methods("GET")	//about page
  r.Handle("/blog", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(handle.BlogPage))).Methods("GET")	//about page
  r.Handle("/blog/", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(handle.BlogPage))).Methods("GET")	//about page
  r.Handle("/blog/{slug}/", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(handle.BlogPage))).Methods("GET")	//about page
  r.Handle("/friend", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(handle.FriendPage))).Methods("GET")	//friends page
  r.Handle("/friend/", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(handle.FriendPage))).Methods("GET")	//friends page
  r.Handle("/policy", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(handle.PolicyPage))).Methods("GET")	//shipping page
  r.Handle("/policy/", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(handle.PolicyPage))).Methods("GET")	//shipping page
  //r.Handle("/time", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(timeFunc))).Methods("GET")	//shipping page
  //
  //site root
  //FrontPage Fails on any slug detected
  r.Handle("/{slug}", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(handle.FrontPage))).Methods("GET")
  r.Handle("/", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(handle.FrontPage))).Methods("GET") //site root
  r.Handle("/{slug}/{id:[0-9]+}", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(handle.FrontPage))).Methods("GET") // /p/ for pagination
  r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(statikFS)))	//statik sources
  r.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    // this handler behaved as I would expect the r.NotFoundHandler to behave..
    w.WriteHeader(501)
    w.Write([]byte(`{"status":501,"message":"501: Not implemented."}`))
  })
  return r
}
