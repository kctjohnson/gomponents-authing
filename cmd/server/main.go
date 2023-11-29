package main

import (
	"authing/internal/db"
	"authing/internal/repositories"
	"authing/internal/web"
	"authing/internal/web/pages"
	"authing/internal/web/pages/home"
	"authing/internal/web/pages/user"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

type Application struct {
	SessionManager *scs.SessionManager
	AccountsRepo   *repositories.Accounts
}

func main() {
	sessionManager := scs.New()
	sessionManager.Lifetime = 24 * time.Hour

	dbCon := db.New("test.db")
	err := dbCon.RunMigrations()
	if err != nil {
		db.Teardown("test.db")
		panic(err)
	}

	app := Application{
		SessionManager: sessionManager,
		AccountsRepo:   repositories.NewAccountsRepository(dbCon.DB),
	}

	mux := http.NewServeMux()

	// Unprotected routes
	mux.Handle("/", app.createHandler(home.IndexPage()))
	mux.Handle("/contact", app.createHandler(home.ContactPage()))
	mux.Handle("/about", app.createHandler(home.AboutPage()))
	mux.Handle("/users", app.createHandler(home.UsersPage(app.AccountsRepo)))
	mux.Handle("/invalid", app.createHandler(pages.InvalidPage()))
	mux.Handle("/login", app.createHandler(home.LoginPage()))
	mux.Handle("/register", app.createHandler(home.RegisterPage()))

	// Protected routes
	mux.Handle("/user/adminpanel", app.createProtectedHandler(user.AdminPanelPage()))
	mux.Handle("/user/users", app.createProtectedHandler(user.UsersPage(app.AccountsRepo)))

	// API routes
	mux.HandleFunc("/auth/register", app.Register)
	mux.HandleFunc("/auth/login", app.Login)
	mux.HandleFunc("/auth/logout", app.Logout)

	// Start the server, wrapping the mux handler in the session manager
	fmt.Printf("Starting server on 8081\n")
	if err := http.ListenAndServe("localhost:8081", app.SessionManager.LoadAndSave(mux)); err != nil &&
		!errors.Is(err, http.ErrServerClosed) {
		log.Println("Error:", err)
	}
}

// Makes it so that calls to this route validate the auth token FIRST before passing through the page
func (app *Application) createProtectedHandler(
	title string,
	bodyFunc web.BodyFunc,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authToken := app.SessionManager.Get(r.Context(), "auth_token")

		if authToken == nil {
			http.Redirect(w, r, "/invalid", http.StatusSeeOther)
			return
		}

		// Rendering a Node is as simple as calling Render and passing an io.Writer
		_ = web.Page(true, title, r.URL.Path, bodyFunc()).Render(w)
	}
}

func (app *Application) createHandler(title string, bodyFunc web.BodyFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authToken := app.SessionManager.Get(r.Context(), "auth_token")

		// Rendering a Node is as simple as calling Render and passing an io.Writer
		_ = web.Page(authToken != nil, title, r.URL.Path, bodyFunc()).Render(w)
	}
}
