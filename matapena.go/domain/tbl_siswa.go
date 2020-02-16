package domain
// Created by https://github.com/opannapo/db-to-struct

// TblSiswa domain @Auto-Generate
type TblSiswa struct {
  SwID int `gorm:"primary_key" json:"sw_id"`
  SwPrnID int `json:"sw_prn_id"`
  SwFullname string `json:"sw_fullname"`
  SwBirthdate int `json:"sw_birthdate"`
  SwGender int `json:"sw_gender"`
  SwSclID int `json:"sw_scl_id"`
  SwPhone string `json:"sw_phone"`
  SwPicture string `json:"sw_picture"`
  SwUsrID int `json:"sw_usr_id"`
}


// TableName for tbl_siswa @Auto-Generate
func (TblSiswa) TableName() string {
  return "tbl_siswa"
}