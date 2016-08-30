package main

import (
	"bytes"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"time"

	"gopkg.in/h2non/bimg.v0"

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
	resizeReq := new(ResizeReq)
	errs := binding.Bind(r, resizeReq)
	if errs.Handle(w) {
		return
	}

	name := resizeReq.Image.Filename
	file, err := resizeReq.Image.Open()
	defer file.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	buf, err = ioutil.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	image, err := bimg.NewImage(buf).Resize(300, 300)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.ServeContent(w, r, name, time.Now(), bytes.NewReader(image))
}
