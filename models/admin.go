package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Admin struct {
	gorm.Model
	Username string `gorm:"unique" json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type AdminLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AdminResponse struct {
	Username string `json:"username"`
	Token    uint   `json:"token"`
}

func (a *Admin) SaveToDb() (*Admin, error) {
	a.Role = "ADMIN"
	err := db.Debug().Create(&a).Error
	if err != nil {
		return &Admin{}, err
	}
	return a, nil
}

func (a *Admin) SignIn(admin Admin) (*Admin, error) {
	fmt.Println(admin)
	err := db.Debug().Where("username = ?", admin.Username).Take(&a).Error
	if err != nil {
		return &Admin{}, err
	}
	if a.Password != admin.Password {
		return &Admin{}, nil
	}
	if a.Role != "ADMIN" {
		return &Admin{}, nil
	}
	return a, nil
}

// find by userID
func (a *Admin) FindAdminByUserId(userId uint) (*Admin, error) {
	err := db.Debug().Where("id = ?", userId).Take(&a).Error
	if err != nil {
		return &Admin{}, err
	}
	return a, nil
}
