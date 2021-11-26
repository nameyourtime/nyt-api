package main

import (
	"net/http"
)

func (app *Application) Home(w http.ResponseWriter, r *http.Request) {

	//if err != nil {
	//	app.ServerError(w, err)
	//	return
	//}

	app.RenderHTML(w, r, "home.page.tmpl", &HTMLData{})

}
