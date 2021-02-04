package models

// Config ...
type Config struct {
	User     string   `yaml:"user"`
	Database string   `yaml:"database"`
	Port     string   `yaml:"port"`
	Server   string   `yaml:"server"`
	Settings []string `yaml:"settings"`
}
