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
	result := g.DB.Where(&Record{Day: date.Day(), Month: date.Month(), Year: date.Year()}).Find(&records)
	return records, result.Error
}

func (g *GormRepository) GetRecordsForMonth(date time.Time) ([]Record, error) {
	var records []Record
	result := g.DB.Where(&Record{Month: date.Month(), Year: date.Year()}).Find(&records)
	return records, result.Error
}

func (g *GormRepository) GetRecordsForYear(date time.Time) ([]Record, error) {
	var records []Record
	result := g.DB.Where(&Record{Year: date.Year()}).Find(&records)
	return records, result.Error
}

func (g *GormRepository) GetRecordForProjectForDay(projectID uint, date time.Time) (Record, error) {
	var record Record
	result := g.DB.Where(&Record{ProjectID: projectID, Day: date.Day(), Month: date.Month(), Year: date.Year()}).Find(&record)
	return record, result.Error
}

func (g *GormRepository) GetRecordsForProjectForMonth(projectID uint, date time.Time) ([]Record, error) {
	var records []Record
	result := g.DB.Where(&Record{ProjectID: projectID, Month: date.Month(), Year: date.Year()}).Find(&records)
	return records, result.Error
}

func (g *GormRepository) GetRecordsForProjectForYear(projectID uint, date time.Time) ([]Record, error) {
	var records []Record
	result := g.DB.Where(&Record{ProjectID: projectID, Year: date.Year()}).Find(&records)
	return records, result.Error
}
