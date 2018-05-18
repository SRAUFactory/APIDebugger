package main

import "flag"
import "fmt"
import "io"

var (
	url      = flag.String("u", "", "Invalid value are http(s)~")
	method   = flag.String("m", "GET", "Invalid value are 'GET', 'POST', 'PUT', 'DELETE', 'HEAD'")
	bodyType = flag.String("t", "params", "Invalid value are 'params', 'path'")
	params   = flag.String("p", "", "If 't' is 'params', this parameter is required. Set request parameters")
	path     = flag.String("f", "", "If 't' is 'path', this parameter is required. Set file path to export for request body")
)

func requestHttp(url string, method string, body string) ([]byte, error) {
	req, err := http.NewRequest("GET", method, bytes.NewBuffer([]byte(body)))
	if err != nil {
		return nil, err
	}

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body), nil
}

func main() {
	flag.Parse()
	fmt.Println("url:", *url)
	fmt.Println("method:", *method)
	fmt.Println("bodyType:", *bodyType)
	fmt.Println("params:", *params)
	fmt.Println("path:", *path)
}
