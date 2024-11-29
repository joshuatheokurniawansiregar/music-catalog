package memberships

import (
	"github.com/gin-gonic/gin"
	"github.com/joshuatheokurniawansiregar/music-catalog/internal/model/memberships"
)

//go:generate mockgen -source=handler.go -destination=handler_mock_test.go -package=memberships
type service interface {
	SignUp(request memberships.SignupRequest) error
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
	var route *gin.RouterGroup = h.Group("memberships")
	route.POST("/sign_up", h.Signup)
}