package sqldatabase

import (
	"database/sql"
	"fmt"
	"log"
)

type SensorQuery struct {
	db *sql.DB
}

func NewSensorQuery(db *sql.DB) *SensorQuery {
	query := &SensorQuery{db}

	return query
}

func (sensorQuery *SensorQuery) Close() {
	sensorQuery.db.Close()
}

func (sensorQuery *SensorQuery) CreateStation(station Station) Station {
	prepareStmt, err := sensorQuery.db.Prepare("INSERT INTO station(Id, label, latitude, longitude) VALUES($1, $2, $3, $4)")
	if err != nil {
		log.Fatal(err)
	}

	defer prepareStmt.Close()

	fmt.Printf("%v\n", station)

	_, err = prepareStmt.Exec(station.Id, station.Label, station.Latitude, station.Longitude)
	if err != nil {
		log.Fatal(err)
	}

	return station
}

func (sensorQuery *SensorQuery) GetAllStations() []Station {

	getAllStationsPrepare, err := sensorQuery.db.Prepare("select Id, label, latitude, longitude from station")
	if err != nil {
		log.Fatal(err)
	}

	defer getAllStationsPrepare.Close()

	rows, err := getAllStationsPrepare.Query()

	if err != nil {
		log.Fatal(err, "9")
	}

	defer rows.Close()
	var (
		id        int
		label     string
		latitude  float64
		longitude float64
	)

	stations := []Station{}

	for rows.Next() {
		err := rows.Scan(&id, &label, &latitude, &longitude)
		if err != nil {
			log.Fatal(err, "0")
		}

		s := Station{id, label, latitude, longitude}

		stations = append(stations, s)
	}

	return stations
}

func (sensorQuery *SensorQuery) GetAllEnhancedParameters() []EnhancedParameter {
	prepareStmt, err := sensorQuery.db.Prepare("select Id, parameter_id, cell_methods, time_interval, vertical_datum from enhanced_parameter")
	if err != nil {
		log.Fatal(err)
	}

	defer prepareStmt.Close()

	rows, err := prepareStmt.Query()

	if err != nil {
		log.Fatal(err, "91")
	}

	defer rows.Close()
	var (
		id            int
		parameterId   int
		cellMethods   string
		interval      string
		verticalDatum string
	)

	enhancedParameters := []EnhancedParameter{}

	for rows.Next() {
		err := rows.Scan(&id, &parameterId, &cellMethods, &interval, &verticalDatum)
		if err != nil {
			log.Fatal(err, "01")
		}

		ep := EnhancedParameter{id, parameterId, cellMethods, interval, verticalDatum}

		enhancedParameters = append(enhancedParameters, ep)
	}

	return enhancedParameters
}

func (sensorQuery *SensorQuery) CreateEnhancedParameter(enhancedParameter EnhancedParameter) EnhancedParameter {
	var prepareStmt *sql.Stmt

	prepareStmt, err := sensorQuery.db.Prepare(
		"INSERT INTO enhanced_parameter(Id, parameter_id, cell_methods, time_interval, vertical_datum) VALUES($1, $2, $3, $4, $5)")
	if err != nil {
		log.Fatal(err)
	}

	_, err = prepareStmt.Exec(enhancedParameter.Id, enhancedParameter.ParameterId,
		enhancedParameter.CellMethods, enhancedParameter.Interval, enhancedParameter.VerticalDatum)
	if err != nil {
		log.Fatal(err)
	}

	return enhancedParameter
}
