package main

import "fmt"

type Server interface {
	HandleRequest(int)
}

type RealServer struct {
}

func (r *RealServer) HandleRequest(i int) {
	fmt.Println("receive request", i)
}

type ProxyServer struct {
	Server
	realServer Server
}

func (r *ProxyServer) HandleRequest(i int) {
	fmt.Println("sleep", i)
	r.realServer.HandleRequest(i)
}

func main(){
	s:=&ProxyServer{realServer: &RealServer{}}
	s.HandleRequest(1)
	s.HandleRequest(2)
	s.HandleRequest(3)
}
