package main

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
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

type UrlTest struct {
	Url    string
	Test   string
	Method string
	Data   []byte
}

func httpCall(u UrlTest) ReqResult {
	verbose := false
	data := bytes.NewBuffer([]byte(u.Data))
	out := ReqResult{}
	client := http.Client{
		Timeout: 30 * time.Second,
	}

	req, err := http.NewRequestWithContext(context.Background(), u.Method, u.Url, data)
	if err != nil {
		panic(err)
	}

	if string(u.Data) != "" {
		// Following must be added in order for urlencoded post request body to be submitted correctly
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
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
		fmt.Printf("req: %v\n", req)
	}

	return out
}

var urls = make(map[string]UrlTest)

func addUrl(url string, u UrlTest) {
	method := http.MethodGet

	if u.Method != "" {
		method = u.Method
	}

	urls[url] = UrlTest{
		Url:    defaultUrl + url,
		Test:   u.Test,
		Method: method,
		Data:   u.Data,
	}
}

func Test_main(t *testing.T) {
	addUrl("/", UrlTest{
		Test: "Hello!!!\n",
	})

	addUrl("/welcome", UrlTest{
		Test: "Welcome User!!!\n",
	})

	addUrl("/welcome/em", UrlTest{
		Test: "Welcome EM!!!\n",
	})

	addUrl("/welcome/adam", UrlTest{
		Test: "Welcome Adam!!!\n",
	})

	addUrl("/users/show?team=alpha&member=subzero", UrlTest{
		Test: fmt.Sprintf("team: %s, member: %s\n", "alpha", "subzero"),
	})

	addUrl("/users/show?team=alpha&member=", UrlTest{
		Test: fmt.Sprintf("team: %s, member: %s\n", "alpha", ""),
	})

	// Generate request body for /users/save endpoint
	saveData := url.Values{}
	saveData.Set("name", "Adam Smith")
	saveData.Set("email", "adam@earth.org")

	addUrl("/users/save", UrlTest{
		Test:   fmt.Sprintf("name: %s, email: %s\n", "Adam Smith", "adam@earth.org"),
		Method: http.MethodPost,
		Data:   []byte(saveData.Encode()),
	})

	for k, v := range urls {
		fmt.Println("<=", k)
		req := httpCall(v)
		fmt.Println("=>", req.Body)
		if req.Body != v.Test {
			t.Error("!!! "+v.Test+" <=> ", req.Body)
		}
	}
}
