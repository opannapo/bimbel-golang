package repository

import (
	"app/domain"
	"github.com/jinzhu/gorm"
)

type PhoneRepo struct {
	Db *gorm.DB
}

func (r *PhoneRepo) Create(domain *domain.TblPhone) (result *domain.TblPhone, err error) {
	err = r.Db.Create(&domain).Error
	if err != nil {
		return
	}
	return domain, err
}
