package configs

type (
	Config struct {
		Service       Service       `mapstructure:"service"`
		Database      Database      `mapstructure:"database"`
		SpotifyConfig SpotifyConfig `mapstructure:"spotifyConfig"`
	}

	Service struct {
		Port      string `mapstructure:"port"`
		SecretKey string `mapstructure:"secretKey"`
	}

	Database struct {
		DataSourceName string `mapstructure:"dataSourceName"`
	}

	SpotifyConfig struct {
		ClientID     string `mapstructure:"clientID"`
		ClientSecret string `mapstructure:"clientSecret"`
	}
)