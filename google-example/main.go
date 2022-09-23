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
		Token: "fsGS3vDzRJynZYK22h0rQt:APA91bFipz2cTbQYxRoEM6tt38-iW_2yxxlWaU7HhEDMkxZHR6i7RYLN_n_23BbWR0OrJ8H-8jzIQSqBs3ebFn6xx9ThqmT-IM41gD2t8G9-JK7cUzBPH60b099JwCKSHQZcFxO-1Lg5",
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
