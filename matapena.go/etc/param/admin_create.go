package param

import "app/domain"

type AdminCreate struct {
	Employee domain.TblEmployee `json:"employee" binding:"required"`
	Address  domain.TblAddress  `json:"address" binding:"required"`
}
