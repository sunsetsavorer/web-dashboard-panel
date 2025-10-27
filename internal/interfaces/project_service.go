package interfaces

import (
	"github.com/sunsetsavorer/web-dashboard-panel/internal/contract"
)

type ProjectService interface {
	GetList() (contract.GetProjectsList, error)
}
