package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)
var ack2 string
type API2 int
var k int

func main(){
	api := new(API2)
	err := rpc.Register(api)
	if err != nil {
		log.Fatal("error registering API", err)
	}

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":1802")

	if err != nil {
		log.Fatal("Listener error", err)
	}
	log.Printf("serving rpc on port %d", 1802)
	http.Serve(listener, nil)

	if err != nil {
		log.Fatal("error serving: ", err)
	}


}

func (a *API2) GetACK(request string, reply *string) error {
	if request =="request"{
		*reply="NO"
	}else if request =="prepare"{
		*reply="YES"
	}else if request =="commit"{
		*reply="YES"
	}

	return nil

}