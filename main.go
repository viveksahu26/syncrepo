package main

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/viveksahu26/syncrepo/config"
	"github.com/viveksahu26/syncrepo/pkg/git"
	"github.com/viveksahu26/syncrepo/pkg/http"
	"github.com/viveksahu26/syncrepo/pkg/syncrepo"
)

func init() {
	config.Init()

	config.ConfidentialInit()
}

func main() {
	echoServer := echo.New()

	// Retrieve the Server port from Server Configuration
	port := config.GetServerConfig().Port

	// errChan := make(chan error)

	// initialize git service
	gitService := git.InitGitServices()

	// init sync repo service
	syncRepoService := syncrepo.InitSyncRepoServices(gitService)

	// init handlers
	http.InitSyncRepoHttpHandler(echoServer, syncRepoService)

	errChan := make(chan error)

	go runner(echoServer, port, errChan)
}

func runner(e *echo.Echo, port int, errChan chan error) {
	err := e.Start(":" + strconv.Itoa(port))
	// return errors.Wrap(err, "error listening on port err "+strconv.Itoa(port))
	if err != nil {
		errChan <- err
	}
}
