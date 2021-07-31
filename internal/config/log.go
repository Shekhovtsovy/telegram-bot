package config

type Log struct {
	Facility      string `envconfig:"LOG_DEFAULT_FACILITY" default:""`
	GraylogEnable bool   `envconfig:"GRAYLOG_ENABLE" default:"false"`
	GraylogUdpUri string `envconfig:"GRAYLOG_UDP_URI" default:"0.0.0.0:12201"`
}
