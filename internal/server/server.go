package Server

import (
	"net"
	"GoChat/internal/logger"
)

type Server struct {
	Ip string 
	Port string
	Listener *net.Listener
	//ClientHandlers map[int] chan 
}

func CreateServer(ip string, port string) *Server {
	return &Server{Ip: ip, Port: port}
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

	// Run in Loop
	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}
		go handleConn(conn)
	}	


	return nil
}

// Each New Client Connection will be assigned a unique ID
// We then run another GoRoutine which is the client handler 
// Each client handler is added to our Map of Client Handlers 
// This Map contains the ID of each handler and its associated Channel to receive input
func (S Server) handleClient(){

	 
	


	}

