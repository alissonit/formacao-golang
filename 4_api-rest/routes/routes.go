package routes

import (
	"log"
	"net/http"

	"github.com/alissonit/api-rest/controllers"
	"github.com/alissonit/api-rest/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.SetMiddlewareContentType) // Use middleware
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalities", controllers.GetPersonalities).Methods("GET")
	r.HandleFunc("/api/personalities/{id}", controllers.GetPersonalityFromId).Methods("GET")
	r.HandleFunc("/api/personalities", controllers.CreatePersonality).Methods("POST")
	r.HandleFunc("/api/personalities/{id}", controllers.DeletePersonality).Methods("DELETE")
	r.HandleFunc("/api/personalities/{id}", controllers.UpdatePersonality).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
