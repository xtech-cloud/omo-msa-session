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

type DBConfig struct {
	Type     string	`yaml:"type"`
	User     string	`yaml:"user"`
	Password string	`yaml:"password"`
	IP      string	`yaml:"ip"`
	Port     string	`yaml:"port"`
	Name     string	`yaml:"name"`
}

type BasicConfig struct {
	Timeout int64 `yaml:"timeout"`
	Secret string `yaml:"secret"`
}

type SchemaConfig struct {
	Service  ServiceConfig `yaml:"service"`
	Logger   LoggerConfig  `yaml:"logger"`
	Database DBConfig      `yaml:"database"`
	Basic BasicConfig `yml:"basic"`
}
