package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/KonstantinP85/s-go-service/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}


func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()


	post := router.Group("/post")
	{
		post.POST("/event", h.postEvent)
	}

	return router
}