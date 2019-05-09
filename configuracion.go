package framework

type configuracion struct {
	ip     string
	namedb string
}

var instance configuracion

func GetInstance() configuracion {

	instance = configuracion{} // <--- NOT THREAD SAFE
	instance.ip = getIP()
	instance.namedb = getNameDB()
	return instance
}

func getIP() string {
	return "192.168.30.111"
}

func getNameDB() string {
	return "faf_multitenant_go"
}
