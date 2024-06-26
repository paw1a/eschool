package web

import (
	"github.com/gin-gonic/gin"
	"github.com/paw1a/eschool/internal/adapter/delivery/web/v1"
	"github.com/paw1a/eschool/internal/app/config"
)

type Handler struct {
	config    *config.Config
	v1Handler *v1.Handler
}

func NewHandler(config *config.Config, v1Handler *v1.Handler, router *gin.Engine) *Handler {
	api := router.Group("/api")
	v1Handler.Init(api)

	return &Handler{
		config:    config,
		v1Handler: v1Handler,
	}
}
