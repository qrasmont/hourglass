package project

import (
	"fmt"

	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Name string
}

type GormRepository struct {
	DB *gorm.DB
}

type Repository interface {
	GetProjects() ([]Project, error)
	CreateProject(name string) (Project, error)
	DeleteProject(projectID uint) error
}

func (g *GormRepository) GetProjects() ([]Project, error) {
	var projects []Project

	if err := g.DB.Find(&projects).Error; err != nil {
		return projects, fmt.Errorf("No projects: %v", err)
	}

	return projects, nil
}

func (g *GormRepository) CreateProject(name string) (Project, error) {
	project := Project{Name: name}

	if err := g.DB.Create(&project).Error; err != nil {
		return project, fmt.Errorf("Cloud not create project: %v", err)
	}

	return project, nil
}

func (g *GormRepository) DeleteProject(id uint) error {
	if err := g.DB.Delete(&Project{}, id).Error; err != nil {
		return fmt.Errorf("Could not delete project: %v", err)
	}

	return nil
}
