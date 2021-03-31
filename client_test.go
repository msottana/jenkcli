package main

import "testing"

type AddParamsResult struct {
	Url      string
	Params   []string
	Expected string
}

type ComposeUrlResult struct {
	Protocol  string
	ServerUrl string
	User      string
	Token     string
	JobPath   string
	Expected  string
}

var addParamsResults = []AddParamsResult{
	{"url", []string{"a=1"}, "url?a=1"},
	{"url", []string{"a=1", "b=2"}, "url?a=1&b=2"},
}

var composeUrlResults = []ComposeUrlResult{
	{"http", "localhost", "user", "1234", "my/job", "http://user:1234@localhost/job/my/job"},
	{"http", "localhost:8080", "user", "1234", "my/job", "http://user:1234@localhost:8080/job/my/job"},
	{"http", "localhost:8080/jenkins", "user", "1234", "my/job", "http://user:1234@localhost:8080/jenkins/job/my/job"},
	{"https", "localhost", "user", "1234", "my/job", "https://user:1234@localhost/job/my/job"},
	{"https", "localhost:8080", "user", "1234", "my/job", "https://user:1234@localhost:8080/job/my/job"},
	{"https", "localhost:8080/jenkins", "user", "1234", "my/job", "https://user:1234@localhost:8080/jenkins/job/my/job"},
}

func TestAddParams(t *testing.T) {
	for _, test := range addParamsResults {
		result := addParams(test.Url, test.Params)
		if result != test.Expected {
			t.Error("Expected: " + test.Expected + " but Result: " + result)
		}
	}
}

func TestComposeUrl(t *testing.T) {
	for _, test := range composeUrlResults {
		result := composeUrl(test.Protocol, test.ServerUrl, test.User, test.Token, test.JobPath)
		if result != test.Expected {
			t.Error("Expected: " + test.Expected + " but Result: " + result)
		}
	}
}
