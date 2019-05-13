package configuracion

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuracion struct {
	Ip     string
	Namedb string
	Url    string
}

var instance Configuracion

func GetInstance() Configuracion {

	if instance.Ip == "" || instance.Namedb == "" {
		obtenerDatosConfiguracion()
	}

	return instance
}

func obtenerDatosConfiguracion() {

	//comento para poder trabajar con el debug
	/*path, err := os.Executable()

	if err != nil {
		fmt.Println("error: ", err)
	}*/

	dir := "/home/paltolaguirre/go/src/github.com/xubiosueldos" //filepath.Dir(path)

	filePath := dir + "/configuracion.json"

	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("error: ", err)
	}

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&instance)
	if err != nil {
		fmt.Println("error:", err)
	}
}
