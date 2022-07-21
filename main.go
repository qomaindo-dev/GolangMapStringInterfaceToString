package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	dataMap := make(map[string]interface{})
	dataMap["nama"] = "Udin"
	dataMap["kelas"] = "Kelas karyawan"
	dataMap["detail_kelas"] = map[string]interface{}{"kelas_id": "KLS01", "nama_kelas": "VIP 7A"}
	fmt.Println("dataMap=> ", dataMap)

	//Proses convert map to string
	b, _ := json.Marshal(dataMap)
	convertB := string(b)
	fmt.Println("data b=> ", convertB)
	fmt.Println("tipe data b => ", reflect.TypeOf(convertB))
	//END

	stringFormat := ""
	stringFormat = "nomor_id:" + "ID01" + ", " + convertB
	fmt.Println("stringFormat=> ", stringFormat)
}
