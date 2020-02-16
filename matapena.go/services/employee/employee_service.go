package services

import (
	"app/domain"
	"app/etc/param"
)

// EmployeeServiceRead .
type EmployeeServiceRead interface {
	GetAdminById(id int) (result domain.TblEmployee, err error)
	GetTeacherById(id int) *domain.TblEmployee
	GetGeneralById(id int) *domain.TblEmployee
	GetAdminAll() (result []*domain.TblEmployee, err error)
	GetTeacherAll() []*domain.TblEmployee
	GetGeneralAll() []*domain.TblEmployee
}

// EmployeeServiceModify .
type EmployeeServiceModify interface {
	CreateAdmin(param param.AdminCreate) (result *domain.TblEmployee, err error)
	CreateTeacher(param param.AdminCreate) (result *domain.TblEmployee, err error)
	CreateGeneral(param param.AdminCreate) (result *domain.TblEmployee, err error)
}
