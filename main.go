package main

import (
	"fmt"
	"log"
	"net/http"
)

// compile passing -ldflags "-X main.versionNumber <build>"
var versionNumber = ""

func main() {
	fmt.Printf("GetVer Version: %s\n", versionNumber)

	router := NewRouter()

	// Port number must be fetched by KV store like consul
	log.Fatal(http.ListenAndServe(":9080", router))

	//crawler()
}
