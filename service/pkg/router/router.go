package router

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jgfranco17/go-with-the-flow/core/pkg/context_settings"
	"github.com/jgfranco17/go-with-the-flow/core/pkg/logger"
	"github.com/jgfranco17/go-with-the-flow/service/pkg/env"
	"github.com/jgfranco17/go-with-the-flow/service/pkg/router/headers"
	system "github.com/jgfranco17/go-with-the-flow/service/pkg/router/system"
	v0 "github.com/jgfranco17/go-with-the-flow/service/pkg/router/v0"
)

// Add the fields we want to expose in the logger to the request context
func addLoggerFields() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !env.IsLocalEnvironment() {
			requestID := uuid.NewString()
			environment := os.Getenv(env.ENV_KEY_ENVIRONMENT)
			version := os.Getenv(env.ENV_KEY_VERSION)

			c.Set(string(context_settings.RequestId), requestID)
			c.Set(string(context_settings.Environment), environment)
			c.Set(string(context_settings.Version), version)

			originInfo, err := headers.CreateOriginInfoHeader(c)

			if err == nil && originInfo.Origin != "" {
				c.Set(string(context_settings.Origin), fmt.Sprintf("%s@%s", originInfo.Origin, originInfo.Version))
			}
		}
		c.Next()
	}
}

// Log the start and completion of a request
func logRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		log := logger.FromContext(c)

		origin := c.Request.Header.Get("Origin")
		log.Infof("Request Started: [%s] %s from %s", c.Request.Method, c.Request.URL, origin)
		c.Next()
		log.Infof("Request Completed: [%s] %s", c.Request.Method, c.Request.URL)
	}
}

// Configure the router adding routes and middlewares
func GetRouter() *gin.Engine {
	router := gin.Default()
	router.Use(addLoggerFields())
	router.Use(logRequest())
	router.Use(GetCors())
	system.SetSystemRoutes(router)
	v0.SetRoutes(router)

	return router
}
