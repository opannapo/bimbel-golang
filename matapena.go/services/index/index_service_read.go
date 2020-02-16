package services

import "app/domain"

type IndexServiceReadImpl struct{}

// Home ...
func (IndexServiceReadImpl) Home() *domain.Index {
	explorer := []string{
		"/me",
		"/directory",
		"/help",
	}
	return &domain.Index{Message: "Welcome", Explore: explorer}
}

// Me ...
func (IndexServiceReadImpl) Me() *domain.Index {
	res := domain.Index{Message: "I'am API"}
	return &res
}

// Help ...
func (IndexServiceReadImpl) Help() *domain.Index {
	explorer := []string{
		"/me",
		"/directory",
		"/help",
	}
	res := domain.Index{Message: "Help", Explore: explorer}
	return &res
}

// Directory ...
func (IndexServiceReadImpl) Directory() *domain.Index {
	explorer := []string{
		"/",
		"/siswa",
		"/siswa/{ID}",
		"/employee",
		"/employee/{ID}",
	}
	res := domain.Index{Message: "Endpoint Directory", Explore: explorer}
	return &res
}
