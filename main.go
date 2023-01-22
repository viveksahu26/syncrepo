package main

import "github.com/viveksahu26/syncrepo/config"

func main() {
	// Initialize configuration values like PORT
	// DEBUG, LOG, SYNC_REPO_PATH
	config.Init()

	// Initialize confidentials values like tokens, password, etc
	config.ConfidentialInit()
}
