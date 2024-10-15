package server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/paw1a/eschool/docs"
	v1 "github.com/paw1a/eschool/internal/adapter/delivery/http/v1"
	"github.com/paw1a/eschool/internal/app/config"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func NewGinRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	return router
}

type ServerParams struct {
	fx.In
	Cfg     *config.Config
	Handler *v1.Handler
	Router  *gin.Engine
	Logger  *zap.Logger
}

func NewServer(lc fx.Lifecycle, params ServerParams) *http.Server {
	server := &http.Server{
		Handler:      params.Router,
		Addr:         fmt.Sprintf(":%s", params.Cfg.Web.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				params.Logger.Info(fmt.Sprintf("server started on port %s", params.Cfg.Web.Port))
				params.Logger.Fatal("server shutdown", zap.Error(server.ListenAndServe()))
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
	return server
}
