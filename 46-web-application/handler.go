package main

import (
	"net/http"
)

const (
	loggedInUserKey = "logged_in_user_id"
)

func (app *Application) home(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "index.html", nil)
}

func (app *Application) about(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "about.html", nil)
}

func (app *Application) login(w http.ResponseWriter, r *http.Request) {

	if app.isAuthenticated(r) {
		http.Redirect(w, r, "/submit_post", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodPost {

		if err := r.ParseForm(); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		form := NewForm(r.PostForm)

		form.Required("email", "password").
			MaxLength("email", 255).MaxLength("password", 255).MinLength("password", 3).MinLength("email", 3).
			MatchesPattern("email", EmailRX)

		if !form.Valid() {
			form.Errors.Add("generic", "the data you send is not valid")
			app.errorLog.Printf("Form validation errors: %v", form.Errors)
			app.render(w, r, "login.html", &templateData{
				Form: form,
			})
			return
		}

		email := r.FormValue("email")
		password := r.FormValue("password")

		id, err := app.userRepo.Authenticate(email, password)

		if err != nil {
			form.Errors.Add("generic", err.Error())
			app.render(w, r, "login.html", &templateData{
				Form: form,
			})
			return
		}

		//login in
		app.session.Put(r, loggedInUserKey, email)
		app.session.Put(r, "flash", "Successfully logged in!")

		app.infoLog.Println("Logged in user ID:", id)

		http.Redirect(w, r, "/submit_post", http.StatusSeeOther)
		return

	}

	app.render(w, r, "login.html", &templateData{
		Form: NewForm(r.PostForm),
	})
}

func (app *Application) register(w http.ResponseWriter, r *http.Request) {

	if app.isAuthenticated(r) {
		http.Redirect(w, r, "/submit_post", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodPost {

		if err := r.ParseForm(); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		form := NewForm(r.PostForm)

		form.Required("email", "password", "name").
			MaxLength("email", 255).MaxLength("password", 255).MinLength("password", 3).MinLength("email", 3).
			MatchesPattern("email", EmailRX)

		if !form.Valid() {
			form.Errors.Add("generic", "the data you send is not valid")
			app.errorLog.Printf("Form validation errors: %v", form.Errors)
			app.render(w, r, "register.html", &templateData{
				Form: form,
			})
			return
		}

		//design pattern
		//fbuildig, stratgey, repositort, observer
		//mvc,model view controller, extend to server layer

		email := r.FormValue("email")
		password := r.FormValue("password")
		name := r.FormValue("name")
		avatar := r.FormValue("avatar")
		bio := "im a software engineer"

		_, err := app.userRepo.CreateUserWithProfile(name, email, password, bio, avatar)

		if err != nil {
			form.Errors.Add("generic", err.Error())
			app.render(w, r, "register.html", &templateData{
				Form: form,
			})
			return
		}

		app.session.Put(r, "flash", "Registration successful. Please log in.")

		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}

	app.render(w, r, "register.html", &templateData{
		Form: NewForm(r.PostForm),
	})

}

func (app *Application) logout(w http.ResponseWriter, r *http.Request) {
	app.session.Remove(r, loggedInUserKey)
	app.session.Put(r, "flash", "You have been logged out successfully.")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *Application) contact(w http.ResponseWriter, r *http.Request) {

	app.render(w, r, "contact.html", nil)
}

func (app *Application) submit_post(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "submit_post.html", nil)
}
