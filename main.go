package main

import (
	"campi/api/router"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// 파일에 로그를 작성합니다.
	f, _ := os.Create("./public/gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	router := router.InitRoutes()
	// model.SampleOrm()
	router.Run() // 서버가 실행 되고 0.0.0.0:8080 에서 요청을 기다립니다.
}
