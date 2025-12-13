package main

import (
	"GoChat/internal/server"
	"fmt"
)


func main() {
	server := Server.CreateServer("127.0.0.1", ":8080")
	err := server.StartServer()
	if err != nil {
		fmt.Println("Failed To Start Server")
	}
	 

}
