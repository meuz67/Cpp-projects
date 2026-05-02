package postman

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Postman(ctx context.Context, wg *sync.WaitGroup, transferPoint chan<- string, n int, mail string) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Number ", n, " done")
			return
		default:
			fmt.Println("Postman number ", n, " taken mail")
			time.Sleep(1 * time.Second)
			fmt.Println("Postman number ", n, " brought mail ", mail)
			transferPoint <- mail
			fmt.Println("Postman number ", n, " brought mail ", mail)
		}
	}
}
func PostmanPool(ctx context.Context, postmanCount int) <-chan string {
	mailTransferPoint := make(chan string)
	wg := &sync.WaitGroup{}
	for i := 1; i <= postmanCount; i++ {
		wg.Add(1)
		go Postman(ctx, wg, mailTransferPoint, i, postmanToMail(i))
	}
	go func() {
		wg.Wait()
		close(mailTransferPoint)
	}()
	return mailTransferPoint
}
func postmanToMail(n int) string {
	ptm := map[int]string{
		1: "Hello ",
		2: "World",
		3: "Invite on Epstein Island",
		4: "Information",
	}
	mail, ok := ptm[n]
	if !ok {
		return "Lottery"
	}
	return mail
}
