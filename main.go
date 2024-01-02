package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/melsheikh92/go-microservices/details"
)

func main() {
	fmt.Println("Hello world!")
	router := mux.NewRouter()

	router.HandleFunc("/book/{id}/{page}", func(wr http.ResponseWriter, rq *http.Request) {
		vars := mux.Vars(rq)
		id := vars["id"]
		page := vars["page"]
		fmt.Fprintf(wr, "Book id: %s, page: %s", id, page)
	})

	router.HandleFunc("/health", func(wr http.ResponseWriter, req *http.Request) {
		response := map[string]string{
			"status":    "up",
			"timestamp": time.Now().String(),
		}
		json.NewEncoder(wr).Encode(response)
	})

	router.HandleFunc("/", func(wr http.ResponseWriter, req *http.Request) {
		wr.WriteHeader(http.StatusOK)
		fmt.Fprintf(wr, "Application is up and running")
	})

	router.HandleFunc("/details", func(wr http.ResponseWriter, req *http.Request) {
		host, err := details.GetHostName()
		if err != nil {
			wr.WriteHeader(http.StatusExpectationFailed)
			return
		}
		wr.WriteHeader(http.StatusOK)
		addrs := details.GetLocalIp()
		fmt.Fprintf(wr, "HostName: %s, local ip: %s", host, addrs)
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Fatal(http.ListenAndServe(":80", router))
}
