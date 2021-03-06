package application

import (
	"github.com/Advanced-Go/Day-6/4-Database/easy-issues/domain"
)

type ProjectService struct {
	ProjectRepository domain.ProjectRepository
}

// Returns all the Projects
func (s ProjectService) Projects() ([]*domain.Project, error) {
	return s.ProjectRepository.All()
}

// Creates a Project
func (s ProjectService) Create(u *domain.Project) error {
	return s.ProjectRepository.Create(u)
}

// Deletes a Project
func (s ProjectService) Delete(id int64) error {
	return s.ProjectRepository.Delete(id)
}

// Get a Project by id
func (s ProjectService) Project(id int64) (*domain.Project, error) {
	return s.ProjectRepository.GetById(id)
}
