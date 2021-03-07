package config

type Configuration struct {
	Server   ServerConfiguration   `yaml:"server"`
	Rabbitmq RabbitMQConfiguration `yaml:"rabbitmq"`
	Fluentd  FluentdConfiguration  `yaml:"fluentd"`
	Consumer ConsumerConfiguration `yaml:"consumer"`
}

func NewDefaultConfig() []byte {
	defaultConf := []byte(`
		server:
			host: "127.0.0.1"
			port: 8080
			runMode: "debug"
		rabbitmq:
			host: "127.0.0.1"
			port: 5672
			account: "guest"
			password: "guest"
		cassandra:
			host: "127.0.0.1"
			port: 9042
		fluentd:
			host: "127.0.0.1"
			port: 24224
		consumer:
			type: "ptt"
		`)
	return defaultConf
}
