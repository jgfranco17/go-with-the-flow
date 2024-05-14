package v0

import (
	"github.com/gin-gonic/gin"
	"github.com/jgfranco17/go-with-the-flow/service/pkg/handlers"
	error_handling "github.com/jgfranco17/go-with-the-flow/service/pkg/router/error_handling"
)

// Adds v0 routes to the router.
func SetRoutes(route *gin.Engine) {
	v0 := route.Group("/v0")
	{
		p := v0.Group("pipelines")
		{
			p.GET("verify/:user", error_handling.WithErrorHandling(handlers.VerificationHandler()))
		}
	}
}
