package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	spotifyTrackHandler "github.com/joshuatheokurniawansiregar/music-catalog/external/spotify/handler/tracks"
	spotifyTracksOutbound "github.com/joshuatheokurniawansiregar/music-catalog/external/spotify/repository/tracks"
	spotifyService "github.com/joshuatheokurniawansiregar/music-catalog/external/spotify/service/tracks"
	"github.com/joshuatheokurniawansiregar/music-catalog/internal/configs"
	membershipsHandlerPack "github.com/joshuatheokurniawansiregar/music-catalog/internal/handler/memberships"
	"github.com/joshuatheokurniawansiregar/music-catalog/internal/model/memberships"
	membershipsRepoPack "github.com/joshuatheokurniawansiregar/music-catalog/internal/repository/memberships"
	membershipsServicePack "github.com/joshuatheokurniawansiregar/music-catalog/internal/service/memberships"
	"github.com/joshuatheokurniawansiregar/music-catalog/pkg/httpclient"
	"github.com/joshuatheokurniawansiregar/music-catalog/pkg/internalsql"
)

func main() {
	router := gin.Default()

	var(
		cfg *configs.Config
		err error
	)

	err = configs.Init(
		configs.WithConfigFolder([]string{"./internal/configs/"}),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)


	cfg = configs.GetConfig()

	if err != nil {
		log.Fatal("Gagal Inisiasi Config",err.Error())
	}


	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil{
		log.Fatal(err)
	}

	db.AutoMigrate(&memberships.User{})
	
	var membershipsRepo = membershipsRepoPack.NewRepository(db)
	var membershipsService = membershipsServicePack.NewService(membershipsRepo, cfg)
	var membershipsHandler = membershipsHandlerPack.NewHandler(router, membershipsService)
	membershipsHandler.RegisterRoute()

	var httpClient = httpclient.NewClient(&http.Client{})
	var spotifyOutbound = spotifyTracksOutbound.NewSpotifyOutbound(cfg, httpClient)
	var spotifySvc = spotifyService.NewService(spotifyOutbound)
	var spotifyTraHandler = spotifyTrackHandler.NewHandler(router, spotifySvc)
	spotifyTraHandler.RegisterRoute()
	
	router.Run(cfg.Service.Port)
}