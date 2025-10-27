package project

import (
	"github.com/sunsetsavorer/web-dashboard-panel/internal/contract"
	"github.com/sunsetsavorer/web-dashboard-panel/internal/interfaces"
)

type GetListUseCase struct {
	projectRepository interfaces.ProjectRepository
}

func NewGetListUseCase(
	projectRepository interfaces.ProjectRepository,
) *GetListUseCase {

	return &GetListUseCase{
		projectRepository: projectRepository,
	}
}

func (uc *GetListUseCase) Exec() (contract.GetProjectsList, error) {

	projects, err := uc.projectRepository.GetList()
	if err != nil {
		return contract.GetProjectsList{}, err
	}

	res := make([]contract.ProjectItem, 0, len(projects))

	for _, project := range projects {
		res = append(res, contract.ProjectItem{
			ID:        project.ID,
			Title:     project.Title,
			HexColor:  project.HexColor,
			CreatedAt: project.CreatedAt.Format("2006-01-02"),
		})
	}

	return contract.GetProjectsList{Data: res}, nil
}
