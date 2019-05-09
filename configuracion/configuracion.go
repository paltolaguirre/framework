package configuracion

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Configuracion struct {
	Ip     string
	Namedb string
}

var instance Configuracion

func GetInstance() Configuracion {

	instance = Configuracion{}

	path, err := os.Executable()

	if err != nil {
		fmt.Println(err)
	}
	dir := filepath.Dir(path)

	filePath := dir + "/configuracion.json"

	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&instance)
	if err != nil {
		fmt.Println("error:", err)
	}

	return instance
}
