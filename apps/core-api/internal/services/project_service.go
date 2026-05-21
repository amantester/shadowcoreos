package services

import (
	"github.com/amantester/shadowcoreos/apps/core-api/internal/models"
	"github.com/amantester/shadowcoreos/apps/core-api/internal/repositories"
)

type ProjectService struct {
	ProjectRepo *repositories.ProjectRepository
}

func NewProjectService(projectRepo *repositories.ProjectRepository) *ProjectService {
	return &ProjectService{
		ProjectRepo: projectRepo,
	}
}

func (s *ProjectService) Create(name, description, createdBy string) (*models.Project, error) {
	project := &models.Project{
		Name:        name,
		Description: description,
		CreatedBy:   createdBy,
	}

	err := s.ProjectRepo.Create(project)
	if err != nil {
		return nil, err
	}

	return project, nil
}

func (s *ProjectService) List() ([]models.Project, error) {
	return s.ProjectRepo.List()
}
