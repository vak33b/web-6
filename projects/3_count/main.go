package main

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

// Счетчик и мьютекс для потокобезопасного доступа к счетчику
var (
	counter int
	mu      sync.Mutex
)

func countHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// Отправляем текущее значение счетчика
		mu.Lock()
		defer mu.Unlock()
		fmt.Fprintf(w, "Счетчик: %d", counter)
	case http.MethodPost:
		// Пытаемся получить значение count из формы
		r.ParseForm()
		countStr := r.FormValue("count")

		// Преобразование строки в число
		count, err := strconv.Atoi(countStr)
		if err != nil {
			http.Error(w, "Это не число", http.StatusBadRequest)
			return
		}

		// Увеличиваем счетчик
		mu.Lock()
		counter += count
		mu.Unlock()

		// Подтверждение успешного добавления
		fmt.Fprintf(w, "Счетчик увеличен на %d. Текущее значение: %d", count, counter)
	default:
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
	}
}

func main() {
	// Определяем обработчик для пути /count
	http.HandleFunc("/count", countHandler)

	fmt.Println("Сервер запущен на порту :3333")
	// Запускаем сервер на порту 3333
	if err := http.ListenAndServe(":3333", nil); err != nil {
		fmt.Println("Ошибка при запуске сервера:", err)
	}
}
