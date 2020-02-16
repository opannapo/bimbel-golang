package services

import "app/domain"

type IndexServiceModifyImpl struct{}

func (IndexServiceModifyImpl) HomeCreate() *domain.Index {
	explorer := []string{
		"/me",
		"/directory",
		"/help",
	}
	return &domain.Index{Message: "Welcome", Explore: explorer}
}

func (IndexServiceModifyImpl) MeCreate() *domain.Index {
	res := domain.Index{Message: "I'am API"}
	return &res
}

func (IndexServiceModifyImpl) HelpCreate() *domain.Index {
	explorer := []string{
		"/me",
		"/directory",
		"/help",
	}
	res := domain.Index{Message: "Help", Explore: explorer}
	return &res
}

func (IndexServiceModifyImpl) DirectoryCreate() *domain.Index {
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
