/*db.go*/

package db

import (
 	"fmt"
	"log"
	"github.com/upper/db/v4/adapter/cockroachdb"
  "github.com/upper/db/v4"
)

func Connect(settings cockroachdb.ConnectionURL) (db.Session) {
// /* database connection */ //
	fmt.Printf("Initializing cockroachDB connection\n")
	sess, err := cockroachdb.Open(settings)		//establish the session
	if err != nil {
		log.Fatal("cockroachdb.Open: ", err)
	}
  defer sess.Close()
  return sess
}
