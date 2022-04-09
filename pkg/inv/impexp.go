/*impexp.go*/

package inv

import(
	"fmt"
	//"github.com/upper/db/v4/adapter/cockroachdb"
	"github.com/upper/db/v4"

	"os/exec"
	"os"

)

//correct way from bash shell:
//cockroach sql --certs-dir=certs -e "SELECT * from product.products;" --format=csv > export01.csv
func ExportProductsCSV() { //the extremely lazy way
	fmt.Printf("Exporting products table to csv\n")
output, err := exec.Command("make", "export-products").CombinedOutput()
if err != nil {
  os.Stderr.WriteString(err.Error())
}
fmt.Println(string(output))
}
/*
func ExportEquipmentsCSV() { //the extremely lazy way
	fmt.Printf("Exporting equipments table to csv\n")
output, err := exec.Command("make", "export-equipments").CombinedOutput()
if err != nil {
  os.Stderr.WriteString(err.Error())
}
fmt.Println(string(output))
}
*/

//todo: improve importing
func ImportProductsCSV(sess db.Session) error {
	fmt.Printf("Importing CSV from http://127.0.0.1:8079/export01.csv\n")
	_, err := sess.SQL().Exec(`IMPORT INTO product.products CSV DATA ('http://127.0.0.1:8079/productsexport01.csv')	WITH skip = '1';`)
	if err != nil {
		return err
	}
	return nil
}
/*
func ImportEquipmentsCSV(sess db.Session) error {
	fmt.Printf("Importing CSV from http://127.0.0.1:8079/equipmentsexport01.csv\n")
	_, err := sess.SQL().Exec(`IMPORT INTO product.equipments CSV DATA ('http://127.0.0.1:8079/equipmentsexport01.csv')	WITH skip = '1';`)
	if err != nil {
		return err
	}
	return nil
}
*/
