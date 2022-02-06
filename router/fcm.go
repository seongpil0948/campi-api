package router

import (
	"fmt"
	"log"
	"net/http"

	fbapp "campi/api/service/fire/app"

	"firebase.google.com/go/messaging"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/iterator"
)

func SetFcmRoutes(g *gin.RouterGroup) {

	g.POST("/fcm/push", SamplePush)
}

// @Summary Pushing Message
// @Schemes
// @Description push messaging
// @Tags Push
// @Accept json
// @Produce json
// @Success 200 {string}
// @Router /fcm/push [post]
func SamplePush(g *gin.Context) {
	var param_tokens, acc_tokens []string
	println("In Sample Push")
	app := fbapp.GetFireInstance()
	msgClient, _ := app.Inst.Messaging(app.Ctx)
	uIds := g.PostFormArray("userIds")
	storeClient, _ := app.Inst.Firestore(app.Ctx)
	param_tokens = g.PostFormArray("tokens")
	iter := storeClient.Collection("users").Where("userId", "in", uIds).Documents(app.Ctx)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		} else if err != nil {
			log.Fatalln(err)
		}
		acc_tokens = doc.Data()["messageToken"].([]string)
		for _, t := range acc_tokens { // Add User Token If not in Param Token
			exist := false
			for _, r := range param_tokens {
				if t == r {
					exist = true
				}
			}
			if !exist {
				param_tokens = append(param_tokens, t)
			}
		}

		fmt.Println()
	}

	message := &messaging.MulticastMessage{
		Data: map[string]string{
			"data1": g.PostForm("data1"),
			"data2": g.PostForm("data2"),
			"data3": g.PostForm("data3"),
		},
		Notification: &messaging.Notification{
			Title: g.PostForm("title"),
			Body:  g.PostForm("body"),
		},
		Tokens: param_tokens,
		// Tokens: []string{
		// 	"eMJPDZgwD0dkmdkGmEX3eR:APA91bEwDknN_4eX_Kz3zXJpKn5adPx79t0nSSrxY0FuW_mp-K8ZsFT3eB4AQ0tjujcuYZuYA8hSjY43K1WBg-xMC6uX_7XK90gwYz-rZbyDAdntgxYS4jzwfLhn1Xaa7wWAnEgsE22i",
		// 	"eQQvrew4S1OLtYH800Thhl:APA91bFsZ8g_31QGFJ8qa0pebgfXPAq7t5OBC9PbTjr5-3XulviqfwUkptY7HK7o2KDHzlixDMrRHZKFYT8KCyfpMX46LZjv8au1iD_C-_mZoQRqBeHCzhxr0cgrSWd4gr9hqVMgK7ly",
		// 	"c7-VtIhFTKqsz8-X7N5pQP:APA91bGbjcWHn55qSnu2pXTid2qO6Rum3yAiCzgkb7MOhkLBxDwitoyygsrSBcDQL2uKNZmaiBAUiu66FCIow5ZUJ8pkpR0HvtUAV-jM33ABVlInd7-C9a50Avn4KX5h4Ls6_paNfvEO",
		// },
	}
	response, err := msgClient.SendMulticast(app.Ctx, message)
	if err != nil {
		log.Fatalln(err)
	}
	// Response is a message ID string.
	g.JSON(http.StatusOK, fmt.Sprintf("Successfully sent message %v", response))
}
