package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"html/template"
	"log"
	"net/http"
	"os"
)

var (
	tplHome *template.Template
	tplInfo *template.Template
	tplRegisterf *template.Template
	tplRegistera *template.Template
	tplSuccess *template.Template
	tplJoinUs *template.Template
)

type Applicant struct{
	gorm.Model
	Name string
	Mobile string
	Position string
	Role string
}

func dbConn() (db *gorm.DB) {
	dbhost     := os.Getenv("DB_HOST")
	dbport     := os.Getenv("DB_PORT")
	dbuser     := os.Getenv("DB_USER")
	dbpassword := os.Getenv("DB_PASSWORD")
	dbname     := os.Getenv("DB_NAME")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbhost, dbport, dbuser, dbpassword, dbname)

	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	return db
}

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

func registerf(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := tplRegisterf.Execute(w, nil); err != nil {
		panic(err)
	}
}

func registera(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := tplRegistera.Execute(w, nil); err != nil {
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

func joinUs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := tplJoinUs.Execute(w, nil); err != nil {
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

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	db := dbConn()

	db.AutoMigrate(&Applicant{})

	tplHome = template.Must(template.ParseFiles("views/layouts/main.gohtml", "views/pages/home.gohtml"))
	tplInfo = template.Must(template.ParseFiles("views/layouts/main.gohtml", "views/pages/freelance.gohtml"))
	tplRegisterf = template.Must(template.ParseFiles("views/layouts/main.gohtml", "views/pages/registerf.gohtml"))
	tplRegistera = template.Must(template.ParseFiles("views/layouts/main.gohtml", "views/pages/registera.gohtml"))
	tplJoinUs = template.Must(template.ParseFiles("views/layouts/main.gohtml", "views/pages/apprentice.gohtml"))
	tplSuccess = template.Must(template.ParseFiles("views/layouts/main.gohtml", "views/pages/success.gohtml"))

	r := mux.NewRouter()
	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/info", info).Methods("GET")
	r.HandleFunc("/freelance", freelance).Methods("GET")
	r.HandleFunc("/apprentice", apprentice).Methods("GET")
	r.HandleFunc("/apprentice", freelanceCreate).Methods("POST")
	r.HandleFunc("/register", apprenticeCreate).Methods("POST")
	r.HandleFunc("/success", success).Methods("GET")

	// Styles
	assetHandler := http.FileServer(http.Dir("./dist/"))
	assetHandler = http.StripPrefix("/dist/", assetHandler)
	r.PathPrefix("/dist/").Handler(assetHandler)

	// JS
	jsHandler := http.FileServer(http.Dir("./dist/"))
	jsHandler = http.StripPrefix("/dist/", jsHandler)
	r.PathPrefix("/dist/").Handler(jsHandler)

	//Images
	imageHandler := http.FileServer(http.Dir("./dist/images/"))
	r.PathPrefix("/dist/images/").Handler(http.StripPrefix("/dist/images/", imageHandler))


	http.HandleFunc("/favicon.ico", faviconHandler)

	log.Printf("Starting server on %s", port)

	http.ListenAndServe(":" + port, r)
}