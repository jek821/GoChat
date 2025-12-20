package Client

import (
	"net"
	"GoChat/internal/logger"
)

type Client struct {
	Ip   string
	Port string
	Conn *net.Conn
	Id int32
}


func CreateClient(ip string, port string) *Client {

	return &Client{Ip: ip, Port: port, Conn: nil, Id: -1}

} 

func (C *Client) StartClient(){
	conn, err := net.Dial("tcp", C.Ip + C.Port)
	if err != nil{
		logger.Logger.Info("Client Failed To Start",
												"IP", C.Ip,
												"PORT", C.Port)
	}
	C.Conn = &conn 
	logger.Logger.Info("Client Started",
											"IP", C.Ip,
											"Port", C.Port)

}
