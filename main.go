/*
 * OIDC IAT Userinfo Endpoint
 *
 * Endpoint for OpenID Connect's ID Assertion Token endpoint for userinfo.
 *
 * API version: 0.4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package main

import (
	"log"
	"net/http"
	"os"

	iat "iat/go"
)

func main() {
	log.Printf("Starting Server...")

	// Load configuration
	log.Printf("Loading configuration...")
	iat.Initialize()

	router := iat.NewRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Configuration loaded")

	log.Printf("Running on port " + port)

	log.Fatal(http.ListenAndServe(":"+port, router))
}
