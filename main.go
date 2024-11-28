package main

import (
	"flag"
	"fmt"
	"learning-go/helloworld"
	"learning-go/middleware"
)

func main() {
	mode := flag.String("mode", "basic", "Choose middleware mode: basic or advanced")
	flag.Parse()

	switch *mode {
	case "basic":
		fmt.Println("Running basic middleware")
		middleware.BasicMiddleware()
	case "advanced":
		fmt.Println("Running advanced middleware")
		middleware.AdvancedMiddleware()
	case "helloworld":
		fmt.Println("Running helloworld")
		helloworld.HelloWorld()
	default:
		fmt.Println("Invalid mode")
	}
}
