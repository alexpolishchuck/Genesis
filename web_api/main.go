package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Print("Starting a server")
	handler := http_request_handler{"5000", file_handler{"emails.csv", sync.RWMutex{}}}

	err := handler.start_server()
	if err != nil {
		fmt.Println("Error occurred while starting server")
	}
}
