package config

type Config struct {
	Port     int
	Database DatabaseConfig
}

type DatabaseConfig struct {
	Driver   string `required:"true" split_words:"true"`
	Username string `required:"true" split_words:"true"`
	Password string `required:"true" split_words:"true"`
	Server   string `required:"true" split_words:"true"`
	Name     string `required:"true" split_words:"true"`
}
