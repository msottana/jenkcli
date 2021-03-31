package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type QueueResponse struct {
	Class   string `json:"_class"`
	Actions []struct {
		Class  string `json:"_class"`
		Causes []struct {
			Class            string `json:"_class"`
			ShortDescription string `json:"shortDescription"`
			UserID           string `json:"userId"`
			UserName         string `json:"userName"`
		} `json:"causes"`
		Parameters []struct {
			Class string      `json:"_class"`
			Name  string      `json:"name"`
			Value interface{} `json:"value"`
		} `json:"parameters"`
	} `json:"actions"`
	Blocked      bool   `json:"blocked"`
	Buildable    bool   `json:"buildable"`
	ID           int64  `json:"id"`
	InQueueSince int64  `json:"inQueueSince"`
	Params       string `json:"params"`
	Stuck        bool   `json:"stuck"`
	Task         struct {
		Class string `json:"_class"`
		Color string `json:"color"`
		Name  string `json:"name"`
		URL   string `json:"url"`
	} `json:"task"`
	Timestamp  int64  `json:"timestamp"`
	URL        string `json:"url"`
	Why        string `json:"why"`
	Executable struct {
		Class  string `json:"_class"`
		Number int64  `json:"number"`
		URL    string `json:"url"`
	} `json:"executable"`
}

func Build(protocol string, serverUrl string, user string, token string, jobPath string) *http.Response {
	fullUrl := composeUrl(protocol, serverUrl, user, token, jobPath) + "/build"

	payloadBuf := new(bytes.Buffer)

	req, err := http.NewRequest("POST", fullUrl, payloadBuf)

	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	fmt.Println(req)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	return resp
}

func BuildWithParameters(protocol string, serverUrl string, user string, token string, jobPath string, params []string) QueueResponse {
	fullUrl := composeUrl(protocol, serverUrl, user, token, jobPath) + "/buildWithParameters"

	urlWithParams := addParams(fullUrl, params)

	payloadBuf := new(bytes.Buffer)

	req, err := http.NewRequest("POST", urlWithParams, payloadBuf)

	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)

	location := resp.Header.Get("Location")

	protocolLen := len(protocol) + 3

	queueResponse := getQueueResponse(location[0:protocolLen] + user + ":" + token + "@" + location[protocolLen:] + "/api/json?pretty=true")

	return queueResponse
}

func composeUrl(protocol string, serverUrl string, user string, token string, jobPath string) string {
	return protocol + "://" + user + ":" + token + "@" + serverUrl + "/job/" + jobPath
}

func addParams(url string, params []string) string {
	newUrl := url + "?"
	if len(params) > 0 {
		newUrl += params[0]
		for i := 1; i < len(params); i++ {
			newUrl += "&" + params[i]
		}
	}
	return newUrl
}

func getQueueResponse(url string) QueueResponse {
	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	reader := bytes.NewReader(body)
	decoder := json.NewDecoder(reader)

	queueResponse := &QueueResponse{}
	err = decoder.Decode(queueResponse)

	if err != nil {
		panic(err)
	}

	return *queueResponse
}
