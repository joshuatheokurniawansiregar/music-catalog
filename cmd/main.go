package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/joshuatheokurniawansiregar/music-catalog/internal/configs"
	"github.com/joshuatheokurniawansiregar/music-catalog/pkg/internalsql"
	"github.com/rs/zerolog/log"
)

func main() {
	var(
		cfg *configs.Config
		err error
	)

	err = configs.Init(
		configs.WithConfigFolder([]string{"./internal/configs/", "./configs/"}),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)

	if err != nil{
		log.Error().Err(err)
	}

	cfg = configs.GetConfig()

	_, err = internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil{
		log.Error().Err(err)
	}
	
	router := gin.Default()
	router.Run(cfg.Service.Port)
}