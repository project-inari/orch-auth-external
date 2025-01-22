// Package di provides dependency injection for the server
package di

import (
	"context"
	"log"
	"log/slog"
	"net/http"

	"github.com/getsentry/sentry-go"
	"github.com/labstack/echo/v4"

	"github.com/project-inari/orch-auth-external/config"
	"github.com/project-inari/orch-auth-external/handler"
	"github.com/project-inari/orch-auth-external/pkg/httpclient"
	"github.com/project-inari/orch-auth-external/repository"
	"github.com/project-inari/orch-auth-external/service"
)

// New injects the dependencies for the server
func New(c *config.Config) {
	ctx := context.Background()

	// Sentry initialization
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:                c.SentryConfig.SentryDSN,
		Debug:              false,
		EnableTracing:      true,
		TracesSampleRate:   1.0,
		ProfilesSampleRate: 1.0,
	}); err != nil {
		slog.Error("error - [main.New] sentry initialization failed", slog.Any("error", err))
	}

	// Echo server initialization
	e := echo.New()
	setupServer(ctx, e, c)

	// HTTP Client initialization
	httpClientCoreAuth := httpclient.NewHTTPClient(httpclient.Options{
		MaxConns:                 c.APICoreAuthConfig.MaxConns,
		MaxRetry:                 c.APICoreAuthConfig.MaxRetry,
		Timeout:                  c.APICoreAuthConfig.Timeout,
		InsecureSkipVerify:       c.APICoreAuthConfig.InsecureSkipVerify,
		MaxTransactionsPerSecond: c.APICoreAuthConfig.MaxTransactionsPerSecond,
	})

	httpClientCoreUser := httpclient.NewHTTPClient(httpclient.Options{
		MaxConns:                 c.APICoreUserConfig.MaxConns,
		MaxRetry:                 c.APICoreUserConfig.MaxRetry,
		Timeout:                  c.APICoreUserConfig.Timeout,
		InsecureSkipVerify:       c.APICoreUserConfig.InsecureSkipVerify,
		MaxTransactionsPerSecond: c.APICoreUserConfig.MaxTransactionsPerSecond,
	})

	// Repository initialization
	coreAuthAPIRepo := repository.NewCoreAuthAPIRepository(repository.CoreAuthAPIRepositoryConfig{
		BaseURL:        c.APICoreAuthConfig.BaseURL,
		SignupPath:     c.APICoreAuthConfig.SignupPath,
		DeleteUserPath: c.APICoreAuthConfig.DeleteUserPath,
	}, repository.CoreAuthAPIRepositoryDependencies{
		Client: httpClientCoreAuth,
	})

	coreUserAPIRepo := repository.NewCoreUserAPIRepository(repository.CoreUserAPIRepositoryConfig{
		BaseURL:    c.APICoreUserConfig.BaseURL,
		SignupPath: c.APICoreUserConfig.SignupPath,
	}, repository.CoreUserAPIRepositoryDependencies{
		Client: httpClientCoreUser,
	})

	// Service initialization
	service := service.New(service.Dependencies{
		CoreAuthAPIRepository: coreAuthAPIRepo,
		CoreUserAPIRepository: coreUserAPIRepo,
	})

	// Handler initialization
	handler.New(e, handler.Dependencies{
		Service: service,
	})

	// HTTP Listening
	if err := e.Start(":" + c.AppConfig.Port); err != nil && err != http.ErrServerClosed {
		log.Panicf("error - [main.New] unable to start server: %v", err)
	}
}
