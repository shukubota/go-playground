package handler

import (
	"context"
	"encoding/base64"
	"fmt"
	//"github.com/ChimeraCoder/anaconda"
	"github.com/aws/aws-lambda-go/events"
	"github.com/shukubota/go-api-template/serverless-template/tweet/anaconda"
	"log"
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
	text := "良さげ"

	fmt.Println(api)
	fmt.Println(text)

	// 小さいサイズ
	im, err := http.Get("https://sample.mgstage.com/sample/manmanland/476mla/075/476mla-075_20220308T135029.mp4")
	//im, err := http.Get("https://sample.mgstage.com/sample/hametaverse/598hmt/010/598hmt-010_20220916T184301.mp4")
	//im, err := http.Get("https://cc3001.dmm.co.jp/litevideo/freepv/h/hoi/hoisw00018/hoisw00018_dmb_w.mp4")
	fmt.Println(im)

	filePath := "/Users/shukubota/Desktop/media/short.mp4"

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	//bytes, err := io.ReadAll(im.Body)
	bytes, err := os.ReadFile(filePath)
	defer im.Body.Close()
	totalBytes := len(bytes)
	media, err := api.UploadVideoInit(totalBytes, "video/mp4")
	fmt.Println(media)
	fmt.Println(err)

	fmt.Println(len(bytes))

	fmt.Println("=============media")

	//mediaMaxLen := 1 * 1024 * 1024
	mediaMaxLen := 5000000
	segment := 0
	for i := 0; i < totalBytes; i += mediaMaxLen {
		var mediaData string
		fmt.Printf("%v\n", segment)
		if i+mediaMaxLen < totalBytes {
			mediaData = base64.StdEncoding.EncodeToString(bytes[i : i+mediaMaxLen])
		} else {
			mediaData = base64.StdEncoding.EncodeToString(bytes[i:])
		}
		if err = api.UploadVideoAppend(media.MediaIDString, segment, mediaData); err != nil {
			fmt.Println("=========err")
			fmt.Printf("%v\n", err)
			return events.APIGatewayProxyResponse{}, err
		}
		segment += 1
	}

	//chunkIndex := 0
	//mediaMaxLen := 5 * 1024 * 1024
	//for i := 0; i < len(data); i += mediaMaxLen {
	//	log.Println("Chunk", chunkIndex)
	//	err = api.UploadVideoAppend(media.MediaIDString, chunkIndex,
	//		base64.StdEncoding.EncodeToString(
	//			data[i:int(math.Min(float64(mediaMaxLen), float64(len(data))))],
	//		),
	//	)
	//	if err != nil {
	//		fmt.Println("=========err")
	//		fmt.Printf("%v\n", err)
	//		return events.APIGatewayProxyResponse{}, err
	//	}
	//	chunkIndex++
	//}

	fmt.Println("============before UploadVideoFinalize")
	//time.Sleep(time.Second * 3)
	fmt.Println("============before UploadVideoFinalize")
	videoMedia, err := api.UploadVideoFinalize(media.MediaIDString)
	if err != nil {
		fmt.Println("=========err UploadVideoFinalize")
		fmt.Printf("%v\n", err)
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
