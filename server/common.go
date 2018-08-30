package main

import (
	"io/ioutil"
	"net/http"

	"github.com/asmexie/go-logger/logger"
)

func doRequest(url string) (content []byte, err error) {
	response, err := http.Get(url)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return contents, nil
}
