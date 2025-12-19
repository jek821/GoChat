package Server

import (
	"GoChat/internal/logger"
	"GoChat/pkg"
	"net"
)

type Server struct {
	Ip string 
	Port string
	Listener *net.Listener
	ClientHandlers map[int32] chan pkg.Transmission
	ChannelIn chan pkg.Transmission
	IdCounter int32
	
}

func CreateServer(ip string, port string) *Server {
	return &Server{Ip: ip, Port: port, Listener: nil, ChannelIn: make(chan pkg.Transmission), IdCounter: 0}
}

func (S *Server) StartServer() error {
	listener, err := net.Listen("tcp", S.Ip + S.Port)
	if err != nil{
		return err
	}
	S.Listener = &listener
	logger.Logger.Info("Server Started",
											"IP", S.Ip,
											"PORT", S.Port,)

	defer logger.Logger.Info("Server Closed")

// Run Connection Acceptance in Loop 
	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}
		//Increment ID Counter 
		S.IdCounter++ 
		NewHandlerChan := CreateNewHandler(S.ChannelIn, S.IdCounter, conn)

		//Add New Handler ID and Chan to map
		S.ClientHandlers[S.IdCounter] = NewHandlerChan
		logger.Logger.Info("New Client Handler Spawned",
												"ID", S.IdCounter)

		

	return nil
}


	 
	


