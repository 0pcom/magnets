/*run.go*/
package run

import (
	"github.com/sirupsen/logrus"
	"github.com/skycoin/skycoin/src/util/logging"
	"github.com/spf13/cobra"
	route "github.com/0pcom/magnets/pkg/route"

)

var (
	webPort	int
)

func init() {
	RootCmd.Flags().IntVarP(&webPort, "port", "p", 8081, "port to serve on")
}

var RootCmd = &cobra.Command{
	Use:   "run",
	Short: "run the web application",
	Run: func(_ *cobra.Command, _ []string) {
		mLog := logging.NewMasterLogger()
		mLog.SetLevel(logrus.InfoLevel)
		route.Server(webPort)
	},
}
