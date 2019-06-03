package configuracion

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Configuracion struct {
	Ip                               string
	Namedb                           string
	Passdb                           string
	Protocolomonolitico              string
	Dominiomonolitico                string
	Puertomonolitico                 string
	Protocolomicroservicio           string
	Dominiomicroservicio             string
	Puertomicroservicio              string
	Puertomicroservicioautenticacion string
	Puertomicroserviciolegajo        string
	Puertomicroservicionovedad       string
	Puertomicroservicioliquidacion   string
	Puertomicroservicioconcepto      string
	Puertomicroserviciohelpers       string
	Checkmonolitico                  bool
	Versionlegajo                    int
	Versionnovedad                   int
	Versionliquidacion               int
	Versionconcepto                  int
	Versionsecurity                  int
	Privacidadtablas                 []PrivacidadTablas
}

type PrivacidadTablas struct {
	Tabla      string
	Privacidad string
}

var instance Configuracion

func GetInstance() Configuracion {

	if instance.Ip == "" || instance.Namedb == "" {
		obtenerDatosConfiguracion()
	}

	return instance
}

func obtenerTablaPrivada(x string) PrivacidadTablas {

	for i, n := range instance.Privacidadtablas {
		if x == n.Tabla {
			return instance.Privacidadtablas[i]
		}
	}
	return instance.Privacidadtablas[0]
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

func GetUrlMonolitico() string {

	url := instance.Protocolomonolitico + "://" + instance.Dominiomonolitico + ":" + instance.Puertomonolitico + "/NXV/"
	return url
}

func GetUrlMicroservicio(microservicioPuerto string) string {
	puerto := instance.Puertomicroservicio
	if puerto == "" {
		puerto = microservicioPuerto
	}

	url := instance.Protocolomicroservicio + "://" + instance.Dominiomicroservicio + ":" + puerto + "/api/"
	return url
}
