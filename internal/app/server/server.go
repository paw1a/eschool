package server

import (
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	v1 "github.com/paw1a/eschool-delivery/http/v1"
	"github.com/paw1a/eschool/internal/app/config"
	log "github.com/sirupsen/logrus"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"go.uber.org/fx"
	"net/http"
	"time"
)

func NewGinRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.Use(cors.Default())
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
				log.Infof("server started on port %s", params.Cfg.Web.Port)
				log.Fatal(server.ListenAndServe())
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
	return server
}
