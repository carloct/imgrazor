package main

import (
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"github.com/mholt/binding"
	"github.com/urfave/negroni"
)

func main() {

	router := httprouter.New()
	router.POST("/image", wh(post))
	router.GET("/", wh(test))

	n := negroni.Classic()
	n.UseHandler(router)
	n.Run(":80")
}

func post(w http.ResponseWriter, r *http.Request) {

	postRequest := new(PostRequest)
	errs := binding.Bind(r, postRequest)
	if errs.Handle(w) {
		spew.Dump(errs)
		return
	}

	spew.Dump(postRequest)
}

func test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Test"))
}

func wh(h http.HandlerFunc) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		context.Set(r, "params", p)
		h.ServeHTTP(w, r)
	}
}
