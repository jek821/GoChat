package main

import (
	"GoChat/internal/logger"
	"GoChat/internal/server"
	"GoChat/pkg"
)


func main() {
	pkg.Init() // Initialize Gob (Registers Protocol Structs)	
	server := Server.CreateServer("127.0.0.1", ":8080")
	errChan := make(chan error,1)
	go func(){
		errChan <- server.StartServer()
	}()

	if err := <-errChan; err != nil {
		logger.Logger.Info("Server Failed To Start")
	}

	 

}
