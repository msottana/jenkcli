package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"

	"github.com/spf13/viper"
)

type Config struct {
	Protocol  string `json:"protocol"`
	ServerUrl string `json:"serverUrl"`
	User      string `json:"user"`
	Token     string `json:"token"`
}

func ReadConfig() Config {

	viper.SetConfigName("jenkcli-auth")    // name of config file (without extension)
	viper.SetConfigType("yaml")            // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")               // optionally look for config in the working directory
	viper.AddConfigPath("$HOME/.jenkcli/") // optionally look for config in the .jenkins directory inside $HOME

	err := viper.ReadInConfig()

	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found
			fmt.Println("### Config file not found, using CLI authentication ###")
			config := readConfig()

			if err != nil {
				panic(err)
			}

			fmt.Println()

			return config
		} else {
			// Config file was found but another error was produced
			panic(err)
		}
	}

	// Config file found and successfully parsed

	config := Config{
		Protocol:  viper.GetString("protocol"),
		ServerUrl: viper.GetString("server_url"),
		User:      viper.GetString("user"),
		Token:     viper.GetString("token"),
	}

	return config
}

func readValue(reader *bufio.Reader, valueName string) string {
	fmt.Print("Enter " + valueName + ":")
	value, err := reader.ReadString('\n')

	if err != nil {
		panic(err)
	}

	return strings.Trim(value, " \n")
}

func readConfig() Config {
	reader := bufio.NewReader(os.Stdin)

	protocol := readValue(reader, "Server protocol")
	serverUrl := readValue(reader, "Server url")
	username := readValue(reader, "Username")

	fmt.Print("Enter Token: ")
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		panic(err)
	}

	token := string(bytePassword)

	return Config{
		Protocol:  protocol,
		ServerUrl: serverUrl,
		User:      username,
		Token:     token,
	}
}
