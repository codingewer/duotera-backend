package models

import "github.com/jinzhu/gorm"

type Detail struct {
	gorm.Model
	Name                string ` gorm:"unique" json:"name"`
	AboutUs             string `json:"aboutus"`
	Phone               string `json:"phone"`
	Email               string `json:"email"`
	Address             string `json:"address"`
	MarkerUrl           string `json:"markerurl"`
	HomeVideo           string `json:"homeVideo"`
	ContactPageImageUrl string `json:"contactPageImageUrl"`
	HomeSubTitle        string `json:"homeSubTitle"`
	HomeTitle           string `json:"homeTitle"`
}

func (d *Detail) CreateDetails() (*Detail, error) {
	err := db.Create(&d).Error
	if err != nil {
		return &Detail{}, err
	}
	return d, nil
}

// get detail by name
func (d *Detail) GetDetailByName(name string) (*Detail, error) {
	err := db.Debug().Table("details").Where("name = ?", name).Take(&d).Error
	if err != nil {
		return &Detail{}, err
	}
	return d, nil
}

// update detail by name
func (d *Detail) UpdateDetailByName(name string) (*Detail, error) {
	err := db.Debug().Table("details").Where("name = ?", name).Updates(&d).Error
	if err != nil {
		return &Detail{}, err
	}
	return d, nil
}
