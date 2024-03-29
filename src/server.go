package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

type tResponse struct {
	Method  string
	Proto   string
	Host    string
	URL     string
	Request tRequest
}

type tRequest struct {
	Params  map[string][]string
	Body    string
	Headers map[string][]string
}

type handler struct{}

func (h *handler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	if len(CLI.ResponseDelay) == 1 {
		time.Sleep(time.Duration(CLI.ResponseDelay[0]) * time.Millisecond)
	}
	if len(CLI.ResponseDelay) == 2 {
		sort.Ints(CLI.ResponseDelay)
		time.Sleep(
			time.Duration(
				randomIntRange(CLI.ResponseDelay[0], CLI.ResponseDelay[1]),
			) * time.Millisecond,
		)
	}

	if len(CLI.BasicAuth) > 0 {
		user, pass, ok := req.BasicAuth()
		if !ok {
			lg.LogError("Error parsing sent basic auth", logrus.Fields{})
			resp.WriteHeader(401)
			return
		}
		if user != CLI.BasicAuth[0] || pass != CLI.BasicAuth[1] {
			lg.LogError("Credentials incorrect", logrus.Fields{
				"user": user,
				"pass": pass,
			})
			resp.WriteHeader(401)
			return
		}
	}

	responseCode := parseResponseCode(req.URL.String())
	resp.WriteHeader(responseCode)
	data := tResponse{
		Method: req.Method,
		Proto:  req.Proto,
		Host:   req.Host,
		URL:    req.URL.String(),
		Request: tRequest{
			Headers: reqHeaders(req),
			Body:    queryBody(resp, req),
			Params:  queryParams(resp, req),
		},
	}

	if !CLI.Verbose {
		lg.LogInfo("Got request", logrus.Fields{
			"method":        data.Method,
			"proto":         data.Proto,
			"host":          data.Host,
			"url":           data.URL,
			"response_code": responseCode,
		})
	} else {
		jsonData, err := json.Marshal(data)
		if err != nil {
			lg.LogError(err.Error(), nil)
		}
		lg.LogInfo("Got request", logrus.Fields{"data": string(jsonData)})
	}
	resp.Header().Set("Content-Type", "application/json")
	json.NewEncoder(resp).Encode(data)
}

func queryParams(w http.ResponseWriter, r *http.Request) (qp map[string][]string) {
	qp = make(map[string][]string)
	for k, v := range r.URL.Query() {
		qp[k] = v
	}
	return
}

func queryBody(w http.ResponseWriter, r *http.Request) (body string) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Error parsing requests body: %q\n", err)
	}
	body = string(bodyBytes)
	return
}

func reqHeaders(r *http.Request) (rh map[string][]string) {
	rh = make(map[string][]string)
	for name, values := range r.Header {
		// loop over all values for the name
		arr := []string{}
		arr = append(arr, values...)
		rh[name] = arr
	}
	return
}

func parseResponseCode(s string) (code int) {
	code = 200
	match := rxFindSubMatch(`/status/(?P<code>\d{3})`, "code", s)
	if match != "" {
		matchInt, err := strconv.Atoi(match)
		if err == nil {
			code = matchInt
		}
	}
	return
}

func randomIntRange(min, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if max <= min {
		max = min + 1
	}
	return r.Intn(max-min) + min
}
