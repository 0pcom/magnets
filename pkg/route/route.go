/*route.go*/
package route
import (
"log"
"fmt"
"github.com/gorilla/mux"
"github.com/gorilla/handlers"
"github.com/0pcom/magnets/pkg/handle"
"net/http"
"os"
)

func Server(webPort int){
  finish := make(chan bool)
  r1 := newRouter()
  Serve := r1
  fmt.Printf("listening on http://127.0.0.1:%d using gorilla router\n", webPort)
go func() {   log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", webPort), Serve))  }()
        //select{} // block forever to prevent exiting
<-finish
//
}

func newRouter() *mux.Router {
  r := mux.NewRouter()
  r.PathPrefix("/img/").Handler(http.StripPrefix("/img/", http.FileServer(http.Dir("./img")))) //images
  r.Handle("/", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(handle.Haltingstate))).Methods("GET") //site root
  return r
}
