package database

import (
	"fmt"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DBCon *gorm.DB
	Env   map[string]string
)

func InitDB() {
	dbURL := "postgres://" + Env["DATABASE_USERNAME"] + ":" + Env["DATABASE_PASSWORD"] + "@" + Env["DATABASE_HOST"] + ":" + Env["DATABASE_PORT"] + "/"
	fmt.Println(dbURL)
	conn, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{}) //this line can be removed
	conn, err = gorm.Open(postgres.Open(dbURL+Env["DATABASE_NAME"]), &gorm.Config{})
	DBCon = conn
	if err != nil {
		panic(err)
	}

}
func IsTableExist(table_exist bool) bool {
	DBCon.Raw("SELECT EXISTS ( SELECT 1 FROM pg_tables WHERE tablename = 'Basic' )").Scan(&table_exist)
	return table_exist
}
