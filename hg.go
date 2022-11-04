package main

import (
	"github.com/qrasmont/hourglass/app"
	"github.com/qrasmont/hourglass/data/project"
	"github.com/qrasmont/hourglass/data/record"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func openDataBase() *gorm.DB {

	db, err := gorm.Open(sqlite.Open("hourglass.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&project.Project{}, &record.Record{})

	return db
}

func main() {
	db := openDataBase()
	project := project.GormRepository{DB: db}
	record := record.GormRepository{DB: db}

	app.Start(project, record)
}
