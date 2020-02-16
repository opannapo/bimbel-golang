package base

import "github.com/jinzhu/gorm"

type Repository struct {
	DB *gorm.DB
}

func Init() Repository {
	return Repository{}
}
