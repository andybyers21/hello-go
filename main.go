package main

// Go code formatter to clean up any ugly indents and such.
import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Route
func router() *mux.Router {
	r := mux.NewRouter()
	r.Handle("/", http.FileServer(http.Dir("./static/")))
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("static/"))))
	return r
}

// Initiate web server
func main() {
	log.Println(" Starting web server at `localhost:9100`\nPress ctrl+C to close.")
	router := router()
	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:9100",
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
