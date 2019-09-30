package app

import (
	"net/http"

	"gwiapi/app/handlers"
	"gwiapi/app/middlewares"
)

func (a *App) setRouters() {

	// Health Check
	a.Get("/", middlewares.JSON(handlers.Ping, a.DB))

	// Auth Routes
	a.Post("/auth/signup", middlewares.JSON(handlers.CreateUser, a.DB))
	a.Post("/auth/login", middlewares.JSON(handlers.Login, a.DB))

	// Users routes
	a.Get("/users", middlewares.Auth(handlers.GetUsers, a.DB))
	a.Get("/users/{id}", middlewares.Auth(handlers.GetUser, a.DB))
	a.Put("/users/{id}", middlewares.Auth(handlers.UpdateUser, a.DB))
	a.Delete("/users/{id}", middlewares.Auth(handlers.DeleteUser, a.DB))

	// Assets Routes
	a.Get("/assets", middlewares.Auth(handlers.GetAssets, a.DB))
	a.Get("/assets/{id:[0-9]+}", middlewares.Auth(handlers.GetAsset, a.DB))
	a.Patch("/assets/{id:[0-9]+}", middlewares.Auth(handlers.UpdateAsset, a.DB))
	a.Delete("/assets/{id:[0-9]+}", middlewares.Auth(handlers.DeleteAsset, a.DB))
	a.Post("/assets/charts", middlewares.Auth(handlers.CreateAssetChart, a.DB))
	a.Post("/assets/insights", middlewares.Auth(handlers.CreateAssetInsight, a.DB))
	a.Post("/assets/audiences", middlewares.Auth(handlers.CreateAssetAudience, a.DB))

	// Participants
	a.Get("/participants", middlewares.Auth(handlers.GetParticipants, a.DB))
	a.Post("/participants", middlewares.Auth(handlers.AddParticipant, a.DB))
	a.Get("/countries", middlewares.JSON(handlers.GetCountries, a.DB))

}

// Get wrap HandleFunc
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Post Get wrap HandleFunc
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Patch wrap HandleFunc
func (a *App) Patch(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PATCH")
}

// Put wrap HandleFunc
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Delete wrap HandleFunc
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}
