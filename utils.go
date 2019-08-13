package framework

import (
	"net/http"
)

func CheckRegistroDefault(p_id int, w http.ResponseWriter) {

	if p_id < 0 {
		RespondError(w, http.StatusNotFound, Eliminarmodificarregistrosdefaults)
		return
	}
}

func CheckParametroVacio(p_id int, w http.ResponseWriter) {

	if p_id == 0 {
		RespondError(w, http.StatusNotFound, IdParametroVacio)
		return
	}
}
