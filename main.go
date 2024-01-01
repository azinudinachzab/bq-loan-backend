package main

import (
	"github.com/azinudinachzab/bq-loan-be-v2/app"

	"log"
	"os"
)

func init() {
	if err := os.Setenv("TZ", "Asia/Jakarta"); err != nil {
		log.Fatalf("cant load timezone %v\n", err)
	}
}

func main() {
	app.New().Run()
}
