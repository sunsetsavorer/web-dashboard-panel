package app

import (
	"fmt"
	"net/http"

	"github.com/sunsetsavorer/web-dashboard-panel/internal/config"
	"github.com/sunsetsavorer/web-dashboard-panel/internal/repository/file"
	"github.com/sunsetsavorer/web-dashboard-panel/internal/router"
	v1 "github.com/sunsetsavorer/web-dashboard-panel/internal/transport/http/v1"
	"github.com/sunsetsavorer/web-dashboard-panel/internal/usecase/project"
)

type App struct {
}

func New() *App {

	return &App{}
}

func (a *App) Run() error {

	config := config.New()

	err := config.Load("configs/config.toml")
	if err != nil {
		return fmt.Errorf("error with config: %v", err)
	}

	projectRepository := file.NewProjectRepository("temp/projects.json")

	getProjectsListUC := project.NewGetListUseCase(projectRepository)

	projectService := project.NewProjectService(
		getProjectsListUC,
	)

	baseHandler := v1.NewBaseHandler()

	projectHandler := v1.NewProjectHandler(
		baseHandler,
		projectService,
	)

	router := router.New()

	projectHandler.RegisterRoutes(router)

	err = http.ListenAndServe(config.Server.Address, router)
	if err != nil {
		return err
	}

	return nil
}
