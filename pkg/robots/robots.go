package robots

import (
	"io/ioutil"
	"net/http"
)

func Get(domain string) string {
	res, _ := http.Get(domain)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	return string(body)
}
