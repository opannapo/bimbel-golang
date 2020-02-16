package services

import (
	"app/domain"
	"app/repository"
	"github.com/jinzhu/gorm"
)

type EmployeeServiceReadImpl struct {
	Repo repository.EmployeeRepo
}

func (e EmployeeServiceReadImpl) GetAdminById(id int) (result domain.TblEmployee, err error) {
	result, err = e.Repo.GetAdminById(id)
	if err != nil {
		println("Error  GetAdminById " + err.Error())
		return
	}
	return
}

func (e EmployeeServiceReadImpl) GetTeacherById(id int) *domain.TblEmployee {
	panic("implement me")
}

func (e EmployeeServiceReadImpl) GetGeneralById(id int) *domain.TblEmployee {
	panic("implement me")
}

func (e EmployeeServiceReadImpl) GetAdminAll() (result []*domain.TblEmployee, err error) {
	result, err = e.Repo.GetAdminAll()
	if err != nil {
		println("Error  GetAdminAll " + err.Error())
		return
	}
	return
}

func (e EmployeeServiceReadImpl) GetTeacherAll() []*domain.TblEmployee {
	panic("implement me")
}

func (e EmployeeServiceReadImpl) GetGeneralAll() []*domain.TblEmployee {
	panic("implement me")
}

func InstanceServiceRead(db *gorm.DB) EmployeeServiceRead {
	return EmployeeServiceReadImpl{Repo: repository.EmployeeRepo{Db: db}}
}
