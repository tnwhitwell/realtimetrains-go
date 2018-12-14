package rtt

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func Request(url string) (SearchResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%v/%v", baseURL, url), nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return SearchResponse{}, nil
	}
	username, ok := os.LookupEnv("RTT_USER")
	if !ok {
		log.Fatal("RTT_USER was not set")
	}
	pass, ok := os.LookupEnv("RTT_PASS")
	if !ok {
		log.Fatal("RTT_PASS was not set")
	}
	req.SetBasicAuth(username, pass)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return SearchResponse{}, nil
	}
	defer resp.Body.Close()
	var sr SearchResponse

	if err := json.NewDecoder(resp.Body).Decode(&sr); err != nil {
		log.Println(err)
		return SearchResponse{}, err
	}
	return sr, nil
}

func Search(from string, to string, dt time.Time) (SearchResponse, error) {
	searchstr := fmt.Sprintf("/search/%v/to/%v/%v/%v/%v/%v%v",
		from, to, dt.Year(), int(dt.Month()), dt.Day(), dt.Hour(), dt.Minute())
	log.Println(searchstr)
	sr, err := Request(searchstr)
	return sr, err
}
