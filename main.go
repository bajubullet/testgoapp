package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	port := GetPort()
	log.Println("Listening on port", port, "...")
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}

// Get the Port from the environment so we can run on Heroku
func GetPort() (port string) {
	port = os.Getenv("PORT")
	if port == "" {
		port = "4747"
		log.Println(
			"INFO: No PORT environment variable detected, defaulting to", port)
	}
	return
}
