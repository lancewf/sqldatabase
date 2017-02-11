package sqldatabase

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"
	"fmt"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "postgres"
	DB_NAME     = "sensor"
)


func BuildSensorQuery() *SensorQuery{
	dbInfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		log.Fatal(err, "1")
	}

	return NewSensorQuery(db)
}
