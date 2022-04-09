package commands

import (
	"log"
	"github.com/spf13/cobra"
	"github.com/skycoin/skycoin/src/util/logging"
	"github.com/sirupsen/logrus"
	inv "github.com/0pcom/magnets/cmd/magnets/commands/inv"
	run "github.com/0pcom/magnets/cmd/magnets/commands/run"
	//sitemap "github.com/0pcom/magnets/cmd/magnets/commands/sitemap"
)

func init() {
	RootCmd.CompletionOptions.DisableDefaultCmd = true
	RootCmd.AddCommand(
		inv.RootCmd,
		run.RootCmd,
//		sitemap.RootCmd,
	)
}

var RootCmd = &cobra.Command{
	Use:   "magnets",
	Short: "magnetosphere.net website implementation",
	Run: func(_ *cobra.Command, _ []string) {
		mLog := logging.NewMasterLogger()
		mLog.SetLevel(logrus.InfoLevel)
	},
}

// Execute executes root CLI command.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Fatal("Failed to execute command: ", err)
	}
}
