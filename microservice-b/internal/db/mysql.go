package db

import (
	"fmt"
	"log"

	"microservice-b/internal/domain"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMySQL() *gorm.DB {
	user := "root"
	pass := "Root@123"
	host := "localhost"
	port := "3306"
	name := "interview_task_db"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		user, pass, host, port, name,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}

	db.AutoMigrate(&domain.SensorData{})
	return db
}
