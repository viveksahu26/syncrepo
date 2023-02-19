package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"gopkg.in/yaml.v3"
)

var repoConfigInfo = new(config)

type config struct {
	Server Server
	Log    log
}

func GetServerConfig() *Server {
	return &repoConfigInfo.Server
}

func GetLogConfig() *log {
	return &repoConfigInfo.Log
}

// server related configuration such as PORT, Debug, Timeout
type Server struct {
	Port    int
	Debug   bool
	Timeout int
}

// log related configuration such as log level
type log struct {
	Level string
}

// main repo configuration such as repo name, username, token, branch, etc.
type MainRepoConfig struct {
	RepoURL             string
	UserName            string
	Token               string
	Branch              string
	FollowersRepoConfig map[string]FollowersRepoConfig
}

// followerd repo configuration such as repo name, username, token, branch, etc.

type FollowersRepoConfig struct {
	RepoURL    string
	UserName   string
	Token      string
	RemoteName string
	Branch     string
}

// Initialization of Server Info such as PORT, Debug, Timeout
func Init() {
	// What to Initialize: config values
	// from a file
	configFile := os.Getenv("CONFIG_FILE")
	if configFile == "" {
		configFile = "./config/local-config.yaml"
	}

	// read the configFile
	file, err := ioutil.ReadFile(configFile)
	if err != nil {
		fmt.Printf("Error in reading config file: %v", configFile)
	}

	// mapping file into our strcut
	err = yaml.Unmarshal(file, repoConfigInfo)
	if err != nil {
		fmt.Printf("error unmarshalling config file: %v\n", err)
	}
	// If env variables are gives which means to change the default values.
	if debug := os.Getenv("DEBUG"); debug != "" {
		val, err := strconv.ParseBool(debug)
		if err != nil {
			fmt.Printf("error parsing debug env var: %v\n", err)
		}
		repoConfigInfo.Server.Debug = val
	}

	if port := os.Getenv("PORT"); port != "" {
		val, err := strconv.Atoi(port)
		if err != nil {
			fmt.Printf("error parsing port env var: %v\n", err)
		}
		// change the default value of port to value externally provided by user
		repoConfigInfo.Server.Port = val
	}
}

// Initialization of Main Repo as well as followers Repo Info
func ConfidentialInit() {
	// Providing credential should be prefer in this way.
	confedentialConfigFile := os.Getenv("SYNC_REPO_CREDENTIAL_PATH")
	if confedentialConfigFile == "" {
		// Basically we can't put credential on local file.
		// Provided this option for testing locally purpose only.
		confedentialConfigFile = "/opt/pushupdate/config/pushupdate-config.yaml"
	}

	file, err := ioutil.ReadFile(confedentialConfigFile)
	if err != nil {
		fmt.Printf("error reading confedential config file: %v", err)
	}

	mainRepoConfig := new(MainRepoConfig)

	// mapping confidential file into confedential custom struct
	err = yaml.Unmarshal(file, mainRepoConfig)
	if err != nil {
		fmt.Printf("error unmarshalling sync repo config file: %v\n", err)
	}

	// TODO: In future provide option for env, so that user can credential
	// info like repoURL, brnach, etc from externally.
}
