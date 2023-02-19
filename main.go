package main

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/viveksahu26/syncrepo/config"
)

func main() {
	// Initialize configuration values like PORT
	// DEBUG, LOG, SYNC_REPO_PATH
	config.Init()

	// Initialize confidentials values like tokens, password, etc
	config.ConfidentialInit()

	// The preceding code block instantiates a new Echo server.
	echoServer := echo.New()

	// get the port
	port := config.TempConfigFile.Server.Port

	// errChan := make(chan error)

	// server is started on port 8000
	echoServer.Logger.Fatal(echoServer.Start(":" + strconv.Itoa(port)))

	/*
		Go’s `net/http` package defines a handler function signature as a f
		unction that takes in an “http.ResponseWriter” and a “http.Request”.

		The handler type in Echo is a function that takes in an Echo “Context”,
		and returns an error.
	*/
}
