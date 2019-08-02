package main

import "net/http"

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := tplHome.Execute(w, nil); err != nil {
		panic(err)
	}
}

func info(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := tplInfo.Execute(w, nil); err != nil {
		panic(err)
	}
}

func registera(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := tplRegistera.Execute(w, nil); err != nil {
		panic(err)
	}
}

func registerf(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := tplRegisterf.Execute(w, nil); err != nil {
		panic(err)
	}
}

func create(w http.ResponseWriter, r *http.Request) {
	db := dbConn()

	ap := Applicant{}
	ap.Name = r.FormValue("name")
	ap.Mobile = r.FormValue("mobile")
	ap.Position = r.FormValue("position")
	ap.Role = r.FormValue("role")

	db.Create(&ap)
	db.Close()

	http.Redirect(w, r, "/success", http.StatusSeeOther)
	return
}

func apprentice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := tplApprentice.Execute(w, nil); err != nil {
		panic(err)
	}
}

func freelance(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := tplFreelance.Execute(w, nil); err != nil {
		panic(err)
	}
}

func success(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := tplSuccess.Execute(w, nil); err != nil {
		panic(err)
	}
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/public/images/favicon.ico")
}