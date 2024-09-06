package config

type Config struct {
	Debug    bool
	FilePath string
}

var Global Config

func (c *Config) Set(debug bool, path string) {
	c.Debug = debug
	c.FilePath = path
}
