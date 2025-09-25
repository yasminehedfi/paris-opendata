# Projet Open Data Paris – Arbres Remarquables

## Description
Projet Open Data Paris : récupérer les données d’arbres via l’API, stocker dans PostgreSQL et visualiser avec un frontend simple.


## Contenu
- `backend/` : code Go pour récupérer les données et exposer une API
- `frontend/` : page HTML avec graphiques (Chart.js)
- `docker-compose.yml` : orchestration des conteneurs
- `Dockerfile` : construction du backend Go

## Lancer 
   docker-compose up --build

## Fonctionnalités
Endpoints API :
 - /api/count_by_arr : nombre d’arbres par arrondissement
 - /api/avg_height_by_arr : hauteur moyenne par arrondissement
 - /api/count_by_genre : nombre d’arbres par genre

Frontend : http://localhost:8080
Backend API : http://localhost:8081