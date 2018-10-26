package robots

import (
	"io/ioutil"
	"net/http"
)

func Get(domain string) string {
	res, err := http.Get(domain)
	if err != nil {
		return err.Error()
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err.Error()
	}

	return string(body)
}
