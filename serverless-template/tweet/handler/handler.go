package handler

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"github.com/aws/aws-lambda-go/events"
	"io"
	"log"
	"math"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

type APIResult struct {
	Request struct {
		Parameters map[string]string `json:"parameters"`
	} `json:"request"`
	Result struct {
		Items []Item `json:"items"`
	}
}

type Item struct {
	Title        string `json:"title"`
	URL          string `json:"URL"`
	AffiliateURL string `json:"affiliateURL"`
	Date         string `json:"date"`
}

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	apiKey := os.Getenv("TWITTER_API_KEY")
	apiSecret := os.Getenv("TWITTER_API_SECRET")
	accessToken := os.Getenv("ACCESS_TOKEN")
	accessTokenSecret := os.Getenv("ACCESS_TOKEN_SECRET")
	anaconda.SetConsumerKey(apiKey)
	anaconda.SetConsumerSecret(apiSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)
	text := "Test Tweet"

	fmt.Println(api)
	fmt.Println(text)

	im, err := http.Get("https://cc3001.dmm.co.jp/litevideo/freepv/h/hoi/hoisw00018/hoisw00018_dmb_w.mp4")
	fmt.Println(im)

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	data, err := io.ReadAll(im.Body)
	defer im.Body.Close()

	//base64Data := base64.StdEncoding.EncodeToString(body)

	//fmt.Println(base64Data)

	//var base64String = "/9j/4AAQSk...LT09DojFy1Z//2Q=="
	//media, _ := api.UploadMedia(base64Data)
	media, err := api.UploadVideoInit(len(data), "video/mp4")
	fmt.Println(media)

	chunkIndex := 0
	for i := 0; i < len(data); i += 5242879 {
		log.Println("Chunk", chunkIndex)
		err = api.UploadVideoAppend(media.MediaIDString, chunkIndex,
			base64.StdEncoding.EncodeToString(
				data[i:int(math.Min(5242879.0, float64(len(data))))],
			),
		)
		if err != nil {
			return events.APIGatewayProxyResponse{}, err
		}
		chunkIndex++
	}

	videoMedia, err := api.UploadVideoFinalize(media.MediaIDString)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	var ids string
	ids += videoMedia.MediaIDString
	ids += ","

	ids = ids[:len(ids)-1]
	log.Println("media_ids:", ids)
	v := url.Values{}
	v.Set("media_ids", ids)
	v.Set("possibly_sensitive", strconv.FormatBool(true))

	v.Add("media_ids", media.MediaIDString)
	tweet, err := api.PostTweet(text, v)
	if err != nil {
		panic(err)
	}
	fmt.Println(tweet.Text)

	return events.APIGatewayProxyResponse{}, nil
}
