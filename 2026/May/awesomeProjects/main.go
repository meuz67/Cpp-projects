package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync/atomic"
)

type payment struct {
	//Описание
	Description string `json:"description"`
	//Сумма
	USD int `json:"usd"`
}

var money = atomic.Int64{}
var paymentHistory = make([]payment, 0)

func payHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	var payment payment
	json.Unmarshal(httpRequestBody, &payment)
	if money.Load() >= int64(payment.USD) {
		money.Add(int64(-payment.USD))
		paymentHistory = append(paymentHistory, payment)
		return
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}
func paymentHistoryHandler(w http.ResponseWriter, r *http.Request) {
	for k, _ := range paymentHistory {
		var element = "Description:" + paymentHistory[k].Description + "; usd:" + strconv.Itoa(paymentHistory[k].USD) + "\n"
		w.Write([]byte(element))
	}
}
func main() {
	money.Add(1000)
	http.HandleFunc("/pay", payHandler)
	http.HandleFunc("/history", paymentHistoryHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err.Error())
	}
}
