package main

import (
	"os"

	"github.com/azinudinachzab/bq-loan-backend/api"
)

func init() {
	os.Setenv("TZ", "Asia/Jakarta")
}

func main() {
	api.Run()
}
