package config

type Log struct {
	Facility string `envconfig:"LOG_DEFAULT_FACILITY" default:""`
}
