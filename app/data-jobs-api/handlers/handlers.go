// Package handlers contains the full set of handler functions and routes
// supported by the web api.
package handlers

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/joshakeman/portal-demo/foundation/web"
	// "github.com/ardanlabs/service/business/auth" // Import is removed in final PR
	// "github.com/ardanlabs/service/foundation/web"
	// "github.com/jmoiron/sqlx"
)

// API constructs an http.Handler with all application routes defined.
// func API(build string, shutdown chan os.Signal, log *log.Logger, db *sqlx.DB, a *auth.Auth) http.Handler {
func API(build string, shutdown chan os.Signal, db *sql.DB) http.Handler {

	// Construct the web.App which holds all routes as well as common Middleware.
	// app := web.NewApp(shutdown, mid.Logger(log), mid.Errors(log), mid.Metrics(), mid.Panics(log))
	app := web.NewApp(shutdown)

	// // Register health check endpoint. This route is not authenticated.
	// c := check{
	// 	build: build,
	// 	db:    db,
	// }
	// app.Handle(http.MethodGet, "/v1/health", c.health)

	// // Register user management and authentication endpoints.
	// u := userHandlers{
	// 	db:   db,
	// 	auth: a,
	// }
	// app.Handle(http.MethodGet, "/v1/users", u.list, mid.Authenticate(a), mid.HasRole(auth.RoleAdmin))
	// app.Handle(http.MethodPost, "/v1/users", u.create, mid.Authenticate(a), mid.HasRole(auth.RoleAdmin))
	// app.Handle(http.MethodGet, "/v1/users/:id", u.retrieve, mid.Authenticate(a))
	// app.Handle(http.MethodPut, "/v1/users/:id", u.update, mid.Authenticate(a), mid.HasRole(auth.RoleAdmin))
	// app.Handle(http.MethodDelete, "/v1/users/:id", u.delete, mid.Authenticate(a), mid.HasRole(auth.RoleAdmin))

	// // This route is not authenticated.
	// app.Handle(http.MethodGet, "/v1/users/token", u.token)

	// Register job  endpoints.
	j := jobHandlers{
		db: db,
	}
	app.Handle(http.MethodGet, "/jobs", j.list)
	// app.Handle(http.MethodPost, "/v1/products", p.create, mid.Authenticate(a))
	// app.Handle(http.MethodGet, "/v1/products/:id", p.retrieve, mid.Authenticate(a))
	// app.Handle(http.MethodPut, "/v1/products/:id", p.update, mid.Authenticate(a))
	// app.Handle(http.MethodDelete, "/v1/products/:id", p.delete, mid.Authenticate(a))

	return app
}
