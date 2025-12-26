package Server

import (
	"GoChat/pkg"
	"net"
	"GoChat/internal/logger"
	"fmt")

type clientHandler struct {
	Conn net.Conn
	Id int32
	InChan chan pkg.Transmission 
	outChan chan pkg.Transmission
}

//Server will give a channel to communicate with it and the new Conn
//Server will also give Handler ID (Will be same as Client ID )
//We will return our channel
// I think we should also make the constructor start the handler routine.
func CreateNewHandler(servChan chan pkg.Transmission, id int32, conn net.Conn) chan pkg.Transmission {

	handlerChan := make(chan pkg.Transmission)
	return handlerChan

}

func (CH *clientHandler) RunAllReceivers() {

}


func (CH *clientHandler) RunServerReceiver() {
// Run listen to in channel in loop
	for t := range CH.InChan {
		HandleTransmission(&t, CH.Id)
	}
}

func (CH *clientHandler) RunClientReceiver() {

	 
}


func HandleTransmission(T *pkg.Transmission, clientId int32) {
	clientIdString := fmt.Sprintf("%d", clientId)
	switch T.Code {
	case pkg.TestDataCode:
		//Logic for Handling TestData HandleTransmission()
		logger.Logger.Info("Test Data Received","Handler ID: ", 
												clientIdString,)	
	default:
		logger.Logger.Error("Unrecognized Transmission type received","Handler ID: ", clientIdString,)
	}

}
