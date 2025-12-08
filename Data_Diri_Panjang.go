package main
import "fmt"

func main() {
	dataUser := map[string]interface{}{
		"nama":		"Fatih",
		"umur":		17,
		"kota":		"Jakarta",
		"status":	false,
		"aktif":	false,
		"ipk":		00,
	}
konversiStatus := map[string]map[bool]string(
	"pernikahan": {
		true: "menikah",
		false: "Belum Menikah"
		}
	"keaktifan": {
		true: "aktif",
		false: "Tidak Aktif",
		}
	"kelulusan": {
		true: "lulus",
		false: "Belum lulus",
		}
	
	
)
}