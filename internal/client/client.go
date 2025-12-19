package Client

import "net"

type Client struct {
	Ip   string
	Port string
	Conn *net.Conn
	Id int32
}


func 
