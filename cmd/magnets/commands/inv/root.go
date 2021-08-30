package inv

import (
	"github.com/spf13/cobra"
)


//func init() {
//	RootCmd.PersistentFlags().StringVarP(&rpcAddr, "rpc", "", "localhost:3435", "RPC server address")
//}

// RootCmd contains commands that interact with the skywire-visor
var RootCmd = &cobra.Command{
	Use:   "inv",
	Short: "sub-commands for inventory management - interact with the local cockroachdb instance",
}
