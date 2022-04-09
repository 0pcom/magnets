package inv

import (
	"github.com/spf13/cobra"
	//mdb "github.com/0pcom/magnets/pkg/db"
	user "github.com/0pcom/magnets/pkg/user"
	inv "github.com/0pcom/magnets/pkg/inv"
	"fmt"
	"log"
		"github.com/upper/db/v4/adapter/cockroachdb"
)

const port = 8022
const siteURL = "https://magnetosphere.net"
//var createpartno string


var (
	importCSV	bool
//	importCSV1	bool
	exportCSV	bool
//	exportCSV1	bool
)


func init() {
	RootCmd.AddCommand(
		impexpCmd,
	)
	impexpCmd.Flags().BoolVarP(&importCSV, "importp", "i", false, "Import products csv from http://127.0.0.1:8079/productsexport01.csv")
	impexpCmd.Flags().BoolVarP(&exportCSV, "exportp", "e", false, "Export products to productsexport01.csv")
	//impexpCmd.Flags().BoolVarP(&importCSV1, "importe", "j", false, "Import equipments csv from http://127.0.0.1:8079/equipmentsexport01.csv")
	//impexpCmd.Flags().BoolVarP(&exportCSV1, "exporte", "f", false, "Export equipments to equipmentsexport01.csv")
	//RootCmd.Flags().Parse()
}



var impexpCmd = &cobra.Command{
	Use:   "impexp",
	Short: "csv import and export operations",
	//Args:  cobra.MinimumNArgs(1),
	Run: func(_ *cobra.Command, _ []string) {
		if exportCSV {
			fmt.Println("Exporting csv to productsexport01.csv")
			inv.ExportProductsCSV()
		}
		/*
		if exportCSV1 {
			fmt.Println("Exporting csv to equipmentsexport01.csv")
			inv.ExportEquipmentsCSV()
		}
		*/
		if importCSV { //|| importCSV1 {
			// /* database connection */ //
			fmt.Printf("Initializing cockroachDB connection\n")
			sess, err := cockroachdb.Open(user.FetchSettings())		//establish the session
			if err != nil {
				log.Fatal("cockroachdb.Open: ", err)
			}
			defer sess.Close()
		if importCSV {
			// // import csv to products table // //
			fmt.Println("Import csv into products table from http://127.0.0.1:8079/productsexport01.csv")
			inv.ImportProductsCSV(sess)
		}
		/*
		if importCSV1 {
			fmt.Println("Import csv into equipments table from http://127.0.0.1:8079/equipmentsexport01.csv")
			inv.ImportEquipmentsCSV(sess)
		}
		*/
	}
	},
}
