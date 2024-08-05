package rest

import (
	"net/http"

	"github.com/naysw/permission/api/rest/handlers"
	"github.com/naysw/permission/internal/core"
)

func registerRoutes(mux *http.ServeMux, app *core.App) {
	policyHandler := handlers.NewPolicyHandler(app.PolicyUsecase())
	authzHandler := handlers.NewAuthzHandler(app.RoleUsecase(), app.PolicyUsecase())

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		w.Write([]byte("Hello World!"))
	})

	mux.HandleFunc("GET /policies", policyHandler.GetList)
	mux.HandleFunc("GET /policies/:id", policyHandler.GetByID)
	mux.HandleFunc("POST /policies", policyHandler.Create)
	mux.HandleFunc("PUT /policies/:id", policyHandler.Update)
	mux.HandleFunc("DELETE /policies/:id", policyHandler.Delete)
	mux.HandleFunc("POST /policies/attach", policyHandler.Attach)

	mux.HandleFunc("POST /is_authorized", authzHandler.IsAuthorized)
}
