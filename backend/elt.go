package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func loadData(db *sql.DB) {
	url := "https://opendata.paris.fr/api/explore/v2.1/catalog/datasets/arbresremarquablesparis/records?limit=100"

	resp, err := http.Get(url)
	if err != nil {
		log.Println("Erreur API:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		log.Println("Erreur JSON:", err)
		return
	}

	results, ok := result["results"].([]interface{})
	if !ok {
		log.Println("Erreur: 'results' n'est pas un tableau")
		return
	}

	for _, r := range results {
		rec := r.(map[string]interface{})

		geom := rec["geom_x_y"].(map[string]interface{})
		lat := geom["lat"].(float64)
		lon := geom["lon"].(float64)

		id := fmt.Sprintf("%v", rec["arbres_idbase"])
		arr := fmt.Sprintf("%v", rec["arbres_arrondissement"])

		hauteur := 0.0
		if rec["arbres_hauteurenm"] != nil {
			hauteur = rec["arbres_hauteurenm"].(float64)
		}

		nomUsuel := fmt.Sprintf("%v", rec["com_nom_usuel"])
		nomLatin := fmt.Sprintf("%v", rec["com_nom_latin"])
		genre := fmt.Sprintf("%v", rec["arbres_genre"])
		espece := fmt.Sprintf("%v", rec["arbres_espece"])
		urlPDF := fmt.Sprintf("%v", rec["com_url_pdf"])
		urlPhoto := fmt.Sprintf("%v", rec["com_url_photo1"])
		resume := fmt.Sprintf("%v", rec["com_resume"])

		_, err := db.Exec(`
			INSERT INTO arbres_remarquables
			(id, nom_usuel, nom_latin, genre, espece, arr, lat, lon, hauteur, url_pdf, url_photo1, resume)
			VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)
			ON CONFLICT (id) DO NOTHING`,
			id, nomUsuel, nomLatin, genre, espece, arr, lat, lon, hauteur, urlPDF, urlPhoto, resume,
		)
		if err != nil {
			log.Println("Erreur insert:", err)
		}
	}

	fmt.Println(" Données importées avec succès depuis l'API v2.1")
}
