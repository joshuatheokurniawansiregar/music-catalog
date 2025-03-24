package tracks

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	spotifyModel "github.com/joshuatheokurniawansiregar/music-catalog/external/spotify/model/tracks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestHandler_Search(t *testing.T) {
	mockCtrl:= gomock.NewController(t)
	defer mockCtrl.Finish()
	mockSvc:= NewMockservice(mockCtrl)
	tests := []struct {
		name string
		expectedStatusCode int64
		expectedBody spotifyModel.SearchResponse
		wantErr bool
		mockFn func()
	}{
		// TODO: Add test cases.
		{
			name: "success",
			expectedStatusCode: 200,
			expectedBody: spotifyModel.SearchResponse{
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
			mockFn: func() {
				mockSvc.EXPECT().Search(gomock.Any(), "bohemian rhapsody", 10, 1).Return(&spotifyModel.SearchResponse{
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
				}, nil)
			},
		},
	}
	for _, tt := range tests {
		tt.mockFn()
			var handler *Handler = &Handler{
				Engine: gin.New(),
				service: mockSvc,
			}
			handler.RegisterRoute()
			w:= httptest.NewRecorder()
			endpoint:= `/tracks/search?query=bohemian+rhapsody&pageSize=10&pageIndex=1`
			
			req, err := http.NewRequest(http.MethodGet, endpoint, nil)
			assert.NoError(t, err)
			handler.ServeHTTP(w, req)
			assert.Equal(t, tt.expectedStatusCode, int64(w.Code))

			if(!tt.wantErr){
				res:= w.Result()
				defer res.Body.Close()
				
				response := spotifyModel.SearchResponse{}
				err = json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)

				assert.Equal(t, tt.expectedBody, response)
			}
	}
}
