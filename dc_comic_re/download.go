package main

import (
	"bytes"
	"net/http"
	"time"

	"github.com/petershen0307/goLogger/logClient"
)

func download(mangaURL string) (*bytes.Buffer, bool) {
	timeoutRequest := http.Client{Timeout: time.Minute * 5}
	response, err := timeoutRequest.Get(mangaURL)
	if err != nil {
		logClient.Log(logClient.LevelErr, "url:%v, err:%v", mangaURL, err)
		return nil, false
	}
	defer response.Body.Close()
	if http.StatusOK != response.StatusCode {
		logClient.Log(logClient.LevelErr, "url:%v, status:%v", mangaURL, response.StatusCode)
		return nil, false
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)
	return buf, true
}
