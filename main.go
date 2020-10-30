package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Dharmik-dbcorp/SampleRestApiGO/helper"
	"github.com/Dharmik-dbcorp/SampleRestApiGO/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
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

func getPatients(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	// created Patients array
	var patients []models.Patient

	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		helper.GetError(err, w)
		return
	}

	// Close the cursor once finished
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		var patient models.Patient
		// & character returns the memory address of the following variable.
		err := cur.Decode(&patient)
		if err != nil {
			log.Fatal(err)
		}

		// add item our array
		patients = append(patients, patient)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(patients) // encode similar to serialize process.
}

func main() {
	//Init Router
	r := mux.NewRouter()

	r.HandleFunc("/api/patients", getPatients).Methods("GET")
	r.HandleFunc("/api/patient", createPatient).Methods("POST")

	config := helper.GetConfiguration()
	log.Fatal(http.ListenAndServe(config.Port, r))

}
