package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	response *http.Response
	err      error
)

type ResponseBody struct {
	Name string `json:"name"`
}

type URLFetcher interface {
	GetVersion(string) error
}

func (r *ResponseBody) GetVersion(getURL string) error {
	response, err = http.Get(getURL)
	if err != nil {
		return err
	}

	defer response.Body.Close()
	if err := json.NewDecoder(response.Body).Decode(r); err != nil {
		return err
	}
	fmt.Println(response.StatusCode)

	return nil
}

func GetLatestRelease(repo string, u URLFetcher) error {
	latestURL := fmt.Sprintf("https://api.github.com/repos/%s/releases/latest", repo)

	if err := u.GetVersion(latestURL); err != nil {
		fmt.Println("Error occuring during fetching url", err.Error())
		return err
	}

	return nil
}

func main() {

	repo := "docker/machine"
	b := ResponseBody{}
	if err := GetLatestRelease(repo, &b); err != nil {
		fmt.Println("Error occured during fetching url", err.Error())
		return
	}

	fmt.Printf("Response body %+v\n", b.Name)

}
