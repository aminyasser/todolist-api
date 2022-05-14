package database

import (
	"log"
     
	"github.com/aminyasser/todo-list/entity/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDb()  (db *gorm.DB, err error) {
	dsn := "host=localhost user=postgres password=root dbname=todolist port=5432"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return
	}



	db.AutoMigrate(&model.Task{})

	return db , nil
}

func CloseDb(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	dbSQL.Close()
}