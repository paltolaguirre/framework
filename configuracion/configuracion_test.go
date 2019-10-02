package configuracion

import "testing"

func Test_GetConfig(t *testing.T) {
	t.Run("Check Ip", func(t *testing.T) {
		expected := "192.168.30.116"
		got := GetInstance()

		if expected != got.Ip {
			t.Errorf("got '%s' expected '%s'", got.Ip, expected)
		}
	})

	t.Run("Check Namedb", func(t *testing.T) {
		expected := "FAF_MULTITENANT_GO"
		got := GetInstance()

		if expected != got.Namedb {
			t.Errorf("got '%s' expected '%s'", got.Namedb, expected)
		}
	})
}

func Test_GetUrlMonolitico(t *testing.T) {
	t.Run("Check GetUrlMonolitico", func(t *testing.T) {
		expected := "https://192.168.30.112:8443/NXV/"
		got := GetUrlMonolitico()

		if expected != got {
			t.Errorf("got '%s' expected '%s'", got, expected)
		}
	})
}

func Test_GetUrlMicroservicio(t *testing.T) {
	t.Run("Check GetUrlMicroservicio", func(t *testing.T) {
		expected := "http://localhost:80/api/"
		got := GetUrlMicroservicio("80")

		if expected != got {
			t.Errorf("got '%s' expected '%s'", got, expected)
		}
	})
}
