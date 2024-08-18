package main

import (
	"go-web-native/config"
	"go-web-native/controller/homecontroller"
	"log"
	"net/http"
)

func main() {

	config.ConnectDB()

	log.Println("Server running in port 8000")

	http.HandleFunc("/", homecontroller.Welcome)
	http.HandleFunc("/About", homecontroller.About)
	http.HandleFunc("/Pasien", homecontroller.Obat)
	http.HandleFunc("/Contact", homecontroller.Contact)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))

	http.ListenAndServe(":8000", nil)
}
