package repository

import (
	"app/domain"
	"github.com/jinzhu/gorm"
)

type AddressRepo struct {
	Db *gorm.DB
}

func (r *AddressRepo) Create(domain *domain.TblAddress) (result *domain.TblAddress, err error) {
	err = r.Db.Create(&domain).Error
	if err != nil {
		return
	}
	return domain, err
}
