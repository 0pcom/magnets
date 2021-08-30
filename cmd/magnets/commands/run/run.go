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
)


//func init() {
//	RootCmd.AddCommand(
//		runCmd,
//	)
//}

func init() {
RootCmd.Flags().IntVarP(&webPort, "port", "p", 8040, "port to serve on")

//RootCmd.Flags().Parse()
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
			sess, err := cockroachdb.Open(user.FetchSettings())		//establish the session
			if err != nil {	log.Fatal("cockroachdb.Open: ", err)	}
		  defer sess.Close()
			inv.DefineProducts(sess)
			inv.DefineEquipments(sess)
			route.Server(webPort)
	},
}
