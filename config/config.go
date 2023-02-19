package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"gopkg.in/yaml.v3"
)

type config struct {
	Server Server
	Log    log
}

type Server struct {
	Port    int
	Debug   bool
	Timeout int
}

type log struct {
	Level string
}

type SyncRepoConfig struct {
	RepoURL             string
	UserName            string
	Token               string
	Branch              string
	FollowersRepoConfig map[string]FollowersRepoConfig
}

type FollowersRepoConfig struct {
	RepoURL    string
	UserName   string
	Token      string
	RemoteName string
	Branch     string
}

var TempConfigFile = new(config)

// func GetServerConfig() *Server {
// 	return &tempConfigFile
// }

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
	err = yaml.Unmarshal(file, TempConfigFile)
	if err != nil {
		fmt.Printf("error unmarshalling config file: %v\n", err)
	}
	// If env variables are gives which means to change the default values.
	if debug := os.Getenv("DEBUG"); debug != "" {
		val, err := strconv.ParseBool(debug)
		if err != nil {
			fmt.Printf("error parsing debug env var: %v\n", err)
		}
		TempConfigFile.Server.Debug = val
	}

	if port := os.Getenv("PORT"); port != "" {
		val, err := strconv.Atoi(port)
		if err != nil {
			fmt.Printf("error parsing port env var: %v\n", err)
		}
		// change the default value of port to value externally provided by user
		TempConfigFile.Server.Port = val
	}
}

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

	syncrepoconfig := new(SyncRepoConfig)

	// mapping confidential file into confedential custom struct
	err = yaml.Unmarshal(file, syncrepoconfig)
	if err != nil {
		fmt.Printf("error unmarshalling sync repo config file: %v\n", err)
	}

	// TODO: In future provide option for env, so that user can credential
	// info like repoURL, brnach, etc from externally.
}
