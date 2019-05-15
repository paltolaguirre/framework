package framework

import (
	"net/http"

	"github.com/jinzhu/gorm"
)

func ServicioList(w http.ResponseWriter, r *http.Request, servicio interface{}, db *gorm.DB) {

	db.Find(&servicio)

}
