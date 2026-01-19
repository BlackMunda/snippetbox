package main

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime/debug"
)

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
		stack  = debug.Stack()
	)

	app.logger.Error(err.Error(), "method", method, "URI", uri, "stack", stack)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) render(w http.ResponseWriter, r *http.Request, status int, page string, data templateData) {
	ts, ok := app.templateCache[page]
	if !ok {
		err := fmt.Errorf("this template %s doesn't exist : ", page)
		app.serverError(w, r, err)
		return
	}

	buffer := new(bytes.Buffer)

	err := ts.ExecuteTemplate(buffer, "base", data)
	if err != nil {
		app.serverError(w, r, err)
	}

	w.WriteHeader(status)

	buffer.WriteTo(w)
}
