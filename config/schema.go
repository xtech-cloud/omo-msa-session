package config

type ServiceConfig struct {
	TTL      int64  `yaml:"ttl"`
	Interval int64  `yaml:"interval"`
	Address  string `yaml:"address"`
}

type LoggerConfig struct {
	Level string `yaml:"level"`
	Dir string `yaml:"dir"`
}

type BasicConfig struct {
	Timeout int64 `yaml:"timeout"`
	Secret string `yaml:"secret"`
}

type SchemaConfig struct {
	Service  ServiceConfig `yaml:"service"`
	Logger   LoggerConfig  `yaml:"logger"`
	Basic BasicConfig `yml:"basic"`
}
