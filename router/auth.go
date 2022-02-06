package router

import (
	"github.com/gin-gonic/gin"
)

func SetAuthRoutes(g *gin.RouterGroup) {

	g.POST("/login", Login)
}

// @Summary Login example
// @Schemes
// @Description Login
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Login
// @Router /auth/Login [post]
func Login(c *gin.Context) {
	getPath := c.Request.URL.String()
	c.JSON(200, gin.H{
		"pathInfo": getPath,
	})
}
