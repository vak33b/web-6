package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/get" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "Hello, web!")
}

func main() {
	http.HandleFunc("/get", helloHandler)

	fmt.Println("Сервер запущен на порту :8080")

	// Разрешаем CORS
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "OPTIONS"})

	if err := http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(http.DefaultServeMux)); err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}
