package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type JenisPendaftaran struct {
	ID        uint64    `gorm:"primary_key json:id`
	Nama      string    `json:nama`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	ExpiredAt time.Time `gorm:"default:null" json:"expired_at"`
}

func (r *JenisPendaftaran) FindAllJenisPendaftaran(db *gorm.DB) (*[]JenisPendaftaran, error) {
	var err error
	jenis_pendaftarans := []JenisPendaftaran{}
	err = db.Debug().Model(&JenisPendaftaran{}).Limit(100).Find(&jenis_pendaftarans).Error
	if err != nil {
		return &[]JenisPendaftaran{}, err
	}
	return &jenis_pendaftarans, nil
}
