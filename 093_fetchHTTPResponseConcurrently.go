package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func fetch(jsonurl string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", jsonurl, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/57.0.2987.133 Safari/537.36")
	res, err := client.Do(req)
	if err != nil {
		return "", nil
	}

	defer res.Body.Close()
	jsonByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(jsonByte), nil
}

func main() {

	urls := []string{
		"http://www.reddit.com/r/aww.json",
		"http://www.reddit.com/r/funny.json",
		"http://www.reddit.com/r/programming.json",
	}

	responseC := make(chan string)
	errorC := make(chan error)

	for _, jsonurl := range urls {
		go func() {
			// Fetch the response & errors and send them on appropriate channels
			res, err := fetch(jsonurl)
			if err != nil {
				errorC <- err
				return
			}
			responseC <- res
		}()
	}

	for i := 0; i < len(urls); i++ {
		select {
		case response := <-responseC:
			fmt.Println("Response", response)
		case err := <-errorC:
			fmt.Println("Error", err)
		}
	}
}
