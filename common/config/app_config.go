package config

// AppConfig defines the structure of shared configuration.
type AppConfig struct {
	Database struct {
		DSN string `json:"dsn"`
	} `json:"database"`
	Server struct {
		Port string `json:"port"`
	} `json:"server"`
}
