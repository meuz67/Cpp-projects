package main

import (
	"awesomeProject/http_server"
	"fmt"
)

func main() {
	fmt.Println("started http server")
	err := http_server.StartHttpServer()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("started http server")
	}
}
