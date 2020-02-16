package domain
// Created by https://github.com/opannapo/db-to-struct

// TblParent domain @Auto-Generate
type TblParent struct {
  PrnID int `gorm:"primary_key" json:"prn_id"`
  PrnFullname string `json:"prn_fullname"`
  PrnPhone string `json:"prn_phone"`
  PrnNik string `json:"prn_nik"`
  PrnBirthdate int `json:"prn_birthdate"`
  PrnGender int `json:"prn_gender"`
  PrnUsrID int `json:"prn_usr_id"`
}


// TableName for tbl_parent @Auto-Generate
func (TblParent) TableName() string {
  return "tbl_parent"
}