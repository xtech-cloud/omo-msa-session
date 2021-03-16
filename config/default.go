package config

const defaultJson string = `{
	"service": {
		"address": ":7089",
		"ttl": 15,
		"interval": 10
	},
	"logger": {
		"level": "info",
		"file": "logs/server.log",
		"std": false
	},
	"basic": {
		"timeout": 43200,
		"secret": "yumei2020-platfom"
	}
}
`
