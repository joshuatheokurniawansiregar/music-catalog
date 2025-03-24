package tracks

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/joshuatheokurniawansiregar/music-catalog/internal/configs"
	"github.com/joshuatheokurniawansiregar/music-catalog/pkg/httpclient"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_outbound_Search(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	
	mockHTTPClient := httpclient.NewMockHTTPClient(mockCtrl)

	type args struct {
		query  string
		limit  int
		offset int
	}
	var next string= "https://api.spotify.com/v1/search?offset=10&limit=10&query=bohemian%20rhapsody&type=track&market=ID&locale=en-US,en;q%3D0.9"
	tests := []struct {
		name    string
		args    args
		want    *SpotifySearchResponse
		wantErr bool
		mockFn func(args args)
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: args{
				query: "bohemian rhapsody",
				limit: 10,
				offset: 0,
			},
			want: &SpotifySearchResponse{
				Tracks: SpotifyTracks{
					Href:"https://api.spotify.com/v1/search?offset=0&limit=10&query=bohemian%20rhapsody&type=track&market=ID&locale=en-US,en;q%3D0.9",
					Limit: 10,
					Next: &next,
					Offset: 0,
					Previous:nil,
					Total:898,
					Items: []SpotifyTrackObject{
						{
							Album: SpotifyAlbumObject{
								AlbumType: "album",
								TotalTracks: 22,
								Images: []SpotifyAlbumImage{
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
							Artists: []SpotifyArtistObject{
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
							Album: SpotifyAlbumObject{
								AlbumType: "compilation",
								TotalTracks: 17,
								Images: []SpotifyAlbumImage{
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
							Artists: []SpotifyArtistObject{
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
			},
			wantErr:false,
			mockFn: func(args args){
				var params url.Values = url.Values{}
				params.Set("q", args.query)
				params.Set("type", "track")
				params.Set("limit", strconv.Itoa(args.limit))
				params.Set("offset", strconv.Itoa(args.offset))
				var basePath string = `https://api.spotify.com/v1/search`
				var urlPath string = fmt.Sprintf("%s?%s", basePath, params.Encode())
				var(
					request *http.Request
					err error
				)
				request, err = http.NewRequest(http.MethodGet, urlPath, nil)
				assert.NoError(t, err)
				request.Header.Set("Authorization", "Bearer accessToken")
				mockHTTPClient.EXPECT().Do(request).Return(&http.Response{
					StatusCode: 200,
					Body: io.NopCloser(bytes.NewBufferString(searchResponse)),
				},nil)
			},
		},
		{
			name: "fail",
			args: args{
				query: "bohemian rhapsody",
				limit: 10,
				offset: 0,
			},
			want: nil,
			wantErr:true,
			mockFn: func(args args){
				var params url.Values = url.Values{}
				params.Set("q", args.query)
				params.Set("type", "track")
				params.Set("limit", strconv.Itoa(args.limit))
				params.Set("offset", strconv.Itoa(args.offset))
				var basePath string = `https://api.spotify.com/v1/search`
				var urlPath string = fmt.Sprintf("%s?%s", basePath, params.Encode())
				var(
					request *http.Request
					err error
				)
				request, err = http.NewRequest(http.MethodGet, urlPath, nil)
				assert.NoError(t, err)
				request.Header.Set("Authorization", "Bearer accessToken")
				mockHTTPClient.EXPECT().Do(request).Return(&http.Response{
					StatusCode: 200,
					Body: io.NopCloser(bytes.NewBufferString("Internal Server Error")),
				},nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			var o *outbound = &outbound{
				cfg: &configs.Config{},
				client: mockHTTPClient,
				AccessToken: "accessToken",
				TokenType: "Bearer",
				ExpiredAt: time.Now().Add(1 * time.Hour),
			}
			got, err := o.Search(context.Background(), tt.args.query, tt.args.limit, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("outbound.Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("outbound.Search() = %v", got)
				t.Errorf("\nwant: %v", tt.want)
			}
		})
	}
}
