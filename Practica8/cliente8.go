package main

import(
	"fmt"
	"net/rpc"
)

func client(){
	c, err := rpc.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	paso := []string{}
	
	err = c.Call("Server.AddScore",paso, &paso)

	fmt.Println(paso);
}

func main() {
	client()
}