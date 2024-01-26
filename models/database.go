package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "CePxZSsHoN0PTmJIy2ZtzrCLdmknLCzL"
	dbname   = "duotera"
)

var db *gorm.DB

func init() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=require",
		host, port, user, password, dbname)

	fmt.Println(psqlInfo)
	conn, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	db = conn
	db.Debug().AutoMigrate(Admin{}, Product{}, ProductCarouselItem{}, SubProduct{}, ProductContainerContent{}, Dealership{}, Detail{})
	fmt.Println("Veri Tabanı bağlantısı başarılı!")

	admin := Admin{
		Username: "Admin",
		Password: "SDfgdfgfbnntyuty",
	}
	_, _ = admin.SaveToDb()

	detail := Detail{
		Name:                "Duotera",
		AboutUs:             "Douter",
		Phone:               "0555555555",
		Email:               "XXXXXXXXXXXXXXXX",
		Address:             "XXXXXXXXXXXXXXXX",
		Instagram:           "https://www.instagram.com/duotera.inc/",
		Facebook:            "https://facebook.com",
		Youtube:             "https://www.youtube.com/@duotera",
		MarkerUrl:           "https://www.github.com",
		HomeVideo:           "https://res.cloudinary.com/ddeatrwxs/video/upload/v1704919064/assets/Backgrounds/s2nvw7oakera1sqq3zwc.webm",
		HomeTitle:           "Model Y Accessories",
		HomeSubTitle:        "Model Y Accessories",
		ContactPageImageUrl: "https://res.cloudinary.com/ddeatrwxs/image/upload/v1704918968/assets/DuoteraProducts/BackSaveSet/qz3sxwx2tdvsoe9iegwv.jpg",
	}
	_, _ = detail.CreateDetails()
}
func GetDB() *gorm.DB {
	return db
}
