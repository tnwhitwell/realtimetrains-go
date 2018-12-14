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

func Request(url string) ([]byte, error) {
	log.Println(url)
	req, err := http.NewRequest("GET", fmt.Sprintf("%v/%v", baseURL, url), nil)
	if err != nil {
		log.Print("NewRequest: ", err)
		return nil, err
	}
	username, ok := os.LookupEnv("RTT_USER")
	if !ok {
		log.Print("RTT_USER was not set")
		return nil, err
	}
	pass, ok := os.LookupEnv("RTT_PASS")
	if !ok {
		log.Print("RTT_PASS was not set")
		return nil, err
	}
	req.SetBasicAuth(username, pass)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Print("Do: ", err)
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		b, _ := ioutil.ReadAll(resp.Body)
		log.Print("Non-200 response: ", string(b))
		return nil, fmt.Errorf(string(b))
	}
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	return bodyBytes, nil
}

func Search(from string, to string, dt time.Time) (SearchResponse, error) {
	searchstr := fmt.Sprintf("/search/%v/to/%v/%v/%v/%v/%v%v",
		from, to, dt.Year(), int(dt.Month()), dt.Day(), fmt.Sprintf("%02d", dt.Hour()), fmt.Sprintf("%02d", dt.Minute()))
	response, err := Request(searchstr)
	if err != nil {
		log.Print(err)
	}
	var sr SearchResponse
	if err := json.Unmarshal(response, &sr); err != nil {
		return SearchResponse{}, err
	}
	return sr, err
}

func ServiceDetail(serviceUID string, dt time.Time) (ServiceResponse, error) {
	searchstr := fmt.Sprintf("/service/%v/%v/%v/%v", serviceUID, dt.Year(), int(dt.Month()), dt.Day())
	response, err := Request(searchstr)
	if err != nil {
		log.Print(err)
	}
	var sr ServiceResponse
	if err := json.Unmarshal(response, &sr); err != nil {
		return ServiceResponse{}, err
	}
	return sr, err
}
