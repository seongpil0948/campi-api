package router

import (
	"campi/api/docs"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func InitRoutes() gin.Engine {
	r := gin.Default()
	g := r.Group("api")

	r.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		log.Print("=== CustomRecovery ===")
		if err, ok := recovered.(string); ok {
			c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
		}
		c.AbortWithStatus(http.StatusInternalServerError)
	}))

	gin.ForceConsoleColor()
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}

	SetAuthRoutes(g.Group("auth"))
	SetCommonRoutes(g.Group("common"))
	SetFcmRoutes(g.Group("fcm"))

	docs.SwaggerInfo.BasePath = "/api"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return *r
}
