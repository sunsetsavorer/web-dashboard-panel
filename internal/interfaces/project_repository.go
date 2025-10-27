package interfaces

import "github.com/sunsetsavorer/web-dashboard-panel/internal/entity"

type ProjectRepository interface {
	GetList() ([]entity.Project, error)
}
