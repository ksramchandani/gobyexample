package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {

	hostPortsUrl := "http://10.1.2.3/v2/listings/1985/hosts/10.0.0.160/ports"

	data := url.Values{}
	data.Set("dt", "2020-09-14 20:26:21.144556")
	data.Add("opened", "true")
	data.Add("ports", "21")
	data.Add("ports", "22")
	data.Add("ports", "23")

	req, _ := http.NewRequest("PUT", hostPortsUrl, strings.NewReader(data.Encode()))

	req.Header.Add("x-auth", "hydra-producer")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}
