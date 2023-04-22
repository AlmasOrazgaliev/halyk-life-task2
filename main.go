package main

import (
	"github.com/AlmasOrazgaliev/halyk-life-task2/config"
	"github.com/AlmasOrazgaliev/halyk-life-task2/controller"
	"github.com/AlmasOrazgaliev/halyk-life-task2/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	db, err := gorm.Open(
		postgres.Open(config.BD),
		&gorm.Config{})
	db.AutoMigrate(&models.Author{})

	db.AutoMigrate(&models.Book{})

	db.AutoMigrate(&models.Member{})

	db.AutoMigrate(&models.BookMember{})
	contr := controller.NewController(db)
	err = controller.Start(contr)
	if err != nil {
		log.Fatal(err)
	}
}
