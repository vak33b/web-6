package main

import (
	"encoding/json"
	"net/http"
	"sync"
)

var (
	count int
	mu    sync.Mutex
)

func main() {
	http.HandleFunc("/count", countHandler)
	http.ListenAndServe(":3333", enableCors(http.DefaultServeMux))
}

func countHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		mu.Lock()
		defer mu.Unlock()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(count)
	} else if r.Method == http.MethodPost {
		var req map[string]int
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "это не число", http.StatusBadRequest)
			return
		}
		value, exists := req["count"]
		if !exists || value < 0 {
			http.Error(w, "это не число", http.StatusBadRequest)
			return
		}
		mu.Lock()
		count += value
		mu.Unlock()
		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, "метод не поддерживается", http.StatusMethodNotAllowed)
	}
}

// enableCors добавляет заголовки CORS
func enableCors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Обработка preflight-запросов
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		h.ServeHTTP(w, r)
	})
}
