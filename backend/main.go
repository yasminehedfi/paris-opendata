package main

import (
	"log"
	"net/http"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	db, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	loadData(db)

	mux := http.NewServeMux()
	mux.HandleFunc("/api/count_by_arr", countByArr(db))
	mux.HandleFunc("/api/avg_height_by_arr", avgHeightByArr(db))
	mux.HandleFunc("/api/count_by_genre", countByGenre(db))

	handler := enableCORS(mux)

	log.Println(" API lancée sur :8081")
	http.ListenAndServe(":8081", handler)
}
