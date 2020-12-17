package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func myhandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello there!!"))
}
func swapoff(w http.ResponseWriter, r *http.Request) {
	cmd, _ := exec.Command("ansible", "-i", "inventory/hosts", "k8s", "-a", "swpapoff -a").Output()
	w.WriteHeader(http.StatusOK)
	w.Write(cmd)
}
func setupMuxRouter() *mux.Router {
	router := mux.NewRouter()
	api := router.PathPrefix("/api/v1").Subrouter()
	apiGeneric := router.PathPrefix("/api").Subrouter()

	apiGeneric.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	})
	api.HandleFunc("/test", myhandler)
	api.HandleFunc("/swapoff", swapoff)
	// apiGet.HandleFunc("/allns", getallns)
	// apiGet.HandleFunc("/onens", getonens).Queries("name", "{name}")
	return router
}
func main() {
	muxRouter := setupMuxRouter()
	loggedRouter := handlers.LoggingHandler(os.Stdout, muxRouter)
	err := http.ListenAndServe(":8080", loggedRouter)
	if err != nil {
		fmt.Println(err)
	}

}
