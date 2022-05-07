package application

import (
	"fmt"
	"net/http"
)

func (app *Application) getMovies(w http.ResponseWriter, r *http.Request, context CustomContext) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("MOVIES"))
}

func (app *Application) getMovieById(w http.ResponseWriter, r *http.Request, context CustomContext) {
	params := r.URL.Query()

	app.logger.Print("koji su paramsi")
	app.logger.Printf(fmt.Sprintf("%v", params))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("MOVIES WITH ID"))
}
