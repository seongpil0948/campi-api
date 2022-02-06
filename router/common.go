package router

import (
	"fmt"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetCommonRoutes(g *gin.RouterGroup) {

	g.GET("/CheckClientIp", CheckClientIp)
}

// @Summary Check Client Ip
// @Schemes
// @Description Health Check from Client Ip
// @Tags Common
// @Accept json
// @Produce json
// @Success 200 {string} CheckClientIp
// @Router /common/CheckClientIp [get]
func CheckClientIp(g *gin.Context) {
	six := g.ClientIP()
	g.JSON(http.StatusOK, fmt.Sprintf("Check Client Ip Ipv6: %v,  IPv4: %v ", six, net.ParseIP(six).To4()))
}
