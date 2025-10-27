package project

import (
	"github.com/sunsetsavorer/web-dashboard-panel/internal/contract"
)

type ProjectService struct {
	getListUC *GetListUseCase
}

func NewProjectService(
	getListUC *GetListUseCase,
) *ProjectService {

	return &ProjectService{
		getListUC: getListUC,
	}
}

func (s *ProjectService) GetList() (contract.GetProjectsList, error) {

	res, err := s.getListUC.Exec()
	if err != nil {
		return contract.GetProjectsList{}, err
	}

	return res, nil
}
