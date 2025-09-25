package main

type Arbre struct {
	ID       string
	NomUsuel string
	NomLatin string
	Genre    string
	Espece   string
	Arr      string
	Lat      float64
	Lon      float64
	Hauteur  float64
	UrlPDF   string
	UrlPhoto string
	Resume   string
}

type CountResult struct {
	Arr   string `json:"arr"`
	Count int    `json:"count"`
}

type AvgHeightResult struct {
	Arr        string  `json:"arr"`
	AvgHauteur float64 `json:"avg_hauteur"`
}

type GenreCountResult struct {
	Genre string `json:"genre"`
	Count int    `json:"count"`
}
