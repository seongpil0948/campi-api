package main

import (
	"campi/api/core/model"
	"campi/api/router"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// 파일에 로그를 작성합니다.
	f, _ := os.Create("./public/gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	// // https://firebase.google.com/docs/cloud-messaging/send-message?hl=ko
	// opt := option.WithCredentialsFile("secret/campi-f8278-firebase-adminsdk-hqj7j-ea56733369.json")
	// ctx := context.Background()
	// app, _ := firebase.NewApp(ctx, nil, opt)
	// msgClient, _ := app.Messaging(ctx)
	// message := &messaging.Message{
	// 	Data:  map[string]string{"score": "850", "time": "2:45"},
	// 	Token: config.FcmServerKey,
	// 	// Topic:        "",
	// 	// Condition:    "",
	// }
	// fmt.Printf("%s", app)

	// // Send a message to the device corresponding to the provided
	// // registration token.
	// response, err := msgClient.Send(ctx, message)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// // Response is a message ID string.
	// fmt.Println("Successfully sent message:", response)

	router := router.InitRoutes()
	model.SampleOrm()
	router.Run() // 서버가 실행 되고 0.0.0.0:8080 에서 요청을 기다립니다.
}
