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
		// tokenダミー
		Token: "cCC7Hyz5h0SsqOmwgShcpf:APA91bF9RwxTjJRB4VaaxXS0bnGznunSrpKoBpdMjMCsj_nOwXtnriu6dFmRW-9KKCWXKy5MZ9FVe5ZAKPWOTXibJyO3Bf4pyGlVKTkWvcjMzvSqa5Z3uRd-mgApIdKx0SXZ-OxSVmxC",
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
