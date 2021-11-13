package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type tResponse struct {
	Method string
	Proto  string
	Host   string
	URL    string
}

type handler struct{}

func (h *handler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	data := tResponse{
		Method: req.Method,
		Proto:  req.Proto,
		Host:   req.Host,
		URL:    fmt.Sprintf("%s", req.URL),
	}
	log.Printf("[INFO] %s %s %s%s", data.Method, data.Proto, data.Host, data.URL)
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusCreated)
	json.NewEncoder(resp).Encode(data)
}
