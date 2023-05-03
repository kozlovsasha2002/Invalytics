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

		bonds := api.Group("/bonds")
		bonds.POST("/", h.CreateBond)
		bonds.GET("/", h.GetAllBonds)
		bonds.GET("/:id", h.GetBondById)
		bonds.PATCH("/:id", h.UpdateBond)
		bonds.DELETE("/:id", h.DeleteBond)

		shares := api.Group("/shares")
		shares.POST("/", h.CreateShare)
		shares.GET("/", h.GetAllShares)
		shares.GET("/:id", h.GetShareById)
		shares.PATCH("/:id", h.UpdateShare)
		shares.DELETE("/:id", h.DeleteShare)
	}

	return router
}
