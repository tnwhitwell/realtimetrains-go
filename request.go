package rtt

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func Request(url string) (SearchResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%v/%v", baseURL, url), nil)
	if err != nil {
		log.Print("NewRequest: ", err)
		return SearchResponse{}, err
	}
	username, ok := os.LookupEnv("RTT_USER")
	if !ok {
		log.Print("RTT_USER was not set")
		return SearchResponse{}, err
	}
	pass, ok := os.LookupEnv("RTT_PASS")
	if !ok {
		log.Print("RTT_PASS was not set")
		return SearchResponse{}, err
	}
	req.SetBasicAuth(username, pass)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Print("Do: ", err)
		return SearchResponse{}, err
	}
	defer resp.Body.Close()
	var sr SearchResponse
	if resp.StatusCode != 200 {
		b, _ := ioutil.ReadAll(resp.Body)
		log.Print("Non-200 response: ", string(b))
		return SearchResponse{}, fmt.Errorf(string(b))
	}
	if err := json.NewDecoder(resp.Body).Decode(&sr); err != nil {
		return SearchResponse{}, err
	}
	return sr, nil
}

func Search(from string, to string, dt time.Time) (SearchResponse, error) {
	searchstr := fmt.Sprintf("/search/%v/to/%v/%v/%v/%v/%v%v",
		from, to, dt.Year(), int(dt.Month()), dt.Day(), fmt.Sprintf("%02d", dt.Hour()), fmt.Sprintf("%02d", dt.Minute()))
	log.Println(searchstr)
	sr, err := Request(searchstr)
	return sr, err
}
