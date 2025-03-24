package tracks

import (
	"context"
	"reflect"
	"testing"

	spotifyModel "github.com/joshuatheokurniawansiregar/music-catalog/external/spotify/model/tracks"
	spotifyTracksRepository "github.com/joshuatheokurniawansiregar/music-catalog/external/spotify/repository/tracks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_service_Search(t *testing.T) {
	mockCtrl:= gomock.NewController(t)
	defer mockCtrl.Finish()
	mockSpotifyOutbound:= NewMockspotifyOutbound(mockCtrl)
	type args struct {
		query     string
		pageSize  int
		pageIndex int
	}
	var next string= "https://api.spotify.com/v1/search?offset=10&limit=10&query=bohemian%20rhapsody&type=track&market=ID&locale=en-US,en;q%3D0.9"
	tests := []struct {
		name    string
		s       *service
		args    args
		want    *spotifyModel.SearchResponse
		wantErr bool
		mockFunc func (args args)
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: args{
				query: "Bohemian Rhapsody",
				pageSize: 10,
				pageIndex: 1,
			},
			want: &spotifyModel.SearchResponse{
				Limit :10,
				Offset:0,
				Items: []spotifyModel.SpotifyTrackObject{
					{
						AlbumType: "album",
						AlbumTotalTracks:22,
						AlbumImagesURL:[]string{
								"https://i.scdn.co/image/ab67616d0000b273e8b066f70c206551210d902b",
								"https://i.scdn.co/image/ab67616d00001e02e8b066f70c206551210d902b",	
								"https://i.scdn.co/image/ab67616d00004851e8b066f70c206551210d902b",
						},
						AlbumName:"Bohemian Rhapsody (The Original Soundtrack)",

						ArtistsName: []string{
							"Queen",
						},

						Explicit: false,
						Href: "https://api.spotify.com/v1/tracks/3z8h0TU7ReDPLIbEnYhWZb",
						ID : "3z8h0TU7ReDPLIbEnYhWZb",
						Name: "Bohemian Rhapsody",
					},
					{
						AlbumType: "compilation",
						AlbumTotalTracks:17,
						AlbumImagesURL:[]string{
								"https://i.scdn.co/image/ab67616d0000b273bb19d0c22d5709c9d73c8263",
								"https://i.scdn.co/image/ab67616d00001e02bb19d0c22d5709c9d73c8263",	
								"https://i.scdn.co/image/ab67616d00004851bb19d0c22d5709c9d73c8263",
						},
						AlbumName:"Greatest Hits (Remastered)",

						ArtistsName: []string{
							"Queen",
						},

						Explicit: false,
						Href: "https://api.spotify.com/v1/tracks/2OBofMJx94NryV2SK8p8Zf",
						ID : "2OBofMJx94NryV2SK8p8Zf",
						Name: "Bohemian Rhapsody - Remastered 2011",
					},
				},
				Total: 898,
			},
			wantErr: false,
			mockFunc: func(args args){
				mockSpotifyOutbound.EXPECT().Search(gomock.Any(), args.query, 10, 0).Return(&spotifyTracksRepository.SpotifySearchResponse{
					Tracks: spotifyTracksRepository.SpotifyTracks{
						Href:"https://api.spotify.com/v1/search?offset=0&limit=10&query=bohemian%20rhapsody&type=track&market=ID&locale=en-US,en;q%3D0.9",
						Limit: 10,
						Next: &next,
						Offset: 0,
						Previous:nil,
						Total:898,
						Items: []spotifyTracksRepository.SpotifyTrackObject{
							{
								Album: spotifyTracksRepository.SpotifyAlbumObject{
									AlbumType: "album",
									TotalTracks: 22,
									Images: []spotifyTracksRepository.SpotifyAlbumImage{
										{
											URL: "https://i.scdn.co/image/ab67616d0000b273e8b066f70c206551210d902b",
										},
										{
											URL: "https://i.scdn.co/image/ab67616d00001e02e8b066f70c206551210d902b",	
										},
										{
											URL: "https://i.scdn.co/image/ab67616d00004851e8b066f70c206551210d902b",
										},
									},
									Name: "Bohemian Rhapsody (The Original Soundtrack)",
								},
								Artists: []spotifyTracksRepository.SpotifyArtistObject{
									{
										HREF: "https://api.spotify.com/v1/artists/1dfeR4HaWDbWqFHLkxsg1d",
										Name: "Queen",
									},
								},
								Explicit: false,
								Href: "https://api.spotify.com/v1/tracks/3z8h0TU7ReDPLIbEnYhWZb",
								ID: "3z8h0TU7ReDPLIbEnYhWZb",
								Name: "Bohemian Rhapsody",
							},
							{
								Album: spotifyTracksRepository.SpotifyAlbumObject{
									AlbumType: "compilation",
									TotalTracks: 17,
									Images: []spotifyTracksRepository.SpotifyAlbumImage{
										{
											URL:"https://i.scdn.co/image/ab67616d0000b273bb19d0c22d5709c9d73c8263",
										},
										{
											URL: "https://i.scdn.co/image/ab67616d00001e02bb19d0c22d5709c9d73c8263",	
										},
										{
											URL: "https://i.scdn.co/image/ab67616d00004851bb19d0c22d5709c9d73c8263",
										},
									},
									Name: "Greatest Hits (Remastered)",
								},
								Artists: []spotifyTracksRepository.SpotifyArtistObject{
									{
										HREF: "https://api.spotify.com/v1/artists/1dfeR4HaWDbWqFHLkxsg1d",
										Name: "Queen",
									},
								},
								Explicit: false,
								Href: "https://api.spotify.com/v1/tracks/2OBofMJx94NryV2SK8p8Zf",
								ID: "2OBofMJx94NryV2SK8p8Zf",
								Name: "Bohemian Rhapsody - Remastered 2011",
							},
						},
					},
				}, nil)
			},
		},
		{
			name: "failed",
			args: args{
				query: "Bohemian Rhapsody",
				pageSize: 10,
				pageIndex: 1,
			},
			want: nil,
			wantErr: true,
			mockFunc: func(args args){
				mockSpotifyOutbound.EXPECT().Search(gomock.Any(), args.query, 10, 0).Return( nil, assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc(tt.args)
			var service *service = &service{
				spotifyOutbound: mockSpotifyOutbound,
			}
			got, err := service.Search(context.Background(), tt.args.query, tt.args.pageSize, tt.args.pageIndex)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
