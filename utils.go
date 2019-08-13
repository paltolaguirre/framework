package framework

import (
	"net/http"
)

func checkRegistroDefault(p_id int, w http.ResponseWriter) {

	if p_id < 0 {
		RespondError(w, http.StatusNotFound, Eliminarmodificarregistrosdefaults)
		return
	}
}

func checkParametroVacio(p_id int, w http.ResponseWriter) {

	if p_id == 0 {
		RespondError(w, http.StatusNotFound, IdParametroVacio)
		return
	}
}
