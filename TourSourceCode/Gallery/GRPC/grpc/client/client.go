package main

import (
	"context"
	"fmt"
	"test1/pb/person"
	"time"

	"google.golang.org/grpc"
)

func main() {
	con, _ := grpc.Dial("localhost:9999", grpc.WithInsecure())
	client := person.NewSearchServiceClient(con)

	/*
		res, err := client.Search(context.Background(), &person.PersonReq{Name: "zxw"})
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(res)

	*/

	c, _ := client.SearchIn(context.Background())
	i := 0
	for {
		if i > 10 {
			res, _ := c.CloseAndRecv()
			fmt.Println(res)
			break
		}
		time.Sleep(time.Second)
		c.SendMsg(&person.PersonReq{Name: "我是进来的消息"})
		i++
	}

	/*
		c, _ := client.SearchOut(context.Background(), &person.PersonReq{Name: "zxw"})
		for {
			req, err := c.Recv()
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println(req)
		}
	*/

	/*
		c, _ := client.SearchIO(context.Background())
		wg := sync.WaitGroup{}
		wg.Add(2)
		//发送消息
		go func() {
			for {
				err := c.Send(&person.PersonReq{Name: "zxw"})
				if err != nil {
					wg.Done()
					break
				}
			}
		}()
		//接收消息
		go func() {
			for {
				req, err := c.Recv()
				if err != nil {
					fmt.Println(err.Error())
					wg.Done()
					break
				}
				fmt.Println(req)
			}
		}()
		wg.Wait()

	*/

}
