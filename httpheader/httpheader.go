package httpheader

import (
	"fmt"
	"net/http"
)

func Httpheader(domain string) (response http.Response, err error) {

	if domain == "" {
		err = fmt.Errorf("Domain is empty")
		return
	}
	client := &http.Client{}

	res, err := client.Get(domain)
	if err != nil {
		return
	}
	response = *res
	return
}
