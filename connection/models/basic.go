package models

import (
	"gomodulename/connection/database"
	"time"
)

type Basic struct {
	Id          int64     `gorm:"AUTO_INCREMENT;primary_key;" json:"id"`
	Name        string    `json:"name" validate:"required"`
	Age         int64     `json:"age" validate:"required"`
	Email       string    `json:"email" validate:"required,email"`
	Password    string    `json:"password" validate:"required"`
	RollNumber string    `json:"phone_number" validate:"required"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
func AutoMigrateTableIfNotExist() {

	var table_exist bool

	table_exist = database.IsTableExist(table_exist)

	if table_exist == false {

		database.DBCon.AutoMigrate(&Basic{})
	}

}