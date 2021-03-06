/*
 * Accounts Analyser
 *
 * This API is for managing bank account statements, for use with [https://github.com/Accounts-Analyser](https://github.com/Accounts-Analyser)
 *
 * API version: 1.0.0
 * Contact: colincampbell321123@hotmail.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package main

import (
	"log"
	"net/http"

	sw "github.com/Accounts-Analyser/statement-api/go"
)

func main() {
	log.Printf("Server started")

	router := sw.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
