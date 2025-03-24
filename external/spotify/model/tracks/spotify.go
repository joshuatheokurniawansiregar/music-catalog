package tracks

type SearchResponse struct {
	Limit  int                  `json:"limit"`
	Offset int                  `json:"offset"`
	Items  []SpotifyTrackObject `json:"items"`
	Total  int                  `json:"total"`
}

type SpotifyTrackObject struct {
	//album related fields
	AlbumType        string   `json:"albumTypes"`
	AlbumTotalTracks int      `json:"totalTracks"`
	AlbumImagesURL   []string `json:"albumImagesURL"`
	AlbumName        string   `json:"albumName"`

	//artists related field
	ArtistsName []string `json:"artistsName"`

	//tracks related fields
	Explicit bool   `json:"explicit"`
	Href     string `json:"href"`
	ID       string `json:"id"`
	Name     string `json:"name"`
}