package main

import (
	db "backend/pkg/db/sqlite"
	"encoding/json"
	"log"
	"net/http"
)

type User struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	BirthDate string `json:"birth_date"`
}

func main() {
	database := db.InitDB()

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
			return
		}

		var user User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, "Erreur JSON", http.StatusBadRequest)
			return
		}

		_, err = database.Exec(`
            INSERT INTO users (email, password, first_name, last_name, birth_date)
            VALUES (?, ?, ?, ?, ?)`,
			user.Email, user.Password, user.FirstName, user.LastName, user.BirthDate,
		)

		if err != nil {
			log.Println("Erreur insertion utilisateur :", err)
			http.Error(w, "Erreur création utilisateur", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Utilisateur enregistré avec succès"})
	})

	log.Println("🚀 Serveur lancé sur : http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
