package main

import (
	"awesomeprojects/miner"
	"awesomeprojects/postman"
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	closed := sync.RWMutex{}
	var coal int
	var mails []string
	minerContext, minerCancel := context.WithCancel(context.Background())
	postmanContext, postmanCancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(3 * time.Second)
		minerCancel()
	}()
	go func() {
		time.Sleep(6 * time.Second)
		postmanCancel()
	}()
	coalTransferPoint := miner.MinerPool(minerContext, 10)
	mailTransferPoint := postman.PostmanPool(postmanContext, 10)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		closed.Lock()
		defer wg.Done()
		for v := range coalTransferPoint {
			coal += v
		}
		closed.Unlock()
	}()
	wg.Add(1)
	go func() {
		closed.Lock()
		defer wg.Done()
		for v := range mailTransferPoint {
			mails = append(mails, v)
		}
		closed.Unlock()
	}()
	wg.Wait()
	closed.RLock()
	fmt.Println("coal:", coal)
	fmt.Println("mails:", len(mails))
	closed.RUnlock()
}
