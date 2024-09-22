package config

type Config struct {
	DB string `env:"APP_DB_DSN" envDefault:"postgres://test-user:test-password@localhost:5432/test"`
}
