// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"sync/atomic"
// )

// type apiConfig struct {
// 	fileserverHits atomic.Int32
// }

// func main() {
// 	mux := http.NewServeMux()

// 	// Create config struct to hold the hit counter
// 	cfg := &apiConfig{}

// 	// Health check endpoint
// 	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
// 		w.WriteHeader(http.StatusOK)
// 		_, _ = w.Write([]byte("OK"))
// 	})

// 	// File server for /app/
// 	fileServer := http.FileServer(http.Dir("./public"))
// 	mux.Handle("/app/", http.StripPrefix("/app", cfg.middlewareMetricsInc(fileServer)))

// 	// Metrics endpoint
// 	mux.Handle("GET /api/metrics", cfg.getFileserverHits())

// 	// Reset endpoint
// 	mux.Handle("POST /api/reset", cfg.resetFileserverHits())

// 	// Create server
// 	srv := &http.Server{
// 		Addr:    ":8080",
// 		Handler: mux,
// 	}

// 	log.Printf("starting server on %s (open http://localhost:8080)\n", srv.Addr)
// 	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
// 		log.Fatalf("server error: %v", err)
// 	}
// }

// // Middleware to increment hits counter
// func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		cfg.fileserverHits.Add(1)
// 		log.Printf("%s %s", r.Method, r.URL.Path)
// 		next.ServeHTTP(w, r)
// 	})
// }

// // Handler to return hit count
// func (cfg *apiConfig) getFileserverHits() http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
// 		w.WriteHeader(http.StatusOK)
// 		_, _ = w.Write([]byte(fmt.Sprintf("Hits: %d", cfg.fileserverHits.Load())))
// 	})
// }

// // Handler to reset hit count
// func (cfg *apiConfig) resetFileserverHits() http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		cfg.fileserverHits.Store(0)
// 		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
// 		w.WriteHeader(http.StatusOK)
// 		_, _ = w.Write([]byte("Hits reset to 0"))
// 	})
// }

package main

import "net/http"



func main () {
	mux := http.NewServeMux()
	api := &api {
		addr: ":8080",
	}
   srv := &http.Server{
		Addr: api.addr,
	Handler: mux,	
}
mux.HandleFunc(http.MethodGet + "/api/users", api.getUserHAndler)
mux.HandleFunc(http.MethodPost + "/api/users", api.createUserHandler)
 srv.ListenAndServe()
}