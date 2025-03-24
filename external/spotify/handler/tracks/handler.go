package tracks

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/joshuatheokurniawansiregar/music-catalog/external/spotify/model/tracks"
	"github.com/joshuatheokurniawansiregar/music-catalog/internal/middleware"
)

//go:generate mockgen -source=handler.go -destination=handler_mock_test.go -package=tracks
type service interface {
	Search(ctx context.Context, query string, pageSize, pageIndex int)(*tracks.SearchResponse, error)	
}

type Handler struct {
	*gin.Engine
	service service
}

func NewHandler(api *gin.Engine, service service)*Handler{
	return &Handler{
		api,
		service,
	}
}

func(h *Handler) RegisterRoute(){
	var route *gin.RouterGroup = h.Group("tracks")
	route.Use(middleware.AuthMiddleware())
	route.GET("/search", h.Search)
}