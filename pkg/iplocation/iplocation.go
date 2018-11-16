package iplcation

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type IpLocation struct {
	Ip          string `json:"ip"`
	City        string `json:"city"`
	Region      string `json:"region"`
	Country     string `json"city"`
	Coordinates string `json:"loc"`
}

func Find(domain string) []string {
	var ipLocation map[string]interface{}

	url := "https://ipinfo.io/" + domain + "/geo"
	res, err := http.Get(url)
	if err != nil {
		return err.Error()
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err.Error()
	}

	json.Unmarshal([]byte(body), &ipLocation)

	return ipLocation
}
