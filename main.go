// /* */ //

package main

import (
	"fmt"
	"log"
	"net/http"
//	"os"
//	"sort"
//	"strings"
	"github.com/0pcom/magnets1/gorilla"
)

const port = 8040

func main() {

	fmt.Printf("listening on port %d using gorilla router\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), gorilla.Serve))
}
