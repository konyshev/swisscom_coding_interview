package main

import (
	"encoding/json"
	"log"
	"net/http"

	. "github.com/konyshev/swisscom_comodir_restapi/config"
	. "github.com/konyshev/swisscom_comodir_restapi/dao"
	. "github.com/konyshev/swisscom_comodir_restapi/models"

	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
)

var config = Config{}
var dao = InstanceDAO{}

func AllInstancesEndPoint(w http.ResponseWriter, r *http.Request) {
	instances, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, instances)
}

func FindInstanceEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	instance, err := dao.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Instance ID")
		return
	}
	respondWithJson(w, http.StatusOK, instance)
}

func CreateInstanceEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var instance Instance
	if err := json.NewDecoder(r.Body).Decode(&instance); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	instance.ID = bson.NewObjectId()
	if err := dao.Insert(instance); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, instance)
}

func UpdateInstanceEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var instance Instance
	if err := json.NewDecoder(r.Body).Decode(&instance); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Update(instance); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func DeleteInstanceEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var instance Instance
	if err := json.NewDecoder(r.Body).Decode(&instance); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Delete(instance); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/instances", AllInstancesEndPoint).Methods("GET")
	r.HandleFunc("/instances", CreateInstanceEndPoint).Methods("POST")
	r.HandleFunc("/instances", UpdateInstanceEndPoint).Methods("PUT")
	r.HandleFunc("/instances", DeleteInstanceEndPoint).Methods("DELETE")
	r.HandleFunc("/instances/{id}", FindInstanceEndpoint).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
