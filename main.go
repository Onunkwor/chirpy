package main

import (
	"log"
	"net/http"
)
func main () {
 mux := http.NewServeMux()

 mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
	// Set Content-Type first (must be set before WriteHeader or Write)
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	// Explicitly write the 200 status code
		w.WriteHeader(http.StatusOK)

		// Write the body. We ignore the returned (n, err) here for brevity.
		_, _ = w.Write([]byte("OK"))
 })

 //http.Dir(".") converts the filesystem path to an http.FileSystem.
	fileServer := http.FileServer(http.Dir("./public"))
	mux.Handle("/app/", http.StripPrefix("/app/", fileServer))

	// 3) Register the file server for the root path "/".
	//    When the browser requests "/", the FileServer will look for "index.html".
	mux.Handle("/", fileServer)
 // (No handlers registered on mux -> requests will return 404)

 srv := &http.Server{
	  Addr:    ":8080",
	  Handler: mux,
 }
 // 3) Start the server
	log.Printf("starting server on %s (open http://localhost:8080)\n", srv.Addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server error: %v", err)
	}
}