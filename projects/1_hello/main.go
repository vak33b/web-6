package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/get" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "Hello, web!")
}

func main() {
	http.HandleFunc("/get", handler)

	fmt.Println("Сервер запущен на порту :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Ошибка при запуске сервера:", err)
	}
}
