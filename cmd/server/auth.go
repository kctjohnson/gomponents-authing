package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func (app *Application) Register(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Error: %s", err.Error())))
		return
	}

	user := r.FormValue("user")
	password := r.FormValue("password")

	// Check to see if the account already exists
	acc, err := app.AccountsRepo.Get(user)
	if err != nil && err != sql.ErrNoRows {
		w.Write([]byte(fmt.Sprintf("Error: %s", err.Error())))
		return
	}

	if acc != nil {
		w.Write([]byte(fmt.Sprintf("Account with username \"%s\" already exists", user)))
		return
	}

	// Generate the hash and add the user to the accounts table
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Error: %s", err.Error())))
		return
	}

	_, err = app.AccountsRepo.Add(user, string(hash))
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Error: %s", err.Error())))
		return
	}

	// Authorize the user with a session auth token, then redirect them to the admin panel
	app.SessionManager.Put(r.Context(), "auth_token", "valid_token")
	http.Redirect(w, r, "/user/adminpanel", http.StatusSeeOther)
}

func (app *Application) Login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Error: %s", err.Error())))
		return
	}

	user := r.FormValue("user")
	password := r.FormValue("password")

	// First, make sure that the user exists
	account, err := app.AccountsRepo.Get(user)
	if err != nil {
		if err == sql.ErrNoRows {
			w.Write([]byte(fmt.Sprintf("User \"%s\" Doesn't Exist", user)))
			return
		}
		w.Write([]byte(fmt.Sprintf("Error: %s", err.Error())))
		return
	}

	// Compare the entered password to the hash password using bcrypt
	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Error: %s", err.Error())))
		return
	}

	// We good? Awesome, add the session auth token and redirect them to the admin panel
	app.SessionManager.Put(r.Context(), "auth_token", "valid_token")
	http.Redirect(w, r, "/user/adminpanel", http.StatusSeeOther)
}

func (app *Application) Logout(w http.ResponseWriter, r *http.Request) {
	// Delete all session stored values and redirect the user to the home page
	app.SessionManager.Destroy(r.Context())
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
