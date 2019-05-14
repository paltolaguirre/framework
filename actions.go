package framework

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/xubiosueldos/autenticacion/publico"
	"github.com/xubiosueldos/conexionBD"
)

func ServicioList(w http.ResponseWriter, r *http.Request, servicio struct{}, db *gorm.DB) {

	tokenAutenticacion, tokenError := checkTokenValido(r)
	fmt.Println(tokenAutenticacion)
	if tokenError != nil {
		errorToken(w, tokenError)
		return
	} else {

		//db := obtenerDB(tokenAutenticacion)

		//defer db.Close()

		//Lista todos los legajos
		db.Find(&servicio)

		RespondJSON(w, http.StatusOK, servicio)
	}

}

func checkTokenValido(r *http.Request) (*publico.Security, *publico.Error) {

	var tokenAutenticacion *publico.Security
	var tokenError *publico.Error

	url := "http://localhost:8081/check-token"

	req, _ := http.NewRequest("GET", url, nil)

	header := r.Header.Get("Authorization")

	req.Header.Add("Authorization", header)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	if res.StatusCode != http.StatusBadRequest {

		// tokenAutenticacion = &(TokenAutenticacion{})
		tokenAutenticacion = new(publico.Security)
		json.Unmarshal([]byte(string(body)), tokenAutenticacion)

	} else {
		tokenError = new(publico.Error)
		json.Unmarshal([]byte(string(body)), tokenError)

	}

	return tokenAutenticacion, tokenError
}

func errorToken(w http.ResponseWriter, tokenError *publico.Error) {
	errorToken := *tokenError
	RespondError(w, errorToken.ErrorCodigo, errorToken.ErrorNombre)

}

func obtenerDB(tokenAutenticacion *publico.Security) *gorm.DB {

	token := *tokenAutenticacion
	tenant := token.Tenant

	return conexionBD.ConnectBD(tenant)

}
