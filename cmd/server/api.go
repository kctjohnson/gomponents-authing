package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func (app *Application) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id != "" {
		idInt, err := strconv.Atoi(id)
		if err != nil {
			w.Write([]byte(fmt.Sprintf("Error: %s", err.Error())))
			return
		}
		err = app.AccountsRepo.Delete(idInt)
		if err != nil {
			w.Write([]byte(fmt.Sprintf("Error: %s", err.Error())))
			return
		}
	}
}
