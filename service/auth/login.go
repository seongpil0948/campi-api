package auth

import "github.com/gin-gonic/gin"

func Login(c *gin.Context) {
	getPath := c.Request.URL.String()
	c.JSON(200, gin.H{
		"pathInfo": getPath,
	})
}
