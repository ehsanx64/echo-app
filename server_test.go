package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

const (
	defaultUrl = "http://localhost" + DefaultPort
)

type ReqResult struct {
	Body     string
	Request  *http.Request
	Response *http.Response
}

func httpGet(url string) ReqResult {
	verbose := false
	out := ReqResult{}
	client := http.Client{
		Timeout: 30 * time.Second,
	}

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	out.Body = string(body)
	out.Request = req
	out.Response = res

	if verbose {
		if res.StatusCode != http.StatusOK {
			fmt.Println("Non-200 status code detected")
		}

		fmt.Printf("res: %v\n", string(body))
	}

	return out
}

var urls = make(map[string]map[string]string)

func addUrl(u string, t string) {
	urls[u] = map[string]string{
		"url":  defaultUrl + u,
		"test": t,
	}
}

func Test_main(t *testing.T) {
	addUrl("/", "Hello!!!\n")
	addUrl("/welcome", "Welcome User!!!\n")
	addUrl("/welcome/em", "Welcome EM!!!\n")
	addUrl("/welcome/adam", "Welcome Adam!!!\n")

	for k, v := range urls {
		fmt.Println("<=", k)
		req := httpGet(v["url"])
		fmt.Println("=>", req.Body)
		if req.Body != v["test"] {
			t.Error("!!! "+v["test"]+" <=> ", req.Body)
		}
	}
}
