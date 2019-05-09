package configuracion

import (
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

	fmt.Println(path)
	fmt.Println("dir")
	fmt.Println(dir)
	instance.Ip = "192.168.30.111"
	instance.Namedb = "faf_multitenant_go"
	return instance
}
