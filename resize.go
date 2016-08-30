package main

import (
	"mime/multipart"
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/mholt/binding"
)

type ResizeReq struct {
	Image *multipart.FileHeader `json:"image"`
}

func (r *ResizeReq) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&r.Image: binding.Field{
			Form:     "image",
			Required: true,
		},
	}
}

func (r *ResizeReq) Validate(req *http.Request, errs binding.Errors) binding.Errors {
	if r.Image == nil {
		errs = append(errs, binding.Error{
			FieldNames:     []string{"image"},
			Classification: "Invalid Input",
			Message:        "Error: no image found in the request",
		})
	}
	return errs
}

func resize(w http.ResponseWriter, r *http.Request) {
	request := new(ResizeReq)
	errs := binding.Bind(r, request)
	if errs.Handle(w) {
		return
	}

	spew.Dump(request)
}
