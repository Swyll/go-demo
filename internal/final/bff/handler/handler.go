package handler

import (
	"go-demo/internal/final/bff/server"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	bs *server.BffServer
}

func NewHandler(bs *server.BffServer) *Handler {
	return &Handler{
		bs: bs,
	}
}

func (h *Handler) QueryAcountInfo(c *gin.Context) {
	acountID := c.Query("acount_id")
	if acountID == "" {
		c.JSON(400, gin.H{
			"data": "参数错误",
		})
	}

	info, err := h.bs.QueryAcountInfo(acountID)
	if err != nil {
		c.JSON(500, gin.H{
			"data": "查询出错",
		})
	}

	c.JSON(200, gin.H{
		"data": info,
	})
}

type Route interface {
	QueryAcountInfo(c *gin.Context)
}

func Init(e *gin.Engine, route Route) {
	e.GET("/swy", route.QueryAcountInfo)
}
