package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Dharmik-dbcorp/SampleRestApiGO/helper"
	"github.com/Dharmik-dbcorp/SampleRestApiGO/models"
	"github.com/gorilla/mux"
)

//Connection mongoDB with helper class
var collection = helper.ConnectDB()

func createPatient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var patient models.Patient

	// we decode our body request params
	_ = json.NewDecoder(r.Body).Decode(&patient)

	result, err := collection.InsertOne(context.TODO(), patient)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func main() {
	//Init Router
	r := mux.NewRouter()

	r.HandleFunc("/api/patient", createPatient).Methods("POST")

	config := helper.GetConfiguration()
	log.Fatal(http.ListenAndServe(config.Port, r))

}
