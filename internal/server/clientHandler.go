package Server

import (
	"GoChat/pkg"
	"net")

type clientHandler struct {
	Conn net.Conn
	Id int32
	inChan chan pkg.Transmission 
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


func (CH *clientHandler) RunHandler(){
// Run listen to in channel in loop
}
