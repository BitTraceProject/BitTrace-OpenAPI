/*
 * BitTrace
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package server

import (
	"net/http"
	"sync"

	"github.com/BitTraceProject/BitTrace-Types/pkg/config"
	"github.com/BitTraceProject/BitTrace-Types/pkg/database"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type (
	OpenAPIServer struct {
		*gin.Engine

		db *gorm.DB // base DB
	}
	// Route is the information for every URI.
	Route struct {
		// Name is the name of this Route.
		Name string
		// Method is the string for the HTTP method. ex) GET, POST etc..
		Method string
		// Pattern is the pattern of the URI.
		Pattern string
		// HandlerFunc is the handler function of this route.
		HandlerFunc gin.HandlerFunc
	}

	// Routes is the list of the generated Route.
	Routes []Route
)

var (
	oas      *OpenAPIServer
	once     sync.Once
	routesV1 = Routes{
		{
			"OpenapiV1Index",
			http.MethodGet,
			"/",
			OpenapiV1Index,
		},

		{
			"OpenapiV1SnapshotGet",
			http.MethodGet,
			"/snapshot",
			OpenapiV1SnapshotGet,
		},

		{
			"OpenapiV1RevisionGet",
			http.MethodGet,
			"/revision",
			OpenapiV1RevisionGet,
		},

		{
			"OpenapiV1BestStateGet",
			http.MethodGet,
			"/best_state",
			OpenapiV1BestStateGet,
		},

		{
			"OpenapiV1EventOrphanGet",
			http.MethodGet,
			"/event/orphan",
			OpenapiV1EventOrphanGet,
		},

		{
			"OpenapiV1TimelineGet",
			http.MethodGet,
			"/timeline",
			OpenapiV1TimelineGet,
		},

		{
			"OpenapiV1PeerGet",
			http.MethodGet,
			"/peer",
			OpenapiV1PeerGet,
		},

		{
			"OpenapiV1AuthRegisterPost",
			http.MethodPost,
			"/auth/register",
			OpenapiV1AuthRegisterPost,
		},

		{
			"OpenapiV1AuthTokenPost",
			http.MethodPost,
			"/auth/token",
			OpenapiV1AuthTokenPost,
		},

		// 临时的，目前是 get
		{
			"OpenapiV1AuthTokenGet",
			http.MethodGet,
			"/auth/token",
			OpenapiV1AuthTokenGet,
		},
	}
)

// InitOpenAPIServer init and returns the openapi server.
func InitOpenAPIServer(dbConf config.DatabaseConfig) *OpenAPIServer {
	once.Do(func() {
		if oas == nil {
			db, err := database.NewDBInstance(dbConf)
			if err != nil {
				panic(err)
			}
			oas = &OpenAPIServer{
				Engine: gin.Default(),
				db:     db,
			}
			oas.register()
		}
	})
	return oas
}

func RunOpenAPIServer(addr string) error {
	return oas.Run(addr)
}

func (s *OpenAPIServer) register() {
	r := s.Group("/openapi/v1")
	for _, route := range routesV1 {
		switch route.Method {
		case http.MethodGet:
			r.GET(route.Pattern, route.HandlerFunc)
		case http.MethodPost:
			r.POST(route.Pattern, route.HandlerFunc)
		case http.MethodPut:
			r.PUT(route.Pattern, route.HandlerFunc)
		case http.MethodDelete:
			r.DELETE(route.Pattern, route.HandlerFunc)
		}
	}
}

func (s *OpenAPIServer) Raw(sql string) (*gorm.DB, error) {
	var (
		db  *gorm.DB
		err error
	)
	db, err = database.NewDBInstanceCopy(s.db)
	if err != nil {
		return nil, err
	}
	db = db.Raw(sql)
	if db.Error != nil {
		return nil, err
	}
	return db, nil
}
