package config

import "time"

type Config struct {
	Port     int
	Database DatabaseConfig
	Cache    CacheConfig
}

type DatabaseConfig struct {
	Driver   string `default:"postgres" split_words:"true" json:"driver,omitempty"`
	Username string `default:"wb" split_words:"true"  json:"username,omitempty"`
	Password string `required:"true" split_words:"true" json:"password,omitempty"`
	Server   string `default:"localhost" split_words:"true" json:"server,omitempty"`
	Name     string `default:"world-builder" split_words:"true" json:"name,omitempty"`
}

type CacheConfig struct {
	DefaultExpiration time.Duration `split_words:"true" default:"5m" json:"default_expiration,omitempty"`
	DefaultCleanup    time.Duration `split_words:"true" default:"10m" json:"default_cleanup,omitempty"`
}
