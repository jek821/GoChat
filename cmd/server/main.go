package main

import (
	"GoChat/internal/server"
	"GoChat/internal/logger"

)


func main() {
	server := Server.CreateServer("127.0.0.1", ":8080")
	errChan := make(chan error,1)
	go func(){
		errChan <- server.StartServer()
	}()

	if err := <-errChan; err != nil {
		logger.Logger.Info("Server Failed To Start")
	}

	 

}
