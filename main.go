package main

// Go code formatter to clean up any ugly indents and such.
import (
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func handler(w, http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

//Route
func router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	return r
}

// Initiate web server
func main() {
	router := router()
	srv := &http.Server{
		Handler: router,
		Addr: "127.0.0.1:9100",
		WriteTimeout: 10 * time.Second,
		ReadTimeout: 10 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
