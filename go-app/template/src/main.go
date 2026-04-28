package main

import (
	"embed"
	"encoding/json"
	"io/fs"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"
)

//go:embed web
var webFS embed.FS

const (
	appName = "${{values.app_name}}"
	appEnv  = "${{values.app_env}}"
)

func info(w http.ResponseWriter, _ *http.Request) {
	host, _ := os.Hostname()
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]string{
		"time":        time.Now().Format("03:04:05PM on January 02, 2006"),
		"hostname":    host,
		"deployed_on": "kubernetes",
		"env":         appEnv,
		"app_name":    appName,
		"language":    "Go " + runtime.Version(),
	})
}

func healthz(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "up"})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/info", info)
	mux.HandleFunc("/api/v1/healthz", healthz)

	staticFS, err := fs.Sub(webFS, "web")
	if err != nil {
		log.Fatalf("static fs: %v", err)
	}
	mux.Handle("/", http.FileServer(http.FS(staticFS)))

	addr := ":5000"
	log.Printf("%s (%s) listening on %s", appName, appEnv, addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
