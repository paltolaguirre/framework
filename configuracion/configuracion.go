package configuracion

import (
	"fmt"
	"os"
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
	fmt.Println(path)
	instance.Ip = "192.168.30.111"
	instance.Namedb = "faf_multitenant_go"
	return instance
}
