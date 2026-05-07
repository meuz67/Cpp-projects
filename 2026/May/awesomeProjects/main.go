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

func pay(w http.ResponseWriter, req *http.Request) {
	httpRequestBody, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	amount, err := strconv.Atoi(string(httpRequestBody))
	if err != nil {
		fmt.Println(err)
		return
	}
	if int64(amount) <= money.Load() {
		w.WriteHeader(http.StatusOK)
		money.Add(int64(-amount))
		fmt.Println("money = ", money.Load())
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}
func saveMoney(w http.ResponseWriter, req *http.Request) {
	httpRequestbody, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	amount, err := strconv.Atoi(string(httpRequestbody))
	if err != nil {
		fmt.Println(err)
		return
	}
	if int64(amount) <= money.Load() {
		money.Add(int64(-amount))
		bankMoney.Add(int64(amount))
		fmt.Println("money = ", money.Load(), "\n", "bankMoney = ", bankMoney.Load())
	} else {
		fmt.Println("Not enough money")
		w.WriteHeader(http.StatusBadRequest)
	}
}
func main() {
	money.Add(1000)
	http.HandleFunc("/pay", pay)
	http.HandleFunc("/saveMoney", saveMoney)
	http.ListenAndServe(":8080", nil)
}
