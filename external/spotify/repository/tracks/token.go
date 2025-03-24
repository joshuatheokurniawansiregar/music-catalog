package tracks

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
)

type SpotifyTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

func (o *outbound) GetTokenDetails() (string, string, error) {
	if o.AccessToken == "" || time.Now().After(o.ExpiredAt){
		err:= o.generateToken()
		if err != nil{
			return "","", err
		}
	}
	return o.AccessToken, o.TokenType, nil
}

func (o *outbound) generateToken()error{
	formData:= url.Values{}
	formData.Set("grant_type", "client_credentials")
	formData.Set("client_id", o.cfg.SpotifyConfig.ClientID)
	formData.Set("client_secret", o.cfg.SpotifyConfig.ClientSecret)

	encoded := formData.Encode()
	req, err := http.NewRequest(http.MethodPost, `https://accounts.spotify.com/api/token`, strings.NewReader(encoded))
	if err != nil{
		log.Error().Err(err).Msg("error create request to https://accounts.spotify.com/api/token")
		return err
	}

	defer req.Body.Close()

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	
	resp, err := o.client.Do(req)
	if err != nil{
		log.Error().Err(err).Msg("error execute request to https://accounts.spotify.com/api/token")
		return err
	}

	defer resp.Body.Close()


	var spotifyTokenResponse SpotifyTokenResponse
	err = json.NewDecoder(resp.Body).Decode(&spotifyTokenResponse)
	if err != nil{
		log.Error().Err(err).Msg("error unmarshal response from https://accounts.spotify.com/api/token")
		return err
	}

	o.AccessToken = spotifyTokenResponse.AccessToken
	o.TokenType = spotifyTokenResponse.TokenType
	o.ExpiredAt = time.Now().Add(time.Duration(spotifyTokenResponse.ExpiresIn)*time.Second)
	return nil
}