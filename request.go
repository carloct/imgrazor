package main

import (
	"mime/multipart"
	"net/http"

	"github.com/mholt/binding"
)

type PostRequest struct {
	Image  *multipart.FileHeader
	Height int
	Width  int
}

func (r *PostRequest) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&r.Image:  "image",
		&r.Height: "h",
		&r.Width:  "w",
	}
}

func (r *PostRequest) Validate(req *http.Request, errs binding.Errors) binding.Errors {
	if r.Image == nil {
		errs = append(errs, binding.Error{
			FieldNames:     []string{"image"},
			Classification: "Invalid Input",
			Message:        "Error: no image found in the request",
		})
	}
	return errs
}
