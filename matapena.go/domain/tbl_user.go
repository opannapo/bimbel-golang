package domain
// Created by https://github.com/opannapo/db-to-struct

// TblUser domain @Auto-Generate
type TblUser struct {
	UsrID          int   `gorm:"primary_key" json:"usr_id"`
	UsrUtypID      int   `json:"usr_utyp_id"`
	UsrCreatedDate int64 `json:"usr_created_date"`
	UsrStatus      int   `json:"usr_status"`
	UsrUpdatedDate int64 `json:"usr_updated_date"`
}

// TableName for tbl_user @Auto-Generate
func (TblUser) TableName() string {
	return "tbl_user"
}
