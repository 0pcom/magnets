/*user.go*/

package user


import 	"github.com/upper/db/v4/adapter/cockroachdb"


func FetchSettings() (cockroachdb.ConnectionURL){
  return settings
}

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
