package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

func countByArr(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT arr, COUNT(*) FROM arbres_remarquables GROUP BY arr")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var results []CountResult
		for rows.Next() {
			var r CountResult
			if err := rows.Scan(&r.Arr, &r.Count); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			results = append(results, r)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(results)
	}
}

func avgHeightByArr(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query(`
			SELECT arr, AVG(hauteur) 
			FROM arbres_remarquables 
			WHERE hauteur > 0 
			GROUP BY arr 
			ORDER BY arr`)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var results []AvgHeightResult
		for rows.Next() {
			var r AvgHeightResult
			if err := rows.Scan(&r.Arr, &r.AvgHauteur); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			results = append(results, r)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(results)
	}
}

func countByGenre(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query(`
			SELECT genre, COUNT(*) 
			FROM arbres_remarquables 
			WHERE genre <> '' 
			GROUP BY genre 
			ORDER BY COUNT(*) DESC`)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var results []GenreCountResult
		for rows.Next() {
			var r GenreCountResult
			if err := rows.Scan(&r.Genre, &r.Count); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			results = append(results, r)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(results)
	}
}
