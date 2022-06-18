package main

import (
	"context"
	"fmt"
	"time"
)

//var wg sync.WaitGroup

func f(ctx context.Context) {
	//defer wg.Done()
LOOP:
	for{
		fmt.Println("hello world!")
		time.Sleep(time.Millisecond*500)
		select {
			case <- ctx.Done():
				break LOOP
			default:
		}
	}

}

func main () {
	ctx, cancel := context.WithCancel(context.Background())
	//wg.Add(1)
	go f(ctx)
	time.Sleep(time.Second*3)
	cancel()
}
