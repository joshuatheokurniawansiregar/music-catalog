package memberships

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joshuatheokurniawansiregar/music-catalog/internal/model/memberships"
)

func (h *Handler) Login(c *gin.Context){
	var request memberships.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil{
		c.JSON(http.StatusBadRequest, &gin.H{
			"err":err.Error(),
		})
		return
	}

	accessToken, err := h.service.Login(request)
	if err != nil {
		fmt.Println("test case terexecute")
		c.JSON(http.StatusInternalServerError, &gin.H{
			"error": err.Error(),
		})
		return
	}
	
	c.JSON(http.StatusOK, &memberships.LoginResponse{
		AccessToken: accessToken,
	})
}