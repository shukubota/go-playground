package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"io"
	"log"
	"net/http"
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

//{
//"starttime": "1661871720",
//"during": "11",
//"sctype": "b",
//"playernum": "4",
//"playerlevel": "0",
//"playlength": "1",
//"kuitanari": "1",
//"akaari": "1",
//"player1": "ゆkんm",
//"player1ptr": "58.1",
//"player2": "エイチデニウソン",
//"player2ptr": "4.3",
//"player3": "ハゲ村の会長",
//"player3ptr": "-15.7",
//"player4": "337tn101",
//"player4ptr": "-46.7"
//},

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("----------sssss")
	// ここにscraping処理

	res, err := http.Get("https://qiita.com/timeline")
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
		return events.APIGatewayProxyResponse{}, err
	}

	req, err := http.NewRequest("GET", "https://api.dmm.com/affiliate/v3/ItemList", nil)

	params := req.URL.Query()
	params.Add("api_id", "qrs7udKgm3hNXa4c9AAC")
	params.Add("affiliate_id", "eloquent-990")
	params.Add("site", "FANZA")
	params.Add("service", "digital")
	params.Add("sort", "date")
	params.Add("output", "json")

	req.URL.RawQuery = params.Encode()

	fmt.Println(req.URL.String())

	client := &http.Client{}
	r, err := client.Do(req)

	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	var b APIResult
	json.Unmarshal(body, &b)

	fmt.Println(b.Result.Items[0])
	fmt.Println("status")
	fmt.Println(res)
	//fmt.Println(err)

	return events.APIGatewayProxyResponse{}, nil
}
