package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)
var ack3 string
type API3 int
var k int

func main(){
	api := new(API3)
	err := rpc.Register(api)
	if err != nil {
		log.Fatal("error registering API", err)
	}

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":1803")

	if err != nil {
		log.Fatal("Listener error", err)
	}
	log.Printf("serving rpc on port %d", 1803)
	http.Serve(listener, nil)

	if err != nil {
		log.Fatal("error serving: ", err)
	}


}

func (a *API3) GetACK(request string, reply *string) error {
	if request =="request"{
		*reply="YES"
	}else if request =="prepare"{
		*reply="YES"
	}else if request =="commit"{
		*reply="YES"
	}
	
	return nil

}