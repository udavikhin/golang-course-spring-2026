package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type Asset struct {
	Name     string `json:"name"`
	PriceUSD string `json:"priceUsd"`
}

type GetAssetsResponse struct {
	Data []Asset `json:"data"`
}

type CoincapClient struct {
	httpClient *http.Client
	apiKey     string
	baseURL    string
}

func NewCoincapClient(apiKey string) *CoincapClient {
	baseURL := "https://rest.coincap.io/v3/"

	return &CoincapClient{
		httpClient: &http.Client{
			Timeout: time.Second * 10,
		},
		apiKey:  apiKey,
		baseURL: baseURL,
	}
}

func (c *CoincapClient) Request(method string, path string) (*http.Response, error) {
	path = strings.TrimLeft(path, "/")
	req, err := http.NewRequest(method, c.baseURL+path, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *CoincapClient) GetAssets() (*GetAssetsResponse, error) {
	res, err := c.Request(http.MethodGet, "assets")
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Printf("Response error status code: %d", res.StatusCode)
		return nil, errors.New("error getting assets")
	}

	assetsRes := GetAssetsResponse{}
	if err := json.NewDecoder(res.Body).Decode(&assetsRes); err != nil {
		log.Println(err)
		return nil, err
	}

	return &assetsRes, nil
}

func main() {
	key := os.Getenv("COINCAP_API_KEY")

	client := NewCoincapClient(key)

	assets, err := client.GetAssets()
	if err != nil {
		log.Println(err)
		return
	}

	for _, asset := range assets.Data {
		fmt.Printf("%+v\n", asset)
	}
}
