package router

import (
	"notification-microservice/cmd/api/default"

	"github.com/julienschmidt/httprouter"
	"notification-microservice/pkg/application"
)

func Get(app *application.Application) *httprouter.Router {
	mux := httprouter.New()

	// index
	//mux.GET("/", _default.Index(app))

	// post
	mux.POST("/whatsapp", _default.Store(app))
	//mux.GET("/users/:id", getuser.runIndex(app))
	return mux
}
