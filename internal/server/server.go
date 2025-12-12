package server

import (
	"net"
	"GoChat/internal/logger"
)

type Server struct {
	Ip string 
	Port string
	Conn *net.Listener
}

func CreateServer(ip string, port string) *Server {
	return &Server{Ip: ip, Port: port}
}

func (S Server) StartServer() error {
	conn, err := net.Listen("tcp", S.Ip + S.Port)
	if err != nil{
		return err
	}
	S.Conn = &conn
	logger.Logger.Info("Server Started",
											"IP", S.Ip,
											"PORT", S.Port,)
	return nil
}


