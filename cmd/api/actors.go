package main

import (
	"net/http"
)

func (app *application) GetActors(w http.ResponseWriter, r *http.Request) {
	actors, err := app.models.ActorModel.GetActors()
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, actors, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}
