package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var helpFlag = flag.Bool("help", false, "Display usage information")
var portFlag = flag.String("port", "80", "Port to run the server on")

func SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		entity := &Greeting{Message: "Hello World"}
		b, err := json.Marshal(entity)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(b)
	})
}

func IsHelp() bool {
	return *helpFlag
}

func DisplayUsage() {
	fmt.Println("Usage:")
	flag.VisitAll(func(f *flag.Flag) {
		if f.Name == "help" {
			return
		}
		fmt.Printf("  -%s: %s\n", f.Name, f.Usage)
	})
}

func Startup() {
	flag.Parse()
	if IsHelp() {
		DisplayUsage()
		return
	}

	mux := http.NewServeMux()

	port := os.Getenv("PORT")
	if port == "" && portFlag != nil {
		port = *portFlag
	}

	SetupRoutes(mux)
	muxChain := AccessLog(mux)

	log.Printf("Server started on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, muxChain))
}

func main() {
	Startup()
}
