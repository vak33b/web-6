package main

import (
	"fmt"
	"net/http"
)

// Обработчик HTTP-запросов
func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello " + r.URL.Query().Get("name") + "!"))
}

func main() {
	// Регистрируем обработчик для пути "/"
	http.HandleFunc("/api/user", handler)

	// Запускаем веб-сервер на порту 8080
	fmt.Println("starting server...")
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}
