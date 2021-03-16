package config

type ServiceConfig struct {
	TTL      int64  `json:"ttl"`
	Interval int64  `json:"interval"`
	Address  string `json:"address"`
}

type LoggerConfig struct {
	Level string `json:"level"`
	File string `json:"file"`
	Std bool `json:"std"`
}

type BasicConfig struct {
	Timeout int64 `json:"timeout"`
	Secret string `json:"secret"`
}

type SchemaConfig struct {
	Service  ServiceConfig `json:"service"`
	Logger   LoggerConfig  `json:"logger"`
	Basic BasicConfig `json:"basic"`
}
