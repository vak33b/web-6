package main

import (
	"fmt"
	"net/http"
)

// Middleware для CORS
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")                   // Разрешаем все origins
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS") // Разрешаем необходимые методы
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")       // Разрешаем заголовки
		// Обрабатываем preflight запрос
		if r.Method == http.MethodOptions {
			return
		}
		next.ServeHTTP(w, r)
	})
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello, %s!", name)
}

func main() {
	http.Handle("/api/user", corsMiddleware(http.HandlerFunc(helloHandler)))
	fmt.Println("Сервер запущен на порту :8083")
	http.ListenAndServe(":8083", nil)
}
