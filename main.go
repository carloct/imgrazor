package main

import (
	"net/http"

	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"github.com/urfave/negroni"
)

func main() {

	router := httprouter.New()
	router.POST("/resize", wh(resize))

	n := negroni.Classic()
	n.UseHandler(router)
	n.Run(":3000")
}

func wh(h http.HandlerFunc) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		context.Set(r, "params", p)
		h.ServeHTTP(w, r)
	}
}
