package domain
// Created by https://github.com/opannapo/db-to-struct

// TblEmployee domain @Auto-Generate
type TblEmployee struct {
	EmplID        int          `gorm:"primary_key" json:"empl_id"`
	EmplFullname  string       `json:"empl_fullname"`
	EmplBirthdate int          `json:"empl_birthdate"`
	EmplGender    int          `json:"empl_gender"`
	EmplPhone     string       `json:"empl_phone"`
	EmplUsrID     int          `json:"empl_usr_id"`
	User          TblUser      `gorm:"foreignkey:EmplUsrID"json:"-"` //ignore json response
	Address       []TblAddress `gorm:"foreignkey:addr_usr_id;association_foreignkey:EmplUsrID"json:"address"`
	Phone         []TblPhone   `gorm:"foreignkey:phn_usr_id;association_foreignkey:EmplUsrID"json:"phone"`
}

// TableName for tbl_employee @Auto-Generate
func (TblEmployee) TableName() string {
	return "tbl_employee"
}
