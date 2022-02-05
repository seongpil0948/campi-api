package router

import (
	"campi/api/docs"
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func InitRoutes() gin.Engine {
	r := gin.Default()
	g := r.Group("api")
	gin.ForceConsoleColor()
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}
	SetAuthRoutes(g.Group("auth"))
	docs.SwaggerInfo.BasePath = "/api"
	g.GET("/helloworld", Helloworld)
	g.GET("/checkIp", checkClientIp)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	g.GET("/push/msg", samplePush)
	return *r
}

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func Helloworld(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

func checkClientIp(g *gin.Context) {
	six := g.ClientIP()
	log.Printf("Check Client Ip Ipv6: %v, \n IPv4: %v ", six, net.ParseIP(six).To4())
	g.JSON(http.StatusOK, "Check Ip API")
}

func samplePush(g *gin.Context) {
	println("In Sample Push")
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("error initializing FCM: %v\n", err)
	}

	message := &messaging.MulticastMessage{
		Data: map[string]string{
			"score": "850",
			"time":  "2:45",
		},
		Notification: &messaging.Notification{
			Title: g.Param("title"),
			Body:  g.Param("body"),
		},
		Tokens: []string{
			"eMJPDZgwD0dkmdkGmEX3eR:APA91bEwDknN_4eX_Kz3zXJpKn5adPx79t0nSSrxY0FuW_mp-K8ZsFT3eB4AQ0tjujcuYZuYA8hSjY43K1WBg-xMC6uX_7XK90gwYz-rZbyDAdntgxYS4jzwfLhn1Xaa7wWAnEgsE22i",
			"eQQvrew4S1OLtYH800Thhl:APA91bFsZ8g_31QGFJ8qa0pebgfXPAq7t5OBC9PbTjr5-3XulviqfwUkptY7HK7o2KDHzlixDMrRHZKFYT8KCyfpMX46LZjv8au1iD_C-_mZoQRqBeHCzhxr0cgrSWd4gr9hqVMgK7ly",
			"c7-VtIhFTKqsz8-X7N5pQP:APA91bGbjcWHn55qSnu2pXTid2qO6Rum3yAiCzgkb7MOhkLBxDwitoyygsrSBcDQL2uKNZmaiBAUiu66FCIow5ZUJ8pkpR0HvtUAV-jM33ABVlInd7-C9a50Avn4KX5h4Ls6_paNfvEO",
		},
	}
	response, err := client.SendMulticast(ctx, message)
	if err != nil {
		log.Fatalln(err)
	}
	// Response is a message ID string.
	g.JSON(http.StatusOK, fmt.Sprintf("Successfully sent message %v", response))
}
