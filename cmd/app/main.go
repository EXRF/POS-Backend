package main

import (
	"github.com/EXRF/POS-Backend/internal/delivery/http"
	"github.com/EXRF/POS-Backend/pkg/utils"
)

func main() {
	// Load environment variables from .env file if it exists
	utils.LoadEnv()
	http.RunServer()
}
