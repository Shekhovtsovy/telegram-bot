package config

type Db struct {
	Host     string `envconfig:"DB_HOST" default:""`
	Port     string `envconfig:"DB_PORT" default:""`
	Name     string `envconfig:"DB_NAME" default:""`
	Username string `envconfig:"DB_USERNAME" default:""`
	Password string `envconfig:"DB_PASSWORD" default:""`
}
