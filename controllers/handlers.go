package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/sergiordob/Golang-API-REST-01/database"
	"github.com/sergiordob/Golang-API-REST-01/models"
)

func Index() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		json.NewEncoder(writer).Encode(models.States)
	}
}

//É importante notar que a funcionalidade de captura de valores de variáveis de caminho com request.PathValue("id") é uma adição recente no Go 1.22 e não está disponível nas versões anteriores do Go
func GetOneState() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var singleState models.State
		id := request.PathValue("id")
		database.DatabaseConnection.First(&singleState, id)
		json.NewEncoder(writer).Encode(singleState)	
	}
}

func GetAllStates() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var states []models.State
		database.DatabaseConnection.Find(&states)
		err := json.NewEncoder(writer).Encode(states)
		if err != nil {
			log.Println("Error: ", err)
		}
	}
}