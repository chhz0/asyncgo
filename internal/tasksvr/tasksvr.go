package tasksvr

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/chhz0/asyncgo/internal/tasksvr/biz/dal"
	"github.com/chhz0/asyncgo/internal/tasksvr/biz/router"
	"github.com/chhz0/asyncgo/internal/tasksvr/conf"
	"github.com/gin-gonic/gin"
)

func Run() {
	conf.GetConf()
	dal.Init()

	gracefulGin()
}

func gracefulGin() {
	gin.SetMode(conf.GetConf().Mode)

	r := gin.Default()

	srv := &http.Server{
		Addr:    conf.GetConf().TaskSvr.Port,
		Handler: r,
	}

	router.Register(r)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("tasksvr listen: error", "error", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		slog.Error("tasksvr shutdown: error", "error", err.Error())
	}
	slog.Info("server exiting...")
}
