package domain
// Created by https://github.com/opannapo/db-to-struct

// TblAddress domain @Auto-Generate
type TblAddress struct {
	AddrID      int    `gorm:"primary_key" json:"addr_id"`
	AddrAddress string `json:"addr_address"`
	AddrArea    string `json:"addr_area"`
	AddrPhone   string `json:"addr_phone"`
	AddrUsrID   int    `json:"addr_usr_id"`
}

// TableName for tbl_address @Auto-Generate
func (TblAddress) TableName() string {
	return "tbl_address"
}
