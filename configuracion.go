package framework

type configuracion struct {
	ip     string
	namedb string
}

var instance configuracion

func GetInstance() configuracion {

	instance = configuracion{} // <--- NOT THREAD SAFE

	return instance
}

func GetIP() string {
	return "192.168.30.111"
}

func GetNameDB() string {
	return "faf_multitenant_go"
}
