package linkgrab

import (
	"io"
	"net/http"

	"golang.org/x/net/html"
)

/*
** Thanks to github.com/vorozhko for his blog post at
** https://vorozhko.net/get-all-links-from-html-page-with-go-lang for
** helping with this solution
 */

var links1 []string
var links2 []string

func GetLinks(domain string) []string {
	res, _ := http.Get(domain)
	/*
		if err != nil {
			return fmt.Println(err)
		}
	*/
	defer res.Body.Close()

	for _, v := range readLinks(res.Body) {
		links2 = append(links2, v)
	}

	return links2
}

func readLinks(body io.Reader) []string {
	t := html.NewTokenizer(body)

	for {
		tt := t.Next()

		switch tt {
		case html.ErrorToken:
			return links1
		case html.StartTagToken, html.EndTagToken:
			token := t.Token()
			if token.Data == "a" {
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						links1 = append(links1, attr.Val)
					}
				}
			}
		}
	}
}
