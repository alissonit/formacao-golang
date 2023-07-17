package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/alissonit/api-rest/database"
	"github.com/alissonit/api-rest/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func GetPersonalities(w http.ResponseWriter, r *http.Request) {

	var personality []models.Personality

	database.DB.Find(&personality)

	json.NewEncoder(w).Encode(personality)
}

func GetPersonalityFromId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
	}

	var personality models.Personality

	database.DB.First(&personality, id)

	json.NewEncoder(w).Encode(personality)
}

func CreatePersonality(w http.ResponseWriter, r *http.Request) {
	var newPersonality models.Personality

	json.NewDecoder(r.Body).Decode(&newPersonality)

	database.DB.Create(&newPersonality)

	json.NewEncoder(w).Encode(newPersonality)
}

func DeletePersonality(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
	}

	var personality models.Personality

	database.DB.Delete(&personality, id)

	json.NewEncoder(w).Encode(personality)
}

func UpdatePersonality(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
	}

	var personality models.Personality

	database.DB.First(&personality, id)

	json.NewDecoder(r.Body).Decode(&personality)

	database.DB.Save(&personality)

	json.NewEncoder(w).Encode(personality)
}
