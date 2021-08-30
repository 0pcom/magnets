package commands

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/skycoin/skycoin/src/util/logging"
		"github.com/sirupsen/logrus"
		inv "github.com/0pcom/magnets/cmd/magnets/commands/inv"
		run "github.com/0pcom/magnets/cmd/magnets/commands/run"
		"net/http"
		"github.com/gorilla/mux"
		_ "github.com/0pcom/magnets/statik"
)

const port = 8022
const siteURL = "https://magnetosphere.net"
//var createpartno string


var (
	runapp	bool
	createtables	bool
	deleteproducts	bool
	deleteequipments	bool
	droptable	bool
	droptable1	bool
	testprod	bool
	testequip	bool
	Createpartno	string
	create100p	bool
	create100e	bool
	create5000p	bool
	create5000e	bool
	importcsv	bool
	importcsv1	bool
	printinv	bool
	printinv1	bool
	vprintinv	bool
	exportcsv	bool
	exportcsv1	bool
	helpmenu	bool
)


func init() {
	RootCmd.CompletionOptions.DisableDefaultCmd = true

	RootCmd.AddCommand(
		inv.RootCmd,
		run.RootCmd,
	)
}

//
func init() {
	/*
RootCmd.Flags().BoolVarP(&droptable, "droptables", "d", false, "Drop products table")
RootCmd.Flags().BoolVarP(&droptable1, "droptables1", "b", false, "Drop equipments table")
RootCmd.Flags().BoolVarP(&deleteproducts, "deleteall", "D", false, "Delete the products in the products database")
RootCmd.Flags().BoolVarP(&deleteequipments, "deleteall1", "E", false, "Delete the equipment in the equipments database")
RootCmd.Flags().BoolVarP(&createtables, "createtables", "c", false, "Create the tables if they do not exist")
RootCmd.Flags().BoolVarP(&testprod, "testprod", "t", false, "create test product")
RootCmd.Flags().BoolVarP(&testequip, "testequip", "u", false, "create test equipment")
RootCmd.Flags().StringVarP(&Createpartno, "createpartno", "C", "", "Create a part by providing the part number")
RootCmd.Flags().BoolVarP(&create5000p, "create5000p", "Y", false, "Create 5000 products with sequential part numbers")
RootCmd.Flags().BoolVarP(&create5000e, "create5000e", "Z", false, "Create 5000 equipments with sequential part numbers")
RootCmd.Flags().BoolVarP(&create100p, "create100p", "y", false, "Create 100 products with sequential part numbers")
RootCmd.Flags().BoolVarP(&create100e, "create100e", "z", false, "Create 100 equipments with sequential part numbers")
RootCmd.Flags().BoolVarP(&exportcsv, "exportcsv", "e", false, "Export products to productsexport01.csv")
RootCmd.Flags().BoolVarP(&exportcsv1, "exportcsv1", "f", false, "Export equipments to equipmentsexport01.csv")
RootCmd.Flags().BoolVarP(&printinv, "printinv", "p", false, "Print the products table to the terminal")
RootCmd.Flags().BoolVarP(&printinv1, "printinv1", "q", false, "Print the equipments table to the terminal")
//RootCmd.Flags().BoolVarP(&printinv2, "printinv2", "q", false, "Print the  inventory - BTCPayserver format")
RootCmd.Flags().BoolVarP(&vprintinv, "vprintinv", "v", false, "More verbose printinventory")
RootCmd.Flags().BoolVarP(&importcsv, "importcsv", "i", false, "Import products csv from http://127.0.0.1:8079/export01.csv")
RootCmd.Flags().BoolVarP(&importcsv1, "importcsv1", "j", false, "Import equipments csv from http://127.0.0.1:8079/export01.csv")
RootCmd.Flags().BoolVarP(&runapp, "run", "r", false, "run the web app")
RootCmd.Flags().BoolVarP(&helpmenu, "help", "h", false, "show this help menu")
//RootCmd.Flags().Parse()
*/
}
//
var RootCmd = &cobra.Command{
	Use:   "magnets",
	Short: "magnetosphere.net website implementation",
//	PreRun: func(_ *cobra.Command, _ []string) {
//	},
	Run: func(_ *cobra.Command, _ []string) {

		mLog := logging.NewMasterLogger()
		mLog.SetLevel(logrus.InfoLevel)
		// // database connection // //
//		settings := user.FetchSettings()
//		sess := mdb.Connect(settings)
/*
//
			// // drop tables // //
//
			if droptable {
					inv.DropProductsTable(sess)

				}
			// // drop tables // //
			if droptable1 {
					inv.DropEquipmentsTable(sess)

				}
			// // create tables // //
			if createtables {
					inv.CreateProductsTableIfNotExists(sess)
					inv.CreateEquipmentsTableIfNotExists(sess)

			}
			// // delete products // //
			if deleteproducts {
				inv.DefineProducts(sess)
				inv.DeleteAllProducts(sess)

			}
			// // delete products // //
			if deleteequipments {
				inv.DefineEquipments(sess)
				inv.DeleteAllEquipments(sess)

			}
			// // Create Test Product // //
			if testprod {
				inv.DefineProducts(sess)
				inv.CreateTestProd(sess)

			}
			// // Create Test Product // //
			if testequip {
				inv.DefineEquipments(sess)
				inv.CreateTestEquip(sess)

			}
			// // create part // //
			if Createpartno != "" {
				fmt.Println("createpart has value ", Createpartno)
				inv.CreateProduct(sess, Createpartno)

			}
			// // Create 100 Product // //
			if create5000p {
				fmt.Println("creating 5000 products")
				inv.Create5kP(sess)

			}
			// // Create 100 Equipment // //
			if create5000e {
				fmt.Println("creating 5000 equipments")
				inv.Create5kE(sess)

			}
			if create100p {
				fmt.Println("creating 100 products")
				inv.CreateHundredP(sess)

			}
			// // Create 100 Equipment // //
			if create100e {
				fmt.Println("creating 100 equipments")
				inv.CreateHundredE(sess)

			}
			// // print products table inventory // //
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
			// // print equipments table inventory// //
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
			// // print inventory so it can be pasted in the btcpayserver raw editor! // //
			//
			if printinv2 {
				inv.DefineProducts(sess)
				log.Printf("products:")
				for i := range products {
					if inv.Mproducts[i].Enable {
						fmt.Printf("%s:\n", inv.Mproducts[i].PartNo)
						fmt.Printf("  price: "); fmt.Printf("%.2f\n", inv.Mproducts[i].Price)
						fmt.Printf("  title:	");	fmt.Printf("%s\n", inv.Mproducts[i].Name)
						fmt.Printf("  description: "); fmt.Printf("%s\n", inv.Mproducts[i].Description1)
						fmt.Printf("  image:	");	fmt.Printf("%s/img/%s\n", siteURL, inv.Mproducts[i].Image1)
						}
					}

				}
				//
			// // more verbosely print inventory // //
			if vprintinv {
				inv.DefineProducts(sess)
				log.Printf("products:")
				for i := range inv.Mproducts {
					fmt.Printf("product #%d: %#v\n", i, inv.Mproducts[i])

						}
					}
					// // export products table to csv RELIES ON MAKEFILE - COCKROACHDB COMMANDS // //
					if exportcsv {
						fmt.Println("Exporting csv to productsexport01.csv")
						exportProductsCSV()

					}
					// // export equipments table csv RELIES ON MAKEFILE - COCKROACHDB COMMANDS // //
					if exportcsv1 {
						fmt.Println("Exporting csv to productsexport01.csv")
						exportEquipmentsCSV()

					}
					// // import csv to products table // //
					if importcsv {
							fmt.Println("Import a csv into products table from http://127.0.0.1:8079/productsexport01.csv")
							importProductsCSV(sess)

					}
					// // import csv to equipments table // //
					if importcsv1 {
							fmt.Println("Import a csv into equipments table from http://127.0.0.1:8079/equipmentsexport01.csv")
							importEquipmentsCSV(sess)
//
					}
			// // run the web app // //
			if runapp {
				statikFS, err := fs.New()
				if err != nil {
					log.Fatal(err)
				}
				inv.DefineProducts(sess)
				inv.DefineEquipments(sess)
				r := mux.NewRouter() //.StrictSlash(true)
		    r.NotFoundHandler = page404()
				r.PathPrefix("/img/").Handler(http.StripPrefix("/img/", http.FileServer(http.Dir("./img")))) //images
				//products table - main site original endpoints
				r.Handle("/list", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(listPage))).Methods("GET") //pagination
				r.Handle("/list/", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(listPage))).Methods("GET") //pagination
				r.Handle("/list/{id:[0-9]+}", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(listPage))).Methods("GET") //pagination
				r.Handle("/post/{slug}", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(findProduct))).Methods("GET")	//individual product page
				r.Handle("/cat/{slug}", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(categoryPage))).Methods("GET")	//category
				r.Handle("/cat/{slug}/{id:[0-9]+}", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(categoryPage))).Methods("GET")	//category pagination
				// equipment table endpoints
				r.Handle("/equipment", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(equipmentPage))).Methods("GET")	//equipment
				r.Handle("/equipment/p/{id:[0-9]+}", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(equipmentPage))).Methods("GET")	//equipment pagination
				r.Handle("/equipment/post/{slug}", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(findEquipment))).Methods("GET")	//individual product page
				r.Handle("/equipment/cat/{slug}", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(categoryEquipmentPage))).Methods("GET")	//category
				r.Handle("/equipment/cat/{slug}/{id:[0-9]+}", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(categoryEquipmentPage))).Methods("GET")	//category pagination
				//single pages
				r.Handle("/about", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(aboutPage))).Methods("GET")	//about page
				r.Handle("/blog", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(blogPage))).Methods("GET")	//about page
				r.Handle("/blog/", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(blogPage))).Methods("GET")	//about page
				r.Handle("/blog/{slug}/", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(blogPage))).Methods("GET")	//about page
				r.Handle("/friend", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(friendPage))).Methods("GET")	//friends page
				r.Handle("/friend/", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(friendPage))).Methods("GET")	//friends page
				r.Handle("/policy", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(policyPage))).Methods("GET")	//shipping page
				r.Handle("/policy/", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(policyPage))).Methods("GET")	//shipping page
				//r.Handle("/time", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(timeFunc))).Methods("GET")	//shipping page
				//
		    r.Handle("/{slug}", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(frontPage))).Methods("GET") //site root
		    r.Handle("/", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(frontPage))).Methods("GET") //site root
		    r.Handle("/p/{id:[0-9]+}", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(pagePage))).Methods("GET") //pagination
				r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(statikFS)))	//statik sources
		    r.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		      // this handler behaved as I would expect the r.NotFoundHandler to behave..
		      w.WriteHeader(501)
		      w.Write([]byte(`{"status":501,"message":"501: Not implemented."}`))
		    })
				Serve = r
				fmt.Printf("listening on http://127.0.0.1:%d using gorilla router\n", port)
				log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), Serve))
//
			}
//		if cmdargs == 0 { helpmenu1() }


//
*/
	},
}









//package gorilla	// Routing based on the gorilla/mux router
var Serve http.Handler
// //  // //
func NewRouter() *mux.Router {
router := mux.NewRouter().StrictSlash(true)
staticDir := "/static/"	// Choose the folder to serve
router.	// Create the route
	PathPrefix(staticDir).
	Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("."+staticDir))))
return router
}
// // database stuff // //
var err error

// Execute executes root CLI command.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Fatal("Failed to execute command: ", err)
	}
}
