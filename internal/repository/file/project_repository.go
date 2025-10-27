package file

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"github.com/sunsetsavorer/web-dashboard-panel/internal/entity"
)

type ProjectRepository struct {
	mu          sync.Mutex
	filePath    string
	projectsMap map[int]entity.Project
}

func NewProjectRepository(filePath string) *ProjectRepository {

	return &ProjectRepository{
		filePath:    filePath,
		mu:          sync.Mutex{},
		projectsMap: make(map[int]entity.Project),
	}
}

func (r *ProjectRepository) GetList() ([]entity.Project, error) {

	r.mu.Lock()
	defer r.mu.Unlock()

	file, err := os.OpenFile(r.filePath, os.O_RDONLY, 0666)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&r.projectsMap)
	if err != nil {
		return nil, fmt.Errorf("failed to decode projects file: %v", err)
	}

	projects := make([]entity.Project, 0, len(r.projectsMap))

	for _, project := range r.projectsMap {
		projects = append(projects, project)
	}

	return projects, nil
}
