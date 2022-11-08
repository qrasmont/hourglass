package record

import (
	"time"

	"gorm.io/gorm"
)

type Record struct {
	gorm.Model
	Duration  time.Duration
	Day       int
	Month     time.Month
	Year      int
	ProjectID uint `gorm:"foreignKey:Project"`
}

type GormRepository struct {
	DB *gorm.DB
}

type Repository interface {
	CreateRecord(duration time.Duration, projectID uint) error
	UpdateRecord(record Record) error
	DeleteRecordForProject(projectID uint) error
	DeleteRecordById(id uint) error
	GetRecords(projectID uint) ([]Record, error)
	GetRecordsForDay(date time.Time) (Record, error)
	GetRecordsForMonth(date time.Time) ([]Record, error)
	GetRecordsForYear(date time.Time) ([]Record, error)
	GetRecordForProjectForDay(projectID uint, date time.Time) (Record, error)
	GetRecordsForProjectForMonth(projectID uint, date time.Time) ([]Record, error)
	GetRecordsForProjectForYear(projectID uint, date time.Time) ([]Record, error)
}

func (g *GormRepository) CreateRecord(duration time.Duration, projectID uint, date time.Time) error {
	record := Record{Duration: duration,
		ProjectID: projectID,
		Day:       date.Day(),
		Month:     date.Month(),
		Year:      date.Year(),
	}

	result := g.DB.Create(&record)
	return result.Error
}

func (g *GormRepository) UpdateRecord(record Record) error {
	result := g.DB.Save(&record)
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

func (g *GormRepository) GetRecordsForDay(date time.Time) ([]Record, error) {
	var records []Record
	result := g.DB.Where(
		"day = ? AND month = ? AND year = ?",
		date.Day(),
		date.Month(),
		date.Year()).
		Find(&records)
	return records, result.Error
}

func (g *GormRepository) GetRecordsForMonth(date time.Time) ([]Record, error) {
	var records []Record
	result := g.DB.Where(
		"month = ? AND year = ?",
		date.Month(),
		date.Year()).
		Find(&records)
	return records, result.Error
}

func (g *GormRepository) GetRecordsForYear(date time.Time) ([]Record, error) {
	var records []Record
	result := g.DB.Where(
		"year = ?",
		date.Year()).
		Find(&records)
	return records, result.Error
}

func (g *GormRepository) GetRecordForProjectForDay(projectID uint, date time.Time) (Record, error) {
	var record Record
	result := g.DB.Where(
		"project_id = ? AND day = ? AND month = ? AND year = ?",
		projectID,
		date.Day(),
		date.Month(),
		date.Year()).
		First(&record)
	return record, result.Error
}

func (g *GormRepository) GetRecordsForProjectForMonth(projectID uint, date time.Time) ([]Record, error) {
	var records []Record
	result := g.DB.Where(
		"project_id = ? AND month = ? AND year = ?",
		projectID,
		date.Month(),
		date.Year()).
		Find(&records)
	return records, result.Error
}

func (g *GormRepository) GetRecordsForProjectForYear(projectID uint, date time.Time) ([]Record, error) {
	var records []Record
	result := g.DB.Where(
		"project_id = ? AND year = ?",
		projectID,
		date.Year()).
		Find(&records)
	return records, result.Error
}
