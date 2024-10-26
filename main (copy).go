// package main

// import (
// 	"database/sql"
// 	"encoding/json"
// 	"fmt"
// 	"html/template"
// 	"net/http"

// 	_ "github.com/go-sql-driver/mysql"
// )

// var DB *sql.DB
// var err error
// var dbUsername, dbPassword string
// var auth bool

// type korisnik struct {
// 	Plastenik int8   `json:"plastenik"`
// 	Biljka    string `json:"biljka"`
// }

// func main() {
// 	DB, err = sql.Open("mysql", "amina:password@tcp(127.0.0.1:3306)/tkm")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer DB.Close()

// 	http.HandleFunc("/", LoginPage)
// 	http.HandleFunc("/login", LoginHandler)
// 	http.HandleFunc("/dashboard", DashboardPage)
// 	http.HandleFunc("/dashboardData", DashboardHandler)

// 	fmt.Println("Server started on http://localhost:8080")
// 	http.ListenAndServe(":8080", nil)
// }

// func LoginPage(w http.ResponseWriter, r *http.Request) {
// 	http.ServeFile(w, r, "index.html")
// }

// func LoginHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == http.MethodPost {
// 		username := r.FormValue("username")
// 		password := r.FormValue("password")
// 		err := DB.QueryRow("SELECT idUser, password FROM user WHERE idUser = ?", username).Scan(&dbUsername, &dbPassword)
// 		if err != nil {
// 			if err == sql.ErrNoRows {
// 				fmt.Fprintf(w, "Invalid credentials. Please try again.")
// 				return
// 			}
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}

// 		if password != dbPassword {
// 			fmt.Fprintf(w, "Invalid credentials. Please try again.")
// 			return
// 		}
// 	}

// 	auth = true
// 	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)

// 	tmpl, err := template.ParseFiles("index.html")
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	tmpl.Execute(w, nil)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	http.ServeFile(w, r, "index.html")
// }

// func fetchUserData(username string) (korisnik, error) {
// 	var userData korisnik
// 	err := DB.QueryRow("SELECT idPlasenik, biljka FROM korisnik WHERE idUser = ?", username).Scan(&userData)
// 	if err != nil {
// 		return korisnik{}, err
// 	}
// 	return userData, nil
// }

// func DashboardHandler(w http.ResponseWriter, r *http.Request) {
// 	username := dbUsername
// 	userData, err := fetchUserData(username)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	jsonData, err := json.Marshal(userData)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")

// 	_, err = w.Write(jsonData)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// }

// func DashboardPage(w http.ResponseWriter, r *http.Request) {
// 	if !auth {
// 		http.Redirect(w, r, "/", http.StatusSeeOther)
// 		return
// 	}
// 	http.ServeFile(w, r, "dashboard/dashboard.html")
// }
