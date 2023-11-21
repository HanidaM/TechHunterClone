package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "baurzhkk"
	password = "bauka2003"
	dbname   = "techbd"
)

//const (
//	host     = "ep-misty-cloud-327990.us-east-2.aws.neon.tech"
//	port     = 5432
//	user     = "baurzhkk"
//	password = "rXGDm0y6ViTF"
//	dbname   = "techbd"
//)

var DB *gorm.DB

func ConnectDB() {

	var err error

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require",
		host, port, user, password, dbname)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	log.Println("Connected to the database successfully")
}
