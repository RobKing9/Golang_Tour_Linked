package main

import (
	"fmt"
	"net"
	"test1/pb/person"

	"google.golang.org/grpc"
)

type PersonServer struct {
	person.UnimplementedSearchServiceServer
}

func main() {
	listen, _ := net.Listen("tcp", ":9999")
	fmt.Println("服务器启动成功！！！\n正在监听9999端口！！！")
	ser := grpc.NewServer()
	person.RegisterSearchServiceServer(ser, &PersonServer{})
	ser.Serve(listen)
}

/*
func (*PersonServer) Search(c context.Context, req *person.PersonReq) (*person.PersonRes, error) {
	name := req.GetName()
	res := &person.PersonRes{Name: "你的名字是"+name}
	return res, nil
}

*/

//服务端流式 RPC
//服务端流式 RPC 是一个单向流，指 Server 为 Stream,,Client 为普通的一元 RPC 请求。
//简单来讲，就是客户端发起一次普通的 RPC 请求，服务端通过流式响应多次发送数据集，客户端 Recv 接收数据集

func (*PersonServer) SearchIn(server person.SearchService_SearchInServer) error {
	for {
		req, err := server.Recv()
		fmt.Println(req)
		if err != nil {
			server.SendAndClose(&person.PersonRes{Name: "完成服务！"})
			break
		}
	}
	return nil
}

/*
func (*PersonServer) SearchOut(req *person.PersonReq, server person.SearchService_SearchOutServer) error {
	name := req.Name
	i := 0
	for {
		if i > 10 {
			break
		}
		time.Sleep(time.Second)
		server.Send(&person.PersonRes{Name: "你的名字是"+name+"i"})
		i++
	}
	return nil
}
*/

/*
func (*PersonServer) SearchIO(server person.SearchService_SearchIOServer) error {
	chanstr  := make(chan string)
	i := 0
	//接受信息
	go func() {
		for {
			i++
			//不断接收消息
			req, _ := server.Recv()
			if i>10{
				chanstr<-"服务完成"
				break
			}
			chanstr<-req.Name
		}
	}()
	//将接收的信息返回
	for {
		s := <-chanstr
		s1 := "你发过来的消息是:"+s+ "\n接收成功！"
		if s == "服务完成" {
			break
		}
		//将不断接收的消息 返回
		server.Send(&person.PersonRes{Name: s1})
	}
	return nil
}
*/
