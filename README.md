# APIDebugger
This tool uses to debug for API Request.

## Usage

```
$go run main.go -u 'http://www.yahoo.co.jp' -m 'POST' -t params -p 'test1=value1&test2=value2&test3=value3' 
```

## Options

```
-f string
If 't' is 'path', this parameter is required. Set file path to export for request body
-m string
Invalid value are 'GET', 'POST', 'PUT', 'DELETE', 'HEAD' (default "GET")
-p string
If 't' is 'params', this parameter is required. Set request parameters
-t string
Invalid value are 'params', 'path' (default "params")
-u string
Invalid value are http(s)~
```

## Run example

```
// GET(no request parameters)
$go run main.go -u 'http://www.yahoo.co.jp'

// GET(no request parameters, has method setting)
$go run main.go -u 'http://www.yahoo.co.jp' -m 'GET'

// POST(has request parameters)
$go run main.go -u 'http://www.yahoo.co.jp' -m 'POST' -t params -p 'test1=value1&test2=value2&test3=value3'

// POST(has request body's file)
$go run main.go -u 'http://www.yahoo.co.jp' -m 'POST' -t path -p ../json.txt
```

