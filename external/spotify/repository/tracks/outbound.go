package tracks

import (
	"time"

	"github.com/joshuatheokurniawansiregar/music-catalog/internal/configs"
	"github.com/joshuatheokurniawansiregar/music-catalog/pkg/httpclient"
)

type outbound struct {
	cfg *configs.Config
	client httpclient.HTTPClient
	AccessToken string
	TokenType string
	ExpiredAt time.Time
}

func NewSpotifyOutbound(cfg *configs.Config, clnt httpclient.HTTPClient) *outbound{
	return &outbound{
		cfg: cfg,
		client: clnt,
	}
}