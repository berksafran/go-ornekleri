package main

import (
	"fmt"
	"log"
	"net/http"

	helpers "login-register-module/helpers"
)

// Çoklu değişken belirleme best practice'i.
var (
	username        string = ""
	email           string = ""
	password        string = ""
	passwordConfirm string = ""
)

func main() {
	// mux'ı projemize ekliyoruz.
	mux := http.NewServeMux()

	// Değişkenlerimizi belirliyoruz.

	// Register için handler oluşturuyoruz.
	/*
		Dikkat! Burada mux.Handle değil mux.HandleFunc kullandık.
		mux.Handle kullansaydık; struct belirleyip ServeHTTP methodunu
		bu structa eklemeliydik. Örneğini daha önceki bölümlerde bulabilirsiniz.
	*/
	mux.HandleFunc("/register", RegisterHandler)

	// Login için handler oluşturuyoruz.
	mux.HandleFunc("/login", LoginHandler)

	// Sunucumuzu 8080 portunda ayağa kaldırıyoruz.
	fmt.Println("Listening on Port: 8080...")
	http.ListenAndServe(":8080", mux)
}

// RegisterHandler is ...
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Request: '/register' ")
	// Önce client'tan gelen requestte bulunan datayı parse ediyoruz.
	r.ParseForm()

	// Gelen tüm form fieldlarını terminalde yazdıralım.
	for key, value := range r.Form {
		fmt.Printf("%s - %s\n", key, value)
	}

	// queryParams veya Body > "x-www-form-encoded" tarafından gönderilenleri parse eder.
	username = r.FormValue("username")
	email = r.FormValue("email")
	password = r.FormValue("password")
	passwordConfirm = r.FormValue("passwordConfirm")

	/*
	 Tüm kontrolleri yazmaya başlayalım..
	*/

	// Değerlerin dolu mu boş mu geldiğini kontrol edelim.
	usernameCheck := helpers.IsEmpty(username)
	emailCheck := helpers.IsEmpty(email)
	passwordCheck := helpers.IsEmpty(password)
	passwordConfirmCheck := helpers.IsEmpty(passwordConfirm)

	if usernameCheck || emailCheck || passwordCheck || passwordConfirmCheck {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Error (21): There is at least one empty required field.")
		return
	}

	// Password ile ConfirmPassword aynı değerlere sahip mi, kontrol edelim.
	if password != passwordConfirm {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Error (21): Passwords didn't matched. Please check it out!")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Success!")
	return

}

// LoginHandler is ...
func LoginHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("Request: '/login' ")

	// Önce client'tan gelen requestte bulunan datayı parse ediyoruz.
	r.ParseForm()
	for key, value := range r.Form {
		fmt.Printf("%s - %s\n", key, value)
	}

	// queryParams veya Body > "x-www-form-encoded" tarafından gönderilenleri parse eder.
	email = r.FormValue("email")
	password = r.FormValue("password")

	// Değerlerin dolu mu boş mu geldiğini kontrol edelim.
	emailCheck := helpers.IsEmpty(email)
	passwordCheck := helpers.IsEmpty(password)

	if emailCheck || passwordCheck {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Error (21): There is at least one empty required field.")
		return
	}

	var dbEmail string = "berk@safran.com"
	var dbPassword string = "12345"

	if password == dbPassword && email == dbEmail {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Success!")
		return
	}

	w.WriteHeader(404)
	fmt.Fprintf(w, "Error: Incorrect password or email!")
	return
}
