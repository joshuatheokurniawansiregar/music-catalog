package tracks

import (
	"context"

	spotifyModel "github.com/joshuatheokurniawansiregar/music-catalog/external/spotify/model/tracks"
	"github.com/joshuatheokurniawansiregar/music-catalog/external/spotify/repository/tracks"
	"github.com/rs/zerolog/log"
)

func (s *service) Search(ctx context.Context, query string, pageSize, pageIndex int)(*spotifyModel.SearchResponse, error){
	limit:= pageSize
	offset := (pageIndex - 1 ) * pageSize

	//var trackDetails *tracks.SpotifySearchResponse from repository, err error
	trackDetails, err := s.spotifyOutbound.Search(ctx, query, limit, offset)
	if err != nil{
		log.Error().Err(err).Msg("error sesarch track spotify")
		return nil, err
	}

	return  modelToSearchResponse(trackDetails), nil
}

func modelToSearchResponse(spotifySearchResponse *tracks.SpotifySearchResponse)*spotifyModel.SearchResponse{
	if spotifySearchResponse == nil{
		return nil
	}
	var items []spotifyModel.SpotifyTrackObject = make([]spotifyModel.SpotifyTrackObject, 0)

	for _, item := range spotifySearchResponse.Tracks.Items{
		var imagesUrl []string = make([]string, len(item.Album.Images))
		for index, image := range item.Album.Images{
			imagesUrl[index] = image.URL
		}
		
		var artistsName []string = make([]string, len(item.Artists))
		for idx, artist := range item.Artists{
			artistsName[idx] = artist.Name
		}
		items = append(items, spotifyModel.SpotifyTrackObject{
			AlbumType	: item.Album.AlbumType,
			AlbumTotalTracks: item.Album.TotalTracks,
			AlbumImagesURL: imagesUrl,
			AlbumName: item.Album.Name,

			ArtistsName: artistsName,

			Explicit: item.Explicit,
			Href: item.Href,
			ID: item.ID,
			Name: item.Name,
		})
	}
	return &spotifyModel.SearchResponse{
		Limit: spotifySearchResponse.Tracks.Limit,
		Offset: spotifySearchResponse.Tracks.Offset,
		Items: items,
		Total: spotifySearchResponse.Tracks.Total,
	}
}
