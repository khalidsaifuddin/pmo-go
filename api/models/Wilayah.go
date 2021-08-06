package models

import (
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Wilayah struct {
	Kode           string    `gorm:"primary_key" json:"kode"`
	Nama           string    `json:"nama"`
	IdLevelWilayah uint64    `json:"id_level_wilayah"`
	IndukKode      string    `json:"induk_kode"`
	Induk          Induk     `json:induk`
	CreatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	ExpiredAt      time.Time `gorm:"default:null" json:"expired_at"`
}

type Induk []Wilayah

func (r *Wilayah) Prepare() {
	r.Kode = ""
	r.Nama = html.EscapeString(strings.TrimSpace(r.Nama))
	r.CreatedAt = time.Now()
	r.UpdatedAt = time.Now()
}

func (r *Wilayah) FindAllWilayah(db *gorm.DB) (*[]Wilayah, error) {
	var err error
	wilayahs := []Wilayah{}
	err = db.Debug().Model(&Wilayah{}).Limit(100).Find(&wilayahs).Error
	if err != nil {
		return &[]Wilayah{}, err
	}
	return &wilayahs, nil
}

func (r *Wilayah) FindWilayahByID(db *gorm.DB, kode string) (*Wilayah, error) {
	var err error
	err = db.Debug().Model(&Wilayah{}).Where("kode = ?", kode).Take(&r).Error
	if err != nil {
		return &Wilayah{}, err
	}
	return r, nil
}

func (r *Wilayah) FindWilayahByInduk(db *gorm.DB, id_level_wilayah uint64, induk_kode string) (*[]Wilayah, error) {
	var err error
	// err = db.Debug().Model(&Wilayah{}).Where("id_level_wilayah = ?", id_level_wilayah).Where("induk_kode = ?", induk_kode).Take(&r).Error
	wilayahs := []Wilayah{}
	if induk_kode != "" {
		err = db.Debug().Model(&Wilayah{}).Where("id_level_wilayah = ?", id_level_wilayah).Where("induk_kode = ?", induk_kode).Limit(100).Order("nama asc").Find(&wilayahs).Error
	} else {
		err = db.Debug().Model(&Wilayah{}).Where("id_level_wilayah = ?", id_level_wilayah).Limit(100).Order("nama asc").Find(&wilayahs).Error
	}
	if err != nil {
		return &[]Wilayah{}, err
	}
	if len(wilayahs) > 0 {
		for i, _ := range wilayahs {

			if wilayahs[i].IndukKode != "" {
				err := db.Debug().Model(&Wilayah{}).Where("kode = ?", wilayahs[i].IndukKode).Take(&wilayahs[i].Induk).Error
				if err != nil {
					return &[]Wilayah{}, err
				}
			}

		}
	}
	return &wilayahs, nil
}
