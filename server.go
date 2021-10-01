package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type TimeStruct struct {
	Time string `json:"time"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Go to /time")
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	result := TimeStruct{Time: time.Now().Format(time.RFC3339)}
	json.NewEncoder(w).Encode(result)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/time", timeHandler)

	log.Println("Server running at http://localhost/:8795")
	http.ListenAndServe(":8795", nil)
}
