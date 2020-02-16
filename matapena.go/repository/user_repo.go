package repository

import (
	"app/domain"
	"github.com/jinzhu/gorm"
)

type UserRepo struct {
	Db *gorm.DB
}

func (r *UserRepo) Create(domain *domain.TblUser) (result *domain.TblUser, err error) {
	err = r.Db.Create(&domain).Error
	if err != nil {
		return
	}
	return domain, err
}
