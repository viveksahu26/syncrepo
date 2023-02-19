package main

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/viveksahu26/syncrepo/config"
	"github.com/viveksahu26/syncrepo/pkg/git"
	"github.com/viveksahu26/syncrepo/pkg/http"
	"github.com/viveksahu26/syncrepo/pkg/syncrepo"
)

func main() {
	// Initialize configuration values like PORT
	// DEBUG, LOG, SYNC_REPO_PATH
	config.Init()

	// initialize log
	//

	// Initialize confidentials values like tokens, password, etc
	config.ConfidentialInit()

	// The preceding code block instantiates a new Echo server.
	echoServer := echo.New()

	// get the port
	port := config.GetServerConfig().Port

	// errChan := make(chan error)

	// initialize git service
	gitService := git.InitGitServices()

	// init sync repo service
	syncRepoService := syncrepo.InitSyncRepoServices(gitService)

	// init handlers
	http.InitSyncRepoHttpHandler(echoServer, syncRepoService)

	// server is started on port 8000
	// echoServer.Logger.Fatal(echoServer.Start(":" + strconv.Itoa(port)))

	errChan := make(chan error)

	go runner(echoServer, port, errChan)

	/*
		Go’s `net/http` package defines a handler function signature as a f
		unction that takes in an “http.ResponseWriter” and a “http.Request”.

		The handler type in Echo is a function that takes in an Echo “Context”,
		and returns an error.
	*/
}

func runner(e *echo.Echo, port int, errChan chan error) {
	err := e.Start(":" + strconv.Itoa(port))
	// return errors.Wrap(err, "error listening on port err "+strconv.Itoa(port))
	if err != nil {
		errChan <- err
	}
}
