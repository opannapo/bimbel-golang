package repository

import (
	"app/domain"
	"github.com/jinzhu/gorm"
)

type EmployeeRepo struct {
	Db *gorm.DB
}

func (r *EmployeeRepo) CreateAdmin(employee *domain.TblEmployee) (result *domain.TblEmployee, err error) {
	err = r.Db.Create(&employee).Error
	if err != nil {
		return nil, err
	}
	return employee, err
}

func (r *EmployeeRepo) GetAdminAll() (result []*domain.TblEmployee, err error) {
	err = r.Db.
		Preload("User").
		Preload("Address").
		Preload("Phone").
		Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, err
}

func (r *EmployeeRepo) GetAdminById(id int) (result domain.TblEmployee, err error) {
	err = r.Db.
		Preload("User").
		Preload("Address").
		Preload("Phone").
		Where("empl_id=?", id).
		Find(&result).Error
	if err != nil {
		return
	}
	return result, err
}
