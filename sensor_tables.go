package sqldatabase

import "strconv"

/*
CREATE TABLE public.station
(
  id integer NOT NULL DEFAULT nextval('station_id_seq'::regclass),
  label character varying,
  latitude double precision NOT NULL,
  longitude double precision NOT NULL,
  CONSTRAINT station_pkey PRIMARY KEY (id)
)
 */
type Station struct {
	id      int
	label string
	latitude float64
	longitude float64
}

func (s Station) String() string {
	return "Station( id: " + strconv.Itoa(s.id) + ", label: " + s.label + ", lat: " +
		strconv.FormatFloat(s.latitude, 'E', -1, 64) + ", lon: " +
		strconv.FormatFloat(s.longitude, 'E', -1, 64) + ")"
}

/*
CREATE TABLE public.enhanced_parameter
(
  id integer NOT NULL DEFAULT nextval('enhanced_parameter_id_seq'::regclass),
  parameter_id integer NOT NULL,
  cell_methods character varying NOT NULL DEFAULT ''::character varying,
  time_interval character varying NOT NULL DEFAULT ''::character varying,
  vertical_datum character varying NOT NULL DEFAULT ''::character varying,
  CONSTRAINT enhanced_parameter_pkey PRIMARY KEY (id)
)
 */
type EnhancedParameter struct {
	id            int
	parameterId   int
	cellMethods   string
	interval      string
	verticalDatum string
}

func (ep EnhancedParameter) String() string {
	return "EnhancedParameter( id: " + strconv.Itoa(ep.id) + ", parameterId: " + strconv.Itoa(ep.parameterId) +
		", cellmethods: " + ep.cellMethods + ", interval: " + ep.interval + ", verticaldatum: " + ep.verticalDatum + ")"
}

