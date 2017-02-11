package sqldatabase

import (
	"database/sql"
	"log"
)

type SensorQuery struct {
	db *sql.DB
	preparedStatementsMap map[string]*sql.Stmt
}

func NewSensorQuery(db *sql.DB) *SensorQuery{
	query := &SensorQuery{db, make(map[string]*sql.Stmt)}

	return query
}

func (sensorQuery *SensorQuery) Close() {
	for _,stmt := range sensorQuery.preparedStatementsMap {
		stmt.Close()
	}
	sensorQuery.db.Close()
}

func (sensorQuery *SensorQuery) CreateStation(station Station) Station {
	var prepareStmt *sql.Stmt

	if sensorQuery.preparedStatementsMap["CreateStation"] == nil {
		stmt, err := sensorQuery.db.Prepare("INSERT INTO station(label, latitude, longitude) VALUES($1, $2, $3) RETURNING id")
		if err != nil {
			log.Fatal(err)
		}

		sensorQuery.preparedStatementsMap["CreateStation"] = stmt

		prepareStmt = stmt
	}

	var lastInsertId int

	err := prepareStmt.QueryRow(station.label, station.latitude, station.longitude).Scan(&lastInsertId)
	if err != nil {
		log.Fatal(err)
	}

	station.id = lastInsertId

	return station
}

func (sensorQuery *SensorQuery) GetAllStations() []Station {

	var getAllStationsPrepare *sql.Stmt

	if sensorQuery.preparedStatementsMap["GetAllStations"] == nil {
		stmt, err := sensorQuery.db.Prepare("select id, label, latitude, longitude from station")
		if err != nil {
			log.Fatal(err)
		}

		sensorQuery.preparedStatementsMap["GetAllStations"] = stmt

		getAllStationsPrepare = stmt
	}

	rows, err := getAllStationsPrepare.Query()

	if err != nil {
		log.Fatal(err, "9")
	}

	defer rows.Close()
	var (
		id int
		label string
		latitude float64
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

	var prepareStmt *sql.Stmt

	if sensorQuery.preparedStatementsMap["GetAllEnhancedParameters"] == nil {
		stmt, err := sensorQuery.db.Prepare("select id, parameter_id, cell_methods, time_interval, vertical_datum from enhanced_parameter")
		if err != nil {
			log.Fatal(err)
		}

		sensorQuery.preparedStatementsMap["GetAllEnhancedParameters"] = stmt

		prepareStmt = stmt
	}

	rows, err := prepareStmt.Query()

	if err != nil {
		log.Fatal(err, "91")
	}

	defer rows.Close()
	var (
		id int
		parameterId int
		cellMethods string
		interval string
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

	if sensorQuery.preparedStatementsMap["CreateEnhancedParameter"] == nil {
		stmt, err := sensorQuery.db.Prepare(
			"INSERT INTO enhanced_parameter(id, parameter_id, cell_methods, time_interval, vertical_datum) VALUES($1, $2, $3, $4, $5)")
		if err != nil {
			log.Fatal(err)
		}

		sensorQuery.preparedStatementsMap["CreateEnhancedParameter"] = stmt

		prepareStmt = stmt
	}

	_, err := prepareStmt.Exec(enhancedParameter.id, enhancedParameter.parameterId,
		enhancedParameter.cellMethods, enhancedParameter.interval, enhancedParameter.verticalDatum)
	if err != nil {
		log.Fatal(err)
	}

	return enhancedParameter
}