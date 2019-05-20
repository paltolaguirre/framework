package versiondbmicroservicio

import (
	"github.com/jinzhu/gorm"
)

type Versiondbmicroservicio struct {
	Nombremicroservicio  string
	Versionmicroservicio string
}

func CrearTablaVersionDBMicroservicio(db *gorm.DB) {

	db.Model(&Versiondbmicroservicio{})
}

func UltimaVersion(nombre string, db *gorm.DB) string {

	var versiondbmicroservicio Versiondbmicroservicio

	db.First(&versiondbmicroservicio, "nombremicroservicio = ", nombre)

	return versiondbmicroservicio.Versionmicroservicio

}
