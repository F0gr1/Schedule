package main

import (
	"Schedule/server/db"

	"Schedule/server/model"
)

func main() {
	db := db.Connection()
	defer db.Close()

	db.AutoMigrate(&model.Login{})
	db.AutoMigrate(&model.Group{})
	db.AutoMigrate(&model.Task{})
	db.AutoMigrate(&model.GTask{})
}
