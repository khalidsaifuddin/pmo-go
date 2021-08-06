package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Registration struct {
	ID                 string           `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	JenisPendaftaranID uint32           `gorm:"not_null" json:"jenis_pendaftaran_id"`
	KodeProvinsi       string           `gorm:"not_null" json:"kode_provinsi"`
	KodeKabupaten      string           `gorm:"not_null" json:"kode_kabupaten"`
	Nama               string           `gorm:"not_null" json:"nama"`
	LembagaAsal        string           `json:"lembaga_asal"`
	Peran              string           `json:"peran"`
	NoHp               string           `json:"no_hp"`
	Email              string           `json:"email"`
	Keterangan         string           `json:"keterangan"`
	SoftDelete         uint64           `gorm:"default:0" json:"soft_delete"`
	CreatedAt          time.Time        `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt          time.Time        `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	Provinsi           Wilayah          `json:"provinsi"`
	Kabupaten          Wilayah          `json:"kabupaten"`
	JenisPendaftaran   JenisPendaftaran `json:"jenis_pendaftaran`
}

func (r *Registration) Prepare() {
	r.ID = ""
	r.Nama = html.EscapeString(strings.TrimSpace(r.Nama))
	r.CreatedAt = time.Now()
	r.UpdatedAt = time.Now()
}

func (r *Registration) SaveRegistration(db *gorm.DB) (*Registration, error) {
	var err error
	err = db.Debug().Model(&Registration{}).Create(&r).Error
	if err != nil {
		return &Registration{}, err
	}
	if r.ID != "" {
		err = db.Debug().Model(&Wilayah{}).Where("kode = ?", r.KodeProvinsi).Take(&r.Provinsi).Error
		if err != nil {
			return &Registration{}, err
		}
		if r.KodeKabupaten != "" {
			err = db.Debug().Model(&Wilayah{}).Where("kode = ?", r.KodeKabupaten).Take(&r.Kabupaten).Error
			if err != nil {
				return &Registration{}, err
			}
		}
	}
	return r, nil
}

func (r *Registration) Validate() error {
	if r.Nama == "" {
		return errors.New("required nama")
	}
	if r.KodeProvinsi == "" {
		return errors.New("required kode provinsi")
	}

	return nil
}

func (r *Registration) FindRegistrationByID(db *gorm.DB, id string) (*Registration, error) {
	var err error
	err = db.Debug().Model(&Registration{}).Where("id = ?", id).Take(&r).Error
	if err != nil {
		return &Registration{}, err
	}

	if r.JenisPendaftaranID != 0 {
		err = db.Debug().Model(&JenisPendaftaran{}).Where("id = ?", r.JenisPendaftaranID).Take(&r.JenisPendaftaran).Error
		if err != nil {
			return &Registration{}, err
		}
		err = db.Debug().Model(&Wilayah{}).Where("kode = ?", r.KodeProvinsi).Take(&r.Provinsi).Error
		if err != nil {
			return &Registration{}, err
		}
		if r.KodeKabupaten != "" {
			err = db.Debug().Model(&Wilayah{}).Where("kode = ?", r.KodeKabupaten).Take(&r.Kabupaten).Error
			if err != nil {
				return &Registration{}, err
			}
		}
	}

	return r, nil
}
