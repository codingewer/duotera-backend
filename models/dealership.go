package models

import "github.com/jinzhu/gorm"

type Dealership struct {
	gorm.Model
	Name        string `json:"name"`
	Address     string `json:"address"`
	Email       string `json:"email"`
	Website     string `json:"websites"`
	Description string `json:"description"`
	IsActive    bool   `json:"is_active"`
}

func (d *Dealership) Create() (*Dealership, error) {
	err := db.Create(&d).Error
	if err != nil {
		return &Dealership{}, err
	}
	return d, nil
}

func (d *Dealership) FindAll() (*[]Dealership, error) {
	var dealerships []Dealership
	err := db.Find(&dealerships).Error
	if err != nil {
		return &[]Dealership{}, err
	}
	return &dealerships, nil

}

func (d *Dealership) FindById(id uint64) (*Dealership, error) {
	err := db.First(&d, id).Error
	if err != nil {
		return &Dealership{}, err
	}
	return d, nil
}
func (d *Dealership) IsActiveChange(id uint64) (*Dealership, error) {
	err := db.Where("id = ?", id).First(&d).Error
	if err != nil {
		return &Dealership{}, err
	}
	err = db.Model(&Dealership{}).Where("id = ?", id).Update("is_active", !d.IsActive).Error
	if err != nil {
		return &Dealership{}, err
	}
	return d, nil
}

// Get by is active true
func (d *Dealership) FindByIsActive() (*[]Dealership, error) {
	var dealerships []Dealership
	err := db.Where("is_active = ?", false).Find(&dealerships).Error
	if err != nil {
		return &[]Dealership{}, err
	}
	return &dealerships, nil
}

// Get by is active false
func (d *Dealership) FindByIsActiveFalse() (*[]Dealership, error) {
	var dealerships []Dealership
	err := db.Where("is_active = ?", true).Find(&dealerships).Error
	if err != nil {
		return &[]Dealership{}, err
	}
	return &dealerships, nil
}
