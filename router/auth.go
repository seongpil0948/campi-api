package router

import (
	"campi/api/service/auth"

	"github.com/gin-gonic/gin"
)

func SetAuthRoutes(g *gin.RouterGroup) {
	// PingExample godoc
	// @Summary ping example
	// @Schemes
	// @Description do ping
	// @Tags example
	// @Accept json
	// @Produce json
	// @Success 200 {string} Helloworld
	// @Router /example/helloworld [get]
	g.POST("/login", auth.Login)
}
