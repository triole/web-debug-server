# Web Debug Server ![build](https://github.com/triole/web-debug-server/actions/workflows/build.yaml/badge.svg)

<!--- mdtoc: toc begin -->

1. [Synopsis](#synopsis)
2. [Usage Examples](#usage-examples)
3. [Custom Response Codes](#custom-response-codes)
4. [Help](#help)<!--- mdtoc: toc end -->

## Synopsis

A simple web debug server echoing requests. Send a something and retrieve a json response resembling the request's information.

## Usage Examples

$ curl http://localhost:9999?name=john

```json
{
  "Host": "localhost:9999",
  "Method": "GET",
  "Proto": "HTTP/1.1",
  "Request": {
    "Body": "",
    "Headers": {
      "Accept": [
        "*/*"
      ],
      "User-Agent": [
        "curl/7.74.0"
      ]
    },
    "Params": {
      "name": [
        "john"
      ]
    }
  },
  "URL": "/?name=james"
}
```

$ curl -X POST -F 'os=linux' http://localhost:9999

```json
{
  "Host": "localhost:9999",
  "Method": "POST",
  "Proto": "HTTP/1.1",
  "Request": {
    "Body": "--------------------------d11f137cf722c9c9\r\nContent-Disposition: form-data; name=\"os\"\r\n\r\nlinux\r\n--------------------------d11f137cf722c9c9--\r\n",
    "Headers": {
      "Accept": [
        "*/*"
      ],
      "Content-Length": [
        "142"
      ],
      "Content-Type": [
        "multipart/form-data; boundary=------------------------d11f137cf722c9c9"
      ],
      "User-Agent": [
        "curl/7.74.0"
      ]
    },
    "Params": {}
  },
  "URL": "/"
}
```

## Custom Response Codes

Custom response status codes can be set through the called url. Use `/status/` and append the required integer. I.e. `/status/404` responds with a 404 status code and so on.

## Help

```go mdox-exec="r -h"

simple web server for debugging purposes

Flags:
  -h, --help                      Show context-sensitive help.
  -p, --port=9999                 port where to serve
  -r, --response-delay=RESPONSE-DELAY,...
                                  server response delay in ms, use twice to
                                  define range for a random value before each
                                  request
  -j, --json-log                  enable json log, instead of text one
  -l, --log-file="/dev/stdout"    log file
  -v, --verbose                   verbose mode, print full response data set
  -V, --version-flag              display version
```
