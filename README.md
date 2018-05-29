# APIDebugger
This tool uses to debug for API Request.

## Usage

```
$go run main.go -u 'http://www.yahoo.co.jp' -m 'POST' -b params -p 'test1=value1&test2=value2&test3=value3' 
```

## Options

```
  -b string
    	Invalid value are 'params', 'path' (default "params")
  -f string
    	If 'b' is 'path', this parameter is required. Set file path to export for request body
  -m string
    	Invalid value are 'GET', 'POST', 'PUT', 'DELETE', 'HEAD' (default "GET")
  -p string
    	If 'b' is 'params', this parameter is required. Set request parameters
  -t string
    	Set authorization's token
  -u string
    	Invalid value are http(s)~
```

## Run example

```
// GET(no request parameters)
$go run main.go -u 'http://www.yahoo.co.jp'

// GET(no request parameters, has method setting)
$go run main.go -u 'http://www.yahoo.co.jp' -m 'GET'

// GET(no request parameters, has method & authorization's token setting)
$go run main.go -u 'http://www.yahoo.co.jp' -m 'GET' -t '******************'

// POST(has request parameters)
$go run main.go -u 'http://www.yahoo.co.jp' -m 'POST' -b params -p 'test1=value1&test2=value2&test3=value3'

// POST(has request body's file)
$go run main.go -u 'http://www.yahoo.co.jp' -m 'POST' -b path -p ../json.txt
```

