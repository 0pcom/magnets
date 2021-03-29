// Routing based on the gorilla/mux router

package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/upper/db/v4"
"github.com/upper/db/v4/adapter/cockroachdb"
)

var settings = cockroachdb.ConnectionURL{
	Host:     "localhost",
	Database: "product",
	User:     "madmin",
	Options: map[string]string{
		// Secure node.
		"sslrootcert": "certs/ca.crt",
		"sslkey":      "certs/client.madmin.key",
		"sslcert":     "certs/client.madmin.crt",
	},
}

func main() {
	fmt.Printf("Initializing cockroachDB connection\n")
	sess, err := cockroachdb.Open(settings)		//establish the session
	if err != nil {
		log.Fatal("cockroachdb.Open: ", err)
	}
	defer sess.Close()
	deleteAll(sess)
}

func deleteAll(sess db.Session) {
fmt.Printf("Clearing tables\n")
//clear tables ; testing
err := Products(sess).Truncate()
if err != nil {
  log.Fatal("Truncate: ", err)
}
}
