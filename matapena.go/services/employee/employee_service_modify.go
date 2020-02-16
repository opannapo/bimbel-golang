package services

import (
	"app/domain"
	"app/etc/param"
	"app/repository"
	"github.com/jinzhu/gorm"
	"time"
)

//EmployeeServiceModifyImpl ...
type EmployeeServiceModifyImpl struct {
	Repo        repository.EmployeeRepo
	RepoUser    repository.UserRepo
	RepoAddress repository.AddressRepo
	RepoPhone   repository.PhoneRepo
	Db          *gorm.DB
}

//CreateAdmin Employee Admin
func (e *EmployeeServiceModifyImpl) CreateAdmin(param param.AdminCreate) (result *domain.TblEmployee, err error) {
	//DB Objects
	var user *domain.TblUser
	var phone *domain.TblPhone
	employee := &param.Employee
	address := &param.Address

	//Initial User
	user = &domain.TblUser{
		UsrUtypID:      3,
		UsrStatus:      1,
		UsrCreatedDate: time.Now().Unix(),
		UsrUpdatedDate: time.Now().Unix(),
	}

	//Initial Phone
	phone = &domain.TblPhone{
		PhnLabel: "Primary",
		PhnNo:    employee.EmplPhone,
	}

	//Transactional begin
	tx := e.Db.Begin()

	//Create User
	user, err = e.RepoUser.Create(user)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	//Create Address
	address.AddrUsrID = user.UsrID
	address, err = e.RepoAddress.Create(address)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	//Create Phone
	phone.PhnUsrID = user.UsrID
	phone, err = e.RepoPhone.Create(phone)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	//Create Employee
	employee.EmplUsrID = user.UsrID
	employee, err = e.Repo.CreateAdmin(employee)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	//Transactional End Commit
	tx.Commit()

	return employee, nil
}

//CreateTeacher create employee teacher
func (e *EmployeeServiceModifyImpl) CreateTeacher(param param.AdminCreate) (result *domain.TblEmployee, err error) {
	panic("implement me")
}

//CreateGeneral create general employee
func (e *EmployeeServiceModifyImpl) CreateGeneral(param param.AdminCreate) (result *domain.TblEmployee, err error) {
	panic("implement me")
}

//InstanceServiceModify create instance
func InstanceServiceModify(db *gorm.DB) EmployeeServiceModify {
	return &EmployeeServiceModifyImpl{
		Db:          db,
		Repo:        repository.EmployeeRepo{Db: db},
		RepoUser:    repository.UserRepo{Db: db},
		RepoAddress: repository.AddressRepo{Db: db},
		RepoPhone:   repository.PhoneRepo{Db: db},
	}
}
