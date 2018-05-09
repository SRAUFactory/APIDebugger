package main

import "flag"
import "fmt"

var (
	url      = flag.String("u", "", "Invalid value are http(s)~")
	method   = flag.String("m", "GET", "Invalid value are 'GET', 'POST', 'PUT', 'DELETE', 'HEAD'")
	bodyType = flag.String("t", "params", "Invalid value are 'params', 'path'")
	params   = flag.String("p", "", "If 't' is 'params', this parameter is required. Set request parameters")
	path     = flag.String("f", "", "If 't' is 'path', this parameter is required. Set file path to export for request body")
)

func main() {
	flag.Parse()
	fmt.Println("url:", *url)
	fmt.Println("method:", *method)
	fmt.Println("bodyType:", *bodyType)
	fmt.Println("params:", *params)
	fmt.Println("path:", *path)
}
