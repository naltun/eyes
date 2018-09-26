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
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	res, err := client.Get(domain)
	if err != nil {
		return
	}
	response = *res
	return
}

func Parseoutput(response http.Response) {
	fmt.Println(response.Proto)
	fmt.Println("Content-Length:", response.ContentLength)
	head := response.Header
	for k := range head {
		fmt.Printf("%s : %s \n", k, head[k][0])
	}
}
