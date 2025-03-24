package tracks

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Search(c *gin.Context) {
	ctx := c.Request.Context()
	query := c.Request.URL.Query().Get("query")
	pageSizeStr:= c.Query("pageSize")
	pageIndexStr := c.Query("pageIndex")

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil{
		pageSize = 10
	}

	pageIndex,err := strconv.Atoi(pageIndexStr)
	if err != nil{
		pageIndex = 1
	}

	searchResponse, err := h.service.Search(ctx, query, pageSize, pageIndex)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"search_response": searchResponse,
	})
}