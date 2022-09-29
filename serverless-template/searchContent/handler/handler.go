package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/aws/aws-lambda-go/events"
	"io"
	"log"
	"net/http"
)

type Result struct {
	Name string              `json:"name""`
	Rate map[int]int         `json:"rate""`
	List []map[string]string `json:"list"'`
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

	req, err := http.NewRequest("GET", "https://nodocchi.moe/api/listuser.php", nil)

	params := req.URL.Query()
	params.Add("name", "ハゲ村の会長")
	req.URL.RawQuery = params.Encode()

	fmt.Println(req.URL.String())

	client := &http.Client{}
	r, err := client.Do(req)

	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	var b Result
	json.Unmarshal(body, &b)

	fmt.Println(b)
	fmt.Println("status")
	fmt.Println(res)
	fmt.Println(err)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("===========doc")
	fmt.Println(doc)

	return events.APIGatewayProxyResponse{}, nil
}
