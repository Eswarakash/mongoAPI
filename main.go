package main

import (
	"fmt"
	"log"
	"net/http"

	"git.com/Eswarakash/mongoAPI/router"
)

func main() {
	fmt.Println("mongoDB API")
	r := router.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at port 4000 ...")
}
