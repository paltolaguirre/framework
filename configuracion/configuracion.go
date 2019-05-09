package configuracion

type Configuracion struct {
	ip     string
	namedb string
}

var instance Configuracion

func GetInstance() Configuracion {

	instance = Configuracion{}
	instance.ip = "192.168.30.111"
	instance.namedb = "faf_multitenant_go"
	return instance
}
