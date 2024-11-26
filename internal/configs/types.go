package configs

type (
	Config struct {
		Service  Service  `mapastructure:"service"`
		Database Database `mapstructure:"database"`
	}
	Service struct {
		Port      string `mapstructure:"port"`
		SecretJWT string `mapstructure:"secretJWT"`
	}
	Database struct {
		DataSourceName string `mapstructure:"dataSoruceName"`
	}
)