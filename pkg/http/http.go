package http

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/viveksahu26/syncrepo/types"
)

type handlerEvents struct {
	service types.SyncRepoService
}

func InitSyncRepoHttpHandler(e *echo.Echo, service types.SyncRepoService) {
	h := &handlerEvents{service: service}
	e.POST("/push-webhook", h.handleGitlabWebhook)
}

// handleGitlabWebhook handles gitlab webhook when push.
func (h handlerEvents) handleGitlabWebhook(ctx echo.Context) error {
	fmt.Println("Inside handleGitlabWebhook")
	event := new(types.PushEventGitlab)
	if err := ctx.Bind(event); err != nil {
		return err
	}

	err := h.service.SyncRepo(ctx.Request().Context(), event)
	if err != nil {
		return err
	}

	return ctx.NoContent(http.StatusOK)
}
