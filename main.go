package main

import (
	"log"
	"net/http"
)

func main()  {
	
	const port string = "8080" 
	mux := http.NewServeMux()
	corsMux := middlewareCors(mux)

	
	// creating new http server
	server := &http.Server{
		Handler: corsMux,
		Addr: ":" + port,
	}

	log.Printf("Listening on port %v", port)
	log.Fatal(server.ListenAndServe())

}

func middlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}