package main

import (
	"fmt"
	"log"
	"net/rpc"

)


func main() {
	var reply1 string
	var reply2 string
	var reply3 string

	var client1 *rpc.Client
	var client2 *rpc.Client
	var client3 *rpc.Client

	client1, err := rpc.DialHTTP("tcp", "127.0.0.1:1801")

	if err != nil {
		log.Fatal("Connection error: ", err)
	}

	client2, err = rpc.DialHTTP("tcp", "127.0.0.1:1802")

	if err != nil {
		log.Fatal("Connection error: ", err)
	}

	client3, err = rpc.DialHTTP("tcp", "127.0.0.1:1803")

	if err != nil {
		log.Fatal("Connection error: ", err)
	}



	client1.Call("API1.GetACK", "request", &reply1)
	//fmt.Println(reply1)
	client1.Call("API1.GetACK", "request", &reply2)
	//fmt.Println(reply1)
	client2.Call("API2.GetACK", "request", &reply3)
	//fmt.Println(reply2)


	client3.Call("API.GetACK3", "prepare", &reply1)
	//fmt.Println(reply3)
	client1.Call("API.GetACK1", "prepare", &reply2)
	//fmt.Println(reply1)
	client2.Call("API.GetACK2", "prepare", &reply3)
	//fmt.Println(reply2)


	client3.Call("API.GetACK3", "commit", &reply1)
	//fmt.Println(reply3)
	client2.Call("API.GetACK2", "commit", &reply2)
	//fmt.Println(reply2)
	client3.Call("API.GetACK3", "commit", &reply3)
	//fmt.Println(reply3)

	if reply1 == "NO" || reply2 == "NO" || reply3 == "NO"{
		fmt.Println("GLOBAL ABORT")
		client2.Call("API2.GetACK", "request", &reply3)
		fmt.Println("Request (Ready to vote)",reply1)
		client2.Call("API.GetACK2", "prepare", &reply3)
		fmt.Println("Prepare to Commit",reply2)
		client3.Call("API.GetACK3", "commit", &reply3)
		fmt.Println("Commit Initialized",reply3)

	}else{
		fmt.Println("GLOBAL COMMIT")
		fmt.Println("Request (Ready to vote)",reply1)
		client3.Call("API.GetACK3", "commit", &reply3)
		fmt.Println("Prepare to Commit",reply2)
		client2.Call("API.GetACK2", "prepare", &reply3)
		fmt.Println("Commit Initialized",reply3)
		client3.Call("API.GetACK3", "commit", &reply3)
	}



}