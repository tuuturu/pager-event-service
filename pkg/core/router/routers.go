/*
 * Paged
 *
 * Handles CRUD operations for events
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package router

import (
	"fmt"
	"net/http"

	"github.com/oslokommune/go-oidc-middleware/pkg/v1/middleware"

	"github.com/tuuturu/paged/pkg/core"
	"github.com/tuuturu/paged/pkg/core/handlers"

	"github.com/tuuturu/paged/pkg/storage/upper"

	"github.com/gin-gonic/gin"
)

// NewRouter returns a new router.
func New(cfg *core.Config) *gin.Engine {
	router := gin.Default()

	var storage core.StorageClient

	switch {
	default:
		storage = upper.NewUpperClient(
			cfg.Database.URI,
			cfg.Database.DatabaseName,
			cfg.Database.Username,
			cfg.Database.Password,
		)
	}

	err := storage.Open()
	if err != nil {
		panic(fmt.Sprintf("error opening storage connection: %v", err))
	}

	router.Use(middleware.NewGinAuthenticationMiddleware(*cfg.DiscoveryURL))

	for _, route := range routes {
		switch route.Method {
		case http.MethodGet:
			router.GET(route.Pattern, route.HandlerFunc(storage))
		case http.MethodPost:
			router.POST(route.Pattern, route.HandlerFunc(storage))
		case http.MethodPut:
			router.PUT(route.Pattern, route.HandlerFunc(storage))
		case http.MethodPatch:
			router.PATCH(route.Pattern, route.HandlerFunc(storage))
		case http.MethodDelete:
			router.DELETE(route.Pattern, route.HandlerFunc(storage))
		}
	}

	return router
}

// Index is the index handler.
func Index(_ core.StorageClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	}
}

var routes = core.Routes{
	{
		Name:        "Index",
		Method:      http.MethodGet,
		Pattern:     "/",
		HandlerFunc: Index,
	},

	{
		Name:        "AddEvent",
		Method:      http.MethodPost,
		Pattern:     "/events",
		HandlerFunc: handlers.AddEvent,
	},

	{
		Name:        "GetEvents",
		Method:      http.MethodGet,
		Pattern:     "/events",
		HandlerFunc: handlers.GetEvents,
	},

	{
		Name:        "DeleteEvent",
		Method:      http.MethodDelete,
		Pattern:     "/events/:id",
		HandlerFunc: handlers.DeleteEvent,
	},

	{
		Name:        "GetEvent",
		Method:      http.MethodGet,
		Pattern:     "/events/:id",
		HandlerFunc: handlers.GetEvent,
	},

	{
		Name:        "UpdateEvent",
		Method:      http.MethodPatch,
		Pattern:     "/events/:id",
		HandlerFunc: handlers.UpdateEvent,
	},
}
