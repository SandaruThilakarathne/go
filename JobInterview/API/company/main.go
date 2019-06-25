package main

import (
	// "github.com/JobInterview/API/geo"
	"github.com/JobInterview/API/company/serivces"
	"fmt"
	// "API\services"
)

var appName = "companySerive"

func main() {
	fmt.Println("Starting %v\n", appName)
	// services.StartWebServer("6767")
	// geo.StartServer(":6767")
	serivces.StartServer(":6768")

	
}