package main

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"fmt"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()
	key := option.WithCredentialsFile("ohsho-shutsujin-firebase-adminsdk-f2en5-a310576096.json")

	app, err := firebase.NewApp(ctx, nil, key)
	if err != nil {
		fmt.Println("firebaseNewAppError")
		return
	}
	client, err := app.Messaging(ctx)
	if err != nil {
		fmt.Println("firebaseNewClientError")
		return
	}

	m := &messaging.Message{
		Notification: &messaging.Notification{
			Title: "title",
			Body:  "body",
		},
		Token: "fZoxKPf6R3Ce-jNWV3LM6f:APA91bFEb0n9Rpb97EBgOqYWGZSYBTiGXtcBxXXunJ-BopCagCejGzFyCQsHPlUWUm1-au-ZSieP52f-wYmd1psgguiUNLS8YH-fIQcHqo_fFFaRzRfe6ROcQVugmPyEjI-EZ85Xsjzf",
		//Android: &messaging.AndroidConfig{
		//	Priority: "high",
		//	Notification: &messaging.AndroidNotification{
		//		Title: "titleAndroid",
		//		Body:  "bodyAndroid",
		//	},
		//},
	}

	fmt.Println(m)
	res, err := client.Send(ctx, m)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)

	fmt.Println(err)
}
