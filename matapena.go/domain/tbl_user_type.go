package domain
// Created by https://github.com/opannapo/db-to-struct

// TblUserType domain @Auto-Generate
type TblUserType struct {
  UtypID int `gorm:"primary_key" json:"utyp_id"`
  Label string `json:"label"`
}


// TableName for tbl_user_type @Auto-Generate
func (TblUserType) TableName() string {
  return "tbl_user_type"
}