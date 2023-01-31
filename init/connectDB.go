package init

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(conf *Config) {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", conf.DBHost, conf.DBUsername, conf.DBPassword, conf.DBName, conf.DBPort)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Connection to database failed . . .")
	}
	fmt.Println("Connection to database successful . . .")
}
