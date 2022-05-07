package application

import (
	"encoding/json"
	"net/http"
)

func (app *Application) statusHandler(w http.ResponseWriter, r *http.Request, context CustomContext) {
	currentStatus := AppStatus{
		Status:      "available",
		Environment: app.config.Env,
		Version:     version,
	}

	js, err := json.MarshalIndent(currentStatus, "", "\t")
	if err != nil {
		app.logger.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}
