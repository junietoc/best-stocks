package stocks

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type StockItem struct {
	Ticker     string `json:"ticker"`
	TargetFrom string `json:"target_from"`
	TargetTo   string `json:"target_to"`
	Company    string `json:"company"`
	Action     string `json:"action"`
	Brokerage  string `json:"brokerage"`
	RatingFrom string `json:"rating_from"`
	RatingTo   string `json:"rating_to"`
	Time       string `json:"time"`
}

type responseData struct {
	Items    []StockItem `json:"items"`
	NextPage string      `json:"next_page"`
}

func FetchStockData() ([]StockItem, error) {
	bearerToken := os.Getenv("API_BEARER_TOKEN")
	log.Println("fetchingg")
	if bearerToken == "" {
		return nil, fmt.Errorf("API_BEARER_TOKEN is not set")
	}

	var allItems []StockItem
	var nextPage string

	for {
		url := "https://8j5baasof2.execute-api.us-west-2.amazonaws.com/production/swechallenge/list"
		if nextPage != "" {
			url += "?next_page=" + nextPage
		}

		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return nil, err
		}

		req.Header.Set("Authorization", "Bearer "+bearerToken)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		var data responseData
		err = json.Unmarshal(body, &data)
		if err != nil {
			return nil, err
		}

		allItems = append(allItems, data.Items...)

		if data.NextPage == "" {
			break
		}
		nextPage = data.NextPage
	}
	log.Println(len(allItems))
	return allItems, nil
}
