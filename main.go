package main

import (
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"github.com/urfave/negroni"
)

func main() {

	router := httprouter.New()
	router.POST("/resize", wh(resize))
	router.GET("/op/*options", wh(options))

	n := negroni.Classic()
	n.UseHandler(router)
	n.Run(":3000")
}

func options(w http.ResponseWriter, r *http.Request) {
	spew.Dump(context.GetAll(r))
}

func wh(h http.HandlerFunc) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		context.Set(r, "params", p)
		h.ServeHTTP(w, r)
	}
}
