package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var counter int

func main() {
	// Загрузка счетчика из файла (опционально)
	loadCounterFromFile("counter.txt")

	http.HandleFunc("/count", handleCount)
	log.Fatal(http.ListenAndServe(":3333", nil))
}

func handleCount(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprintln(w, counter) // Возвращаем счетчик

	case "POST":
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Ошибка чтения тела запроса", http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		var data struct {
			Count int `json:"count"`
		}
		err = json.Unmarshal(body, &data)
		if err != nil {
			http.Error(w, "Ошибка разбора JSON", http.StatusBadRequest)
			return
		}

		if data.Count == 0 {
			http.Error(w, "это не число", http.StatusBadRequest)
			return
		}
		counter += data.Count
		saveCounterToFile("counter.txt")
		fmt.Fprintln(w, counter)

	default:
		http.Error(w, "Недопустимый метод", http.StatusMethodNotAllowed)
	}
}

func loadCounterFromFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			log.Println("Файл счетчика не найден, устанавливаем значение 0")
			return
		}
		log.Printf("Ошибка открытия файла: %v", err)
		return
	}
	defer file.Close()

	_, err = fmt.Fscanln(file, &counter)
	if err != nil {
		log.Printf("Ошибка чтения файла: %v", err)
		return
	}
	log.Println("Значение счетчика загружено из файла")
}

func saveCounterToFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Printf("Ошибка создания файла: %v", err)
		return
	}
	defer file.Close()

	_, err = fmt.Fprintln(file, counter)
	if err != nil {
		log.Printf("Ошибка записи в файл: %v", err)
		return
	}
	log.Println("Счетчик сохранен в файл")
}
