package main

import (
	"awesomeProject/httpServer"
	"fmt"
)

func main() {
	fmt.Println("Http server started")
	err := httpServer.StartHttpServer()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("httpServer.StartHttpServer() done")
	}
}
