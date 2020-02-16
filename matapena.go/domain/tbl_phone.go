package domain
// Created by https://github.com/opannapo/db-to-struct

// TblPhone domain @Auto-Generate
type TblPhone struct {
	PhnID    int    `gorm:"primary_key" json:"phn_id"`
	PhnUsrID int    `json:"phn_usr_id"`
	PhnLabel string `json:"phn_label"`
	PhnNo    string `json:"phn_no"`
}

// TableName for tbl_phone @Auto-Generate
func (TblPhone) TableName() string {
	return "tbl_phone"
}
