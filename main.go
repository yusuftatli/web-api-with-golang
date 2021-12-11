package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)




func main() {
	fmt.Println("starting server")



	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		bin, _ := json.Marshal("params")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(bin)
	})

log.Fatal(http.ListenAndServe(":" + os.Getenv("PORT"), nil))
}
// func main(){
// 	fmt.Println("starting server")

// 	api := mux.NewRouter().PathPrefix("/api").Subrouter()
// 	api.HandleFunc("/getdata/{params:[0-9]+}", Handler).Methods(http.MethodGet)
// 	// http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), api)
// 	log.Fatal(http.ListenAndServe(":" + os.Getenv("PORT"),api))

// }

// func Handler(w http.ResponseWriter, r *http.Request) {
// 	params, _ := mux.Vars(r)["params"]

// 	// if err != nil {
// 	// 	w.WriteHeader(http.StatusInternalServerError)
// 	// 	return
// 	// }

// 	bin, _ := json.Marshal(params)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	_, _ = w.Write(bin)
// }