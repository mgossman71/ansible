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
	cmd, _ := exec.Command("ansible", "-i", "./inventory/hosts.yaml", "k8s", "-a", "\"swapoff -a\"").Output()

	// cmd, _ := exec.Command("ansible", "--version").Output()
	// log.Printf("Trying to run Swapoff -a")

	w.WriteHeader(http.StatusOK)
	w.Write(cmd)
}
func test() {
	// fmt.Println("Test Function.")
	cmd := exec.Command("ansible", "-i", "./inventory/hosts.yaml", "k8s", "-a", "\"swapoff -a\"").String()
	fmt.Println(cmd)

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
	test()
	// os.Stdout.WriteString("your message here\n")

	muxRouter := setupMuxRouter()
	loggedRouter := handlers.LoggingHandler(os.Stdout, muxRouter)
	err := http.ListenAndServe(":8080", loggedRouter)
	if err != nil {
		fmt.Println(err)
	}
}
