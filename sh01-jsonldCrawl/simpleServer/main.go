package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type MyServer struct {
	r *mux.Router
}

func main() {
	fmt.Println("JSON-LD simple server")

	rcommon := mux.NewRouter()
	rcommon.PathPrefix("/files/").Handler(http.StripPrefix("/files/", http.FileServer(http.Dir("./files"))))
	http.Handle("/files/", &MyServer{rcommon})

	// Start the server...
	log.Printf("About to listen on 9900. Go to http://127.0.0.1:9900/")

	err := http.ListenAndServe(":9900", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func (s *MyServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	rw.Header().Set("Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	s.r.ServeHTTP(rw, req)
}

func addDefaultHeaders(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		fn(w, r)
	}
}
