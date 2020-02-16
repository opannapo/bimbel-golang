package services

import "app/domain"
 
type IndexServiceRead interface {
	Home() *domain.Index
	Me() *domain.Index
	Help() *domain.Index
	Directory() *domain.Index
}

type IndexServiceModify interface {
	HomeCreate() *domain.Index
	MeCreate() *domain.Index
	HelpCreate() *domain.Index
	DirectoryCreate() *domain.Index
}
