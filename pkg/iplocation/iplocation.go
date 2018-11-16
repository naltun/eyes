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

func Find(domain string) (IpLocation, error) {
	var ipLocation IpLocation

	url := "https://ipinfo.io/" + domain + "/geo"
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	json.Unmarshal([]byte(body), &ipLocation)

	return ipLocation
}
