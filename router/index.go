package router

import (
	"campi/api/docs"
	"log"
	"net/http"

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
	// g.Use(appid.AppIDMiddleWare())
	// SetHelloRoutes(g)
	// SetAuthRoutes(g) // SetAuthRoutes invoked
	// g.Use(token.TokenAuthMiddleWare())  //secure the API From this line to bottom with JSON Auth
	// g.Use(appid.ValidateAppIDMiddleWare())
	// SetTaskRoutes(g)
	// SetUserRoutes(g)

	docs.SwaggerInfo.BasePath = "/api"
	g.GET("/helloworld", Helloworld)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
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
