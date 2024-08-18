package homecontroller

import (
	"html/template"
	"net/http"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("assets/all/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, nil)
}

func About(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("assets/all/about.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, nil)
}

func Obat(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("assets/all/pasien.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, nil)
}

func Contact(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("assets/all/contact.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, nil)
}
