package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/WYC51/restful-api-demo/utils"

	"github.com/gorilla/mux"
)

// json_file_path is the path to the example json file
const json_file_path = "./data/example_data.json"

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", SayHello).Methods("GET")
	r.HandleFunc("/show_demo_data", ShowDemo).Methods("GET")
	r.HandleFunc("/show_demo_data", AddData).Methods("POST")
	r.HandleFunc("/show_demo_data/{id}", ShowID).Methods("GET")
	r.HandleFunc("/show_demo_data/{id}", UpdateData).Methods("PUT")
	r.HandleFunc("/show_demo_data/{id}", DeleteData).Methods("DELETE")
	
	log.Fatal(http.ListenAndServe(":8080", r))
}

func DeleteData(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	err = utils.Delete_Data(json_file_path, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "{\"message\":\"OK\"}\n")
}


func UpdateData(w http.ResponseWriter, r *http.Request) {
	params :=mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	var data utils.Info

	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = utils.Update_Data(json_file_path, data, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "{\"message\":\"OK\"}\n")

}

func ShowID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	found := false
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	json_data, _ := utils.Load_Json(json_file_path)
	for i, v := range json_data {
		if int(v["id"].(float64)) == id {
			found = true
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(json_data[i])
			break
		} 
	}
	if !found {
			http.Error(w, "ID not found", http.StatusBadRequest)
		}
}

func AddData(w http.ResponseWriter, r *http.Request) {
	var data utils.Info
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = utils.Add_Data(json_file_path, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return

	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "{\"message\":\"OK\"}\n")
}

func ShowDemo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json_data, _ := utils.Load_Json(json_file_path)
	json.NewEncoder(w).Encode(json_data)
}

func SayHello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Hello, Welcome to my RESTful API!")
}
