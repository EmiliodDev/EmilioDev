package models

import (
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Name        string `gorm:"type:varchar(100)"`
	Description string `gorm:"type:text"`
	URL         string `gorm:"type:varchar(255)"`
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&Project{})
}

func (p *Project) Create(db *gorm.DB) error {
	return db.Create(p).Error
}

func (p *Project) Delete(db *gorm.DB) error {
	return db.Delete(p).Error
}

func (p *Project) Update(db *gorm.DB) error {
	return db.Save(p).Error
}

func GetAllProjects(db *gorm.DB) ([]Project, error) {
	var projects []Project
	if err := db.Find(&projects).Error; err != nil {
		return nil, err
	}

	return projects, nil
}
