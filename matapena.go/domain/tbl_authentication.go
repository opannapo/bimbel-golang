package domain
// Created by https://github.com/opannapo/db-to-struct

// TblAuthentication domain @Auto-Generate
type TblAuthentication struct {
  AuthID int `gorm:"primary_key" json:"auth_id"`
  UsrID int `json:"usr_id"`
  AuthEmail string `json:"auth_email"`
  AuthPasswordMd5 string `json:"auth_password_md5"`
  AuthPasswordPlain string `json:"auth_password_plain"`
}


// TableName for tbl_authentication @Auto-Generate
func (TblAuthentication) TableName() string {
  return "tbl_authentication"
}