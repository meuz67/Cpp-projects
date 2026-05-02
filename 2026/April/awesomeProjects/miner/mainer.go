package miner

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Miner(ctx context.Context, wg *sync.WaitGroup, transferPoint chan<- int, n int, power int) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("miner ", n, " exit")
			return
		default:
			fmt.Println("Miner number ", n, " started mining")
			time.Sleep(1 * time.Second)
			fmt.Println("Miner number ", n, " finished mining ", power)
			transferPoint <- power
			fmt.Println("Miner number ", n, " had given ", power, " coal")
		}
	}
}
func MinerPool(ctx context.Context, mineCount int) <-chan int {
	wg := &sync.WaitGroup{}
	coalTransferPoint := make(chan int)
	for i := 1; i < mineCount; i++ {
		wg.Add(1)
		go Miner(ctx, wg, coalTransferPoint, i, i*10)
	}
	go func() {
		wg.Wait()
		close(coalTransferPoint)
	}()
	return coalTransferPoint
}
