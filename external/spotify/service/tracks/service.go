package tracks

import (
	"context"

	"github.com/joshuatheokurniawansiregar/music-catalog/external/spotify/repository/tracks"
)

//go:generate mockgen -source=service.go -destination=service_mock_test.go -package=tracks
type spotifyOutbound interface{
	Search(ctx context.Context, query string, limit, offset int)(*tracks.SpotifySearchResponse, error)
}

type service struct{
	spotifyOutbound spotifyOutbound
}

func NewService(spotifyOutboud spotifyOutbound) *service{
	return &service{
		spotifyOutbound: spotifyOutboud,
	}
}