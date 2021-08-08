package config

type Config struct {
	Directory string // Location of template files
	Snip      string // Name of template files
}

func New() Config {
	cfg := Config{}
	return cfg
}
