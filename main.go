package main

// Go code formatter to clean up any ugly indents and such.
import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// func handler(w http.ResponseWriter, r *http.Request) {
// 	io.WriteString(w, "../static/index.html")
// }

//Route
func router() *mux.Router {
	r := mux.NewRouter()
	// r.HandleFunc("/", handler)
	r.Handle("/", http.FileServer(http.Dir("./static/")))
	// r.HandleFunc("/", homeHandler)
	// r.HandleFunc("/about", aboutHandler)
	// r.HandleFunc("/contact", contactHandler)
	return r
}

// Initiate web server
func main() {
	fmt.Println(" Starting web server at `localhost:9100`\n Press ctrl+C to close.")
	router := router()
	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:9100",
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
