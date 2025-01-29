package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/FedericoBarberon/Go-Exercises/cyoa"
)

func main() {
	book, err := cyoa.GetBookFromFS(os.DirFS("./"), "gopher.json")

	if err != nil {
		log.Fatal("error getting book from fs", err)
		os.Exit(1)
	}

	cyoaHandler, err := cyoa.NewHandler(book)

	if err != nil {
		log.Fatal("error creating handler", err)
		os.Exit(1)
	}

	port := 8000

	log.Printf("Listening server on port %d", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), cyoaHandler)
}
