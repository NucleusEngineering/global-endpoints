package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/NucleusEngineering/global-endpoints/runinspect"
)

func main() {
	region, err := runinspect.Region()
	if err != nil {
		log.Fatalf("failed to read region: %v", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "I am a web server, responding from %s!", region)
	})

	port, err := runinspect.Port()
	if err != nil {
		log.Fatalf("failed to read port: %v", err)
	}

	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
