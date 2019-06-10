package main

import (
	"github.com/kimsw9599/alice/sheets/doc"
	"log"
)

func main() {
	id,  err := doc.Make_sheets("TestFiles222")

	if err!= nil{
		log.Fatalf("Create sheets error: %v", err)
	}

	print(id)
}