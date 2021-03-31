package main

import (
	"fmt"
	"os"
)

const HELP string = "help"
const START string = "start"
const STATUS string = "status"

func main() {
	args := os.Args[1:]
	jobPath := args[0]
	var params []string

	config := ReadConfig()

	switch args[0] {
	case HELP:
		help()
		return
	case START:
		fmt.Println("Connecting to server " + config.ServerUrl + " with protocol " + config.Protocol)
		fmt.Println("User authenticated as: " + config.User)

		fmt.Println("Starting job: " + jobPath)
		if len(args) == 2 {
			resp := Build(
				config.Protocol,
				config.ServerUrl,
				config.User,
				config.Token,
				jobPath,
			)

			fmt.Println("Response Status:", resp.Status)
			fmt.Println("Response Headers:", resp.Header)
		} else {
			params = args[2:]

			response := BuildWithParameters(
				config.Protocol,
				config.ServerUrl,
				config.User,
				config.Token,
				jobPath,
				params,
			)

			fmt.Print("Build started with parameters:")
			fmt.Println(response.Params)
		}
	case STATUS:
		fmt.Println("Status:")
		// TODO WIP
	default:
		fmt.Println("Use help to get instructions")
	}

}

func help() {
	fmt.Println("Jenkcli is a simple Jenkins client written in Go")
	fmt.Println("To start a build without parameters use: jenkcli start <build/path>")
	fmt.Println("To start a build with parameters use: jenkcli start <build/path> <param1>=<value1> <param2>=<value2>")
}
