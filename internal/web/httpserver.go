package web

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// ServeHttp will start http server with graceful shutdown
func ServeHttp(addr, serviceName string, router http.Handler) {
	srv := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	go func() {
		logrus.Infof("[%s] http listen: %v", serviceName, srv.Addr)

		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			logrus.Error("server listen err: ", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logrus.Warn("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logrus.Fatal("server forced to shutdown: ", err)
	}

	logrus.Warn("server exited")
}
