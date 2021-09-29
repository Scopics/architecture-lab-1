package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type ResultTime struct {
	Time string `json:"time"`
}

func helloHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "text/html")
	rw.WriteHeader(http.StatusAccepted)
	rw.Write([]byte("<p>Hello traveler, you probably came here to see the " +
		"<a href=\"/time\">time</a>, do it as soon as possible</p>"))
}

func timeHandler(rw http.ResponseWriter, req *http.Request) {
	resObj := ResultTime{time.Now().Format(time.RFC3339)}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusAccepted)

	if err := json.NewEncoder(rw).Encode(resObj); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/time", timeHandler)

	log.Println("Server started at http://localhost:8795")
	log.Fatal(http.ListenAndServe(":8795", nil))
}
