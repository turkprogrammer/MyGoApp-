package main

import (
	"fmt"
	"mygoapp/routes"
)

func main() {
	fmt.Println("Starting the application...")
	routes.SetupRoutes()
}
