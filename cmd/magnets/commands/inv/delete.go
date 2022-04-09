package inv

import (
	"fmt"
	"github.com/upper/db/v4/adapter/cockroachdb"
	"log"
	"github.com/spf13/cobra"
	"github.com/skycoin/skycoin/src/util/logging"
		"github.com/sirupsen/logrus"

		inv "github.com/0pcom/magnets/pkg/inv"
	user "github.com/0pcom/magnets/pkg/user"
)



var (
	deleteproducts	bool
//	deleteequipments	bool
	droptable	bool
//	droptable1	bool
)


func init() {
	RootCmd.AddCommand(
		deleteCmd,
	)
}

func init() {
deleteCmd.Flags().BoolVarP(&droptable, "dropp", "d", false, "Drop products table")
//deleteCmd.Flags().BoolVarP(&droptable1, "drope", "e", false, "Drop equipments table")
deleteCmd.Flags().BoolVarP(&deleteproducts, "deletep", "D", false, "Delete the products in the products database")
//deleteCmd.Flags().BoolVarP(&deleteequipments, "deletee", "E", false, "Delete the equipment in the equipments database")
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete products and drop tables from the database",
//	PreRun: func(_ *cobra.Command, _ []string) {
//	},
	Run: func(_ *cobra.Command, _ []string) {

		mLog := logging.NewMasterLogger()
		mLog.SetLevel(logrus.InfoLevel)
		// /* database connection */ //
		fmt.Printf("Initializing cockroachDB connection\n")
		sess, err := cockroachdb.Open(user.FetchSettings())		//establish the session
		if err != nil {
			log.Fatal("cockroachdb.Open: ", err)
		}
		defer sess.Close()
			// /* drop tables */ //
			if droptable {
					inv.DropProductsTable(sess)
				}
			// /* drop tables */ //
			/*
			if droptable1 {
					inv.DropEquipmentsTable(sess)
				}
				*/
			// /* delete products */ //
			if deleteproducts {
				inv.DefineProducts(sess)
				inv.DeleteAllProducts(sess)
			}
			// /* delete products */ //
			/*
			if deleteequipments {
				inv.DefineEquipments(sess)
				inv.DeleteAllEquipments(sess)
			}
			*/
	},
}
