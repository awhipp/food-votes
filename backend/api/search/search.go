package search

import (
	"food-votes/structs"

	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
)

func Request(zipcode string) structs.Response {
	// Get Foursquare API credentials from environment variables
	apiKey := os.Getenv("FOURSQUARE_API_KEY")

	baseURL := "https://api.foursquare.com"
	resource := "/v3/places/search"
	params := url.Values{}
	params.Add("near", zipcode)
	params.Add("limit", "50")
	params.Add("categories", "13000")

	u, _ := url.ParseRequestURI(baseURL)
	u.Path = resource
	u.RawQuery = params.Encode()
	urlStr := fmt.Sprintf("%v", u)
	fmt.Println("URL: ", urlStr)

	req, _ := http.NewRequest("GET", urlStr, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", apiKey)

	res, _ := http.DefaultClient.Do(req)

	if res.StatusCode != 200 {
		log.Fatal("Error: ", res.Status)
	}

	defer res.Body.Close()

	// Decode JSON response into Response struct
	body := structs.Response{}
	json.NewDecoder(res.Body).Decode(&body)

	return body
}
