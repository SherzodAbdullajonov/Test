package database

import "fmt"

func Load() (Config, error) {
	var cfg config

	return Config{
		host:             cfg.Host,
		port:             cfg.Port,
		postgresHost:     cfg.PostgresHost,
		postgresPort:     cfg.PostgresPort,
		postgresUser:     cfg.PostgresUser,
		postgresPassword: cfg.PostgresPassword,
		postgresDB:       cfg.PostgresDB,
	}, nil
}

type Config struct {
	host             string
	port             int
	postgresHost     string
	postgresPort     int
	postgresUser     string
	postgresPassword string
	postgresDB       string
}

func (c Config) ListenAddress() string {
	return fmt.Sprintf("%s:%d", c.host, c.port)
}

func (c Config) PostgresConnString() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.postgresHost, c.postgresPort, c.postgresUser, c.postgresPassword, c.postgresDB,
	)
}

type config struct {
	Environment      string `envconfig:"ENVIRONMENT" default:"development"`
	Host             string `envconfig:"HOST" default:"0.0.0.0"`
	Port             int    `envconfig:"PORT" default:"8000"`
	PostgresHost     string `envconfig:"POSTGRES_HOST" required:"true"`
	PostgresPort     int    `envconfig:"POSTGRES_PORT" required:"true"`
	PostgresUser     string `envconfig:"POSTGRES_USER" required:"true"`
	PostgresPassword string `envconfig:"POSTGRES_PASSWORD" required:"true"`
	PostgresDB       string `envconfig:"POSTGRES_DB" required:"true"`
}
