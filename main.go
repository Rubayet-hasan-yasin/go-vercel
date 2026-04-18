package main

import (
	"fmt"
	handler "go-vercel/api"
	"log"
	"net/http"
	"os"
)
 
func main() {
  mux := http.NewServeMux()
  mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello from Go on Vercel")
  })
  mux.HandleFunc("/static", handler.Handler)

  port := os.Getenv("PORT")
  if port == "" {
    port = "3000"
  }

  log.Fatal(http.ListenAndServe(":"+port, mux))
}