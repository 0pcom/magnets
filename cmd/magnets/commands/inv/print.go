package inv

import (

	"github.com/spf13/cobra"
	"github.com/skycoin/skycoin/src/util/logging"
		"github.com/sirupsen/logrus"
		//"github.com/upper/db/v4"
		"github.com/upper/db/v4/adapter/cockroachdb"

	//	mdb "github.com/0pcom/magnets/pkg/db"
		inv "github.com/0pcom/magnets/pkg/inv"
	user "github.com/0pcom/magnets/pkg/user"
	"fmt"
	"log"
//	"strconv"

)

var (
	printinv	bool
//	printinv1	bool
	vprintinv	bool

)


func init() {
	RootCmd.AddCommand(
		printCmd,
	)
}

func init() {
	printCmd.Flags().BoolVarP(&printinv, "printinv", "p", false, "Print the inv.Mproducts table to the terminal")
//	printCmd.Flags().BoolVarP(&printinv1, "printinv1", "q", false, "Print the inv.Mequipments table to the terminal")
	printCmd.Flags().BoolVarP(&vprintinv, "vprintinv", "v", false, "More verbose printinventory")
}

var printCmd = &cobra.Command{
	Use:   "print",
	Short: "\nprint parts from the database in the terminal",
	Args:  cobra.MinimumNArgs(0),
	Run: func(_ *cobra.Command, _ []string) {
		mLog := logging.NewMasterLogger()
		mLog.SetLevel(logrus.InfoLevel)
		//  database connection  //
			fmt.Printf("Initializing cockroachDB connection\n")
			sess, err := cockroachdb.Open(user.FetchSettings())		//establish the session
			if err != nil {
				log.Fatal("cockroachdb.Open: ", err)
			}
		  defer sess.Close()
			// /* print inv.Mproducts table inventory */ //
			if printinv {
				inv.DefineProducts(sess)
				log.Printf("products:")
				for i := range inv.Mproducts {
					//fmt.Printf("product #%d: %#v\n", i, inv.Mproducts[i])
						fmt.Printf("product[%d]:\n", inv.Mproducts[i].Id)
						fmt.Printf("\tpartno:		");	fmt.Printf("%s\n", inv.Mproducts[i].PartNo)
						if inv.Mproducts[i].Image1 != "" {	fmt.Printf("\tImage1:		");	fmt.Printf("%s\n", inv.Mproducts[i].Image1)	}
						if inv.Mproducts[i].Name != "" {	fmt.Printf("\tName:		");	fmt.Printf("%s\n", inv.Mproducts[i].Name)	}
						fmt.Printf("\tQty:		"); fmt.Printf("%d\n", inv.Mproducts[i].Qty)
						fmt.Printf("\tPrice:		"); fmt.Printf("%.2f\n", inv.Mproducts[i].Price)
						fmt.Printf("\tEnable:		"); fmt.Printf("%t\n", inv.Mproducts[i].Enable)
				}
			}
			// /* print inv.Mequipments table inventory*/ //
			/*
			if printinv1 {
				inv.DefineEquipments(sess)
				log.Printf("equipments:")
				for i := range inv.Mequipments {
					//fmt.Printf("product #%d: %#v\n", i, inv.Mproducts[i])
						fmt.Printf("equipment[%d]:\n", inv.Mequipments[i].Id)
						fmt.Printf("\tpartno:		");	fmt.Printf("%s\n", inv.Mequipments[i].PartNo)
						if inv.Mequipments[i].Image1 != "" {	fmt.Printf("\tImage1:		");	fmt.Printf("%s\n", inv.Mequipments[i].Image1)	}
						if inv.Mequipments[i].Name != "" {	fmt.Printf("\tName:		");	fmt.Printf("%s\n", inv.Mequipments[i].Name)	}
						fmt.Printf("\tQty:		"); fmt.Printf("%d\n", inv.Mequipments[i].Qty)
						fmt.Printf("\tPrice:		"); fmt.Printf("%.2f\n", inv.Mequipments[i].Price)
						fmt.Printf("\tEnable:		"); fmt.Printf("%t\n", inv.Mequipments[i].Enable)
				}
			}
			*/
			// /* more verbosely print inventory */ //
			if vprintinv {
				inv.DefineProducts(sess)
				log.Printf("products:")
				for i := range inv.Mproducts {
					fmt.Printf("product #%d: %#v\n", i, inv.Mproducts[i])
						}
					}



	},
}
