package configuracion

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Configuracion struct {
	Ip                 string
	Namedb             string
	Url                string
	Checkmonolitico    bool
	Versionlegajo      int
	Versionnovedad     int
	Versionliquidacion int
	Versionconcepto int
}

var instance Configuracion

func GetInstance() Configuracion {

	if instance.Ip == "" || instance.Namedb == "" {
		obtenerDatosConfiguracion()
	}

	return instance
}

func obtenerDatosConfiguracion() {

	path, err := os.Executable()

	if err != nil {
		fmt.Println("error: ", err)
	}

	dir := filepath.Dir(path)

	filePath := dir + "/configuracion.json"

	file, err := os.Open(filePath)

	if err != nil {
		filePath = "/configuracion.json"
		file, err = os.Open(filePath)
		if err != nil {
			fmt.Println("error: ", err)
		}
	}

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&instance)
	if err != nil {
		fmt.Println("error:", err)
	}
}
