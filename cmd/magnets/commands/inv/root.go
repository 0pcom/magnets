package inv

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "inv",
	Short: "sub-commands for inventory management - interact with the local cockroachdb instance",
}
