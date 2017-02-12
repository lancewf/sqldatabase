package sqldatabase

import "strconv"

/*
CREATE TABLE public.station
(
  Id integer NOT NULL DEFAULT nextval('station_id_seq'::regclass),
  Label character varying,
  Latitude double precision NOT NULL,
  Longitude double precision NOT NULL,
  CONSTRAINT station_pkey PRIMARY KEY (Id)
)
 */
type Station struct {
	Id        int
	Label     string
	Latitude  float64
	Longitude float64
}

func NewStation(id int, label string, latitude float64, longitude float64) Station {
	return Station{id, label, latitude, longitude}
}

func (s Station) String() string {
	return "Station( Id: " + strconv.Itoa(s.Id) + ", Label: " + s.Label + ", lat: " +
		strconv.FormatFloat(s.Latitude, 'E', -1, 64) + ", lon: " +
		strconv.FormatFloat(s.Longitude, 'E', -1, 64) + ")"
}

/*
CREATE TABLE public.enhanced_parameter
(
  Id integer NOT NULL DEFAULT nextval('enhanced_parameter_id_seq'::regclass),
  parameter_id integer NOT NULL,
  cell_methods character varying NOT NULL DEFAULT ''::character varying,
  time_interval character varying NOT NULL DEFAULT ''::character varying,
  vertical_datum character varying NOT NULL DEFAULT ''::character varying,
  CONSTRAINT enhanced_parameter_pkey PRIMARY KEY (Id)
)
 */
type EnhancedParameter struct {
	Id            int
	ParameterId   int
	CellMethods   string
	Interval      string
	VerticalDatum string
}

func NewEnhancedParameter(id int, parameterId int, cellMethods string, interval string, verticalDatum string) EnhancedParameter {
	return EnhancedParameter{id, parameterId, cellMethods, interval, verticalDatum}
}

func (ep EnhancedParameter) String() string {
	return "EnhancedParameter( Id: " + strconv.Itoa(ep.Id) + ", ParameterId: " + strconv.Itoa(ep.ParameterId) +
		", cellmethods: " + ep.CellMethods + ", Interval: " + ep.Interval + ", verticaldatum: " + ep.VerticalDatum + ")"
}

