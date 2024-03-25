package handler

import (
	"github.com/gin-gonic/gin"
	"task/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.POST("/check", h.check)
	router.POST("/set-param", h.setParam)

	return router
}
