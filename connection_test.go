package sqldatabase

import (
	"testing"
	"log"
)


func TestConnection(t *testing.T){

	sensorQuery := BuildSensorQuery()

	defer sensorQuery.Close()

	stations := sensorQuery.GetAllStations()

	log.Println("length" , len(stations))

	for _, s := range stations {
		log.Println(s)
	}

	eps := sensorQuery.GetAllEnhancedParameters()

	for _, ep := range eps {
		log.Println(ep)
	}
}
