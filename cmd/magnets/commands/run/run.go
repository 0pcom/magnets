/*run.go*/
package run

import (
	"github.com/sirupsen/logrus"
	"github.com/skycoin/skycoin/src/util/logging"
	"github.com/spf13/cobra"
	inv "github.com/0pcom/magnets/pkg/inv"
	user "github.com/0pcom/magnets/pkg/user"
	route "github.com/0pcom/magnets/pkg/route"
	"fmt"
	"log"
	"github.com/upper/db/v4/adapter/cockroachdb"
)

var (
	webPort	int
	webPort1 int
)

func init() {
	RootCmd.Flags().IntVarP(&webPort, "port", "p", 8040, "port to serve on")
	RootCmd.Flags().IntVarP(&webPort1, "port1", "q", 8041, "auxilliary port to serve on")
}

var RootCmd = &cobra.Command{
	Use:   "run",
	Short: "run the web application",
//	PreRun: func(_ *cobra.Command, _ []string) {
//	},
	Run: func(_ *cobra.Command, _ []string) {
		mLog := logging.NewMasterLogger()
		mLog.SetLevel(logrus.InfoLevel)
		//  database connection  //
		fmt.Printf("Initializing cockroachDB connection\n")
		//establish the session
		sess, err := cockroachdb.Open(user.FetchSettings())
		if err != nil {	log.Fatal("cockroachdb.Open: ", err)	}
		defer sess.Close()
		inv.DefineProducts(sess)
		route.Server(webPort, webPort1)
	},
}
