package record

import (
	"time"

	"gorm.io/gorm"
)

type Record struct {
	gorm.Model
	duration  time.Duration
	ProjectID uint `gorm:"foreignKey:Project"`
}

type GormRepository struct {
	DB *gorm.DB
}

type Repository interface {
	CreateRecord(duration time.Duration, projectID uint) error
	DeleteRecordForProject(projectID uint) error
	DeleteRecordById(id uint) error
	GetRecords(projectID uint) ([]Record, error)
}

func (g *GormRepository) CreateRecord(duration time.Duration, projectID uint) error {
	record := Record{duration: duration, ProjectID: projectID}

	result := g.DB.Create(&record)
	return result.Error
}

func (g *GormRepository) DeleteRecordForProject(projectID uint) error {
	result := g.DB.Where("project_id = ?", projectID).Delete(&Record{})
	return result.Error
}

func (g *GormRepository) DeleteRecodById(id uint) error {
	result := g.DB.Delete(&Record{}, id)
	return result.Error
}

func (g *GormRepository) GetRecords(projectID uint) ([]Record, error) {
	var records []Record

	result := g.DB.Where("project_id = ?", projectID).Find(&records)
	return records, result.Error
}
