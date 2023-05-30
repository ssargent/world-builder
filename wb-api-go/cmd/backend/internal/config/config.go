package config

import "time"

type Config struct {
	Port     int
	Database DatabaseConfig
	Cache    CacheConfig
}

type DatabaseConfig struct {
	Driver   string `split_words:"true" default:"postgres"`
	Username string `split_words:"true" default:"wb"`
	Password string `required:"true" split_words:"true" `
	Server   string `split_words:"true" default:"localhost"`
	Name     string `split_words:"true" default:"world-builder"`
}

type CacheConfig struct {
	DefaultExpiration time.Duration `split_words:"true" default:"5m"`
	DefaultCleanup    time.Duration `split_words:"true" default:"10m"`
}
