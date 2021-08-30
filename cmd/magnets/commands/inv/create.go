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
	createtables	bool
	testprod	bool
	testequip	bool
	createpartno	string
	createseriesp	int
	createseriese	int
	create5000p	bool
	create5000e	bool
)


func init() {
	RootCmd.AddCommand(
		createCmd,
	)
}

func init() {
createCmd.Flags().BoolVarP(&createtables, "createtables", "a", false, "Create the tables if they do not exist")
createCmd.Flags().BoolVarP(&testprod, "testprod", "b", false, "create test product")
createCmd.Flags().BoolVarP(&testequip, "testequip", "d", false, "create test equipment")
createCmd.Flags().StringVarP(&createpartno, "createpartno", "n", "", "Create a part by providing the part number")
createCmd.Flags().IntVarP(&createseriesp, "createseriesp", "y", 0, "Create products with sequential part numbers")
createCmd.Flags().IntVarP(&createseriese, "createseriese", "z", 0, "Create equipments with sequential part numbers")
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "add parts to the database",
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
			//  create tables  //
			if createtables {
					inv.CreateProductsTableIfNotExists(sess)
					inv.CreateEquipmentsTableIfNotExists(sess)
			}
			//  Create Test Product  //
			if testprod {
				inv.DefineProducts(sess)
				inv.CreateTestProd(sess)
			}
			//  Create Test Product  //
			if testequip {
				inv.DefineEquipments(sess)
				inv.CreateTestEquip(sess)
			}
			//  create part  //
			if createpartno != "" {
				fmt.Println("createpart has value ", createpartno)
				inv.CreateProduct(sess, createpartno)
			}

			if createseriesp != 0 {
				fmt.Printf("creating %d products\n", createseriesp)
				inv.CreateSeriesP(sess, createseriesp)
			}
			if createseriese != 0 {
				fmt.Printf("creating %d equipments\n", createseriese)
				inv.CreateSeriesE(sess, createseriese)
			}


	},
}
