package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	route := vars["route"]

	url, err := getURL(route)
	if err != nil {
		log.Println(err)
	}

	w.Write([]byte(url + "\n"))
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	route := vars["route"]

	url, err := getURL(route)
	if err != nil {
		log.Println(err)
	}

	http.Redirect(w, r, url, 307)
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	route := vars["route"]

	url, err := getURL(route)
	if err != nil {
		log.Println(err)
	}

	type Url struct {
		URL string `json:"url"`
	}

	w.Header().Set("Content-Type", "application/json")
	j, _ := json.Marshal(Url{url})
	w.Write(j)
}

func getURL(route string) (string, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	url, err := client.SRandMember(route).Result()
	if err != nil {
		return "", err
	}

	return url, nil
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{route:(?:success|fail)}", redirectHandler).Methods("GET").Queries("redirect", "1")
	r.HandleFunc("/{route:(?:success|fail)}.gif", redirectHandler).Methods("GET")
	r.HandleFunc("/{route:(?:success|fail)}", jsonHandler).Methods("GET").Headers("Accept", "application/json")
	r.HandleFunc("/{route:(?:success|fail)}", jsonHandler).Methods("GET").Queries("json", "1")
	r.HandleFunc("/{route:(?:success|fail)}.json", jsonHandler).Methods("GET")
	r.HandleFunc("/{route:(?:success|fail)}", handler).Methods("GET")
	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	log.Fatal(http.ListenAndServe(":8000", loggedRouter))
}
