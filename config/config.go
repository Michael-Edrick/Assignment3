package config

type DataPass struct {
	Username       string `env:"username"`
	Password       string `env:"password"`
	GoogleClientID string `env:"google_client_id"`
}