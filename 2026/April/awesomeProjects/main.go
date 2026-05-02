package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
)

var mtx = sync.Mutex{}
var money = 1000
var bankMoney = 0

func payHandler(w http.ResponseWriter, r *http.Request) {
	httpRequstBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)
	}
	amount, err := strconv.ParseInt(string(httpRequstBody), 10, 64)
	if err != nil {
		fmt.Println("Failed to parse amount")
	}
	mtx.Lock()
	if amount <= int64(money) {
		money -= int(amount)
	}
	mtx.Unlock()
	fmt.Println("Money:", money)
}
func saveHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)
	}
	amount, err := strconv.ParseInt(string(httpRequestBody), 10, 64)
	if err != nil {
		fmt.Println("Failed to parse amount")
	}
	mtx.Lock()
	if amount <= int64(money) {
		money -= int(amount)
		bankMoney += int(amount)
	}
	mtx.Unlock()
	fmt.Println("Money:", money)
	fmt.Println("Bank money:", bankMoney)
}
func main() {
	http.HandleFunc("/", payHandler)
	http.HandleFunc("/save", saveHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Failed to start server")
	}

}
