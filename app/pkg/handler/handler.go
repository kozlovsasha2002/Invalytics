package handler

import (
	"Invalytics/app/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/signup", h.signUp)
		auth.POST("/login", h.login)
	}

	api := router.Group("/api", h.userIdentity)
	{
		deps := api.Group("/deps")
		deps.POST("/", h.createDeposit)
		deps.GET("/", h.getAllDeposits)
		deps.GET("/:id", h.getDepositById)
		deps.PATCH("/:id", h.updateDeposit)
		deps.DELETE("/:id", h.deleteDeposit)
	}

	return router
}
