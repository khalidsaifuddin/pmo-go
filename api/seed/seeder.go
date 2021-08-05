package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/khalidsaifuddin/pmo/api/models"
)

func Load(db *gorm.DB) {
	err := db.Debug().AutoMigrate(&models.JenisPendaftaran{}, &models.Wilayah{}, &models.Registration{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&models.Registration{}).AddForeignKey("kode_provinsi", "wilayahs(kode)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	err = db.Debug().Model(&models.Registration{}).AddForeignKey("kode_kabupaten", "wilayahs(kode)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}
}
