package main

import (
	"github.com/JobInterview/API/geo"
	// "github.com/JobInterview/API/company"
	"fmt"
	// "API\services"
)

var appName = "geoservices"

func main() {
	fmt.Println("Starting %v\n", appName)
	// services.StartWebServer("6767")
	geo.StartServer(":6767")
	// company.StartServer(":6768")

	
}