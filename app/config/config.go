package config

type Config struct {
	LogFile string
}
// NewConfig  init of Config
func NewConfig(logFile string) *Config {
	conf := Config{LogFile: logFile}
	return &conf
}