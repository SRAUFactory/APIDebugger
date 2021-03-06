package main

import "bytes"
import "flag"
import "fmt"
import "io/ioutil"
import "net/http"

var (
	argURL      = flag.String("u", "", "Invalid value are http(s)~")
	argMethod   = flag.String("m", "GET", "Invalid value are 'GET', 'POST', 'PUT', 'DELETE', 'HEAD'")
	argBodyType = flag.String("b", "params", "Invalid value are 'params', 'path'")
	argParams   = flag.String("p", "", "If 'b' is 'params', this parameter is required. Set request parameters")
	argPath     = flag.String("f", "", "If 'b' is 'path', this parameter is required. Set file path to export for request body")
	argToken    = flag.String("t", "", "Set authorization's token")
)

func requestHTTP(url string, method string, body string, token string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	if err != nil {
		return nil, err
	}
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}

	client := new(http.Client)
	return client.Do(req)
}

func printResponseContent(res *http.Response, err error) {
	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()
	byteArray, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(byteArray))
}

func printArgs() {
	fmt.Println("url: ", *argURL)
	fmt.Println("method: ", *argMethod)
	fmt.Println("bodyType: ", *argBodyType)
	fmt.Println("params: ", *argParams)
	fmt.Println("path: ", *argPath)
	fmt.Println("token: ", *argToken)
}

func main() {
	flag.Parse()
	printArgs()

	body := *argParams
	// @ToDo Implement to set body for case of bodyType == 'path'

	res, err := requestHTTP(*argURL, *argMethod, body, *argToken)
	printResponseContent(res, err)
}
