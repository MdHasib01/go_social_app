package main

import "net/http"

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Connection-Type", "application/json")
	w.Write([]byte(`{"status":"ok"}`))

}
