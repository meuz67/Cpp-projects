package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync/atomic"
)

var money = atomic.Int64{}
var bankMoney = atomic.Int64{}

func payHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestBodyString, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Failed to read request body")
		http.Error(w, err.Error(), 400)
		return
	}
	amount, err := strconv.Atoi(string(httpRequestBodyString))
	if err != nil {
		fmt.Println("Failed to parse amount")
		http.Error(w, err.Error(), 400)
		return
	}
	if int64(amount) <= money.Load() {
		money.Add(int64(-amount))
	}
	fmt.Println("amount:", money.Load())
	w.WriteHeader(http.StatusOK)
}
func saveHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestBodyString, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Failed to read request body")
		http.Error(w, err.Error(), 400)
		return
	}
	amount, err := strconv.Atoi(string(httpRequestBodyString))
	if err != nil {
		fmt.Println("Failed to parse amount")
		http.Error(w, err.Error(), 400)
		return
	}
	if int64(amount) <= money.Load() {
		bankMoney.Add(int64(amount))
		money.Add(int64(-amount))
	}
	w.WriteHeader(http.StatusOK)
	fmt.Println("bank money:", bankMoney.Load())
	fmt.Println("money:", money.Load())
}
func main() {
	money.Add(1000)
	http.HandleFunc("/pay", payHandler)
	http.HandleFunc("/save", saveHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {

	}
}
