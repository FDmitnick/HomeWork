package main
 
import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "tGrpc/hello"
)
//访问服务的IP及端口，要与服务端的IP、端口对应
const address  ="127.0.0.1:8088"
/*
    1. 创建groc连接器
    2. 创建grpc客户端,并将连接器赋值给客户端
    3. 向grpc服务端发起请求
    4. 获取grpc服务端返回的结果
*/
func main() {
	// 创建一个grpc的连接
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err !=nil{
		fmt.Println(err)
		return
	}
	defer conn.Close()
	//创建grpc客户端
	c := pb.NewHelloMyServiceClient(conn)
	request:=&pb.HelloRequest{}
	request.Greeting="我是客户端，请求连接..."
	fmt.Println("开始请求服务...")
	//客户端向服务端发送请求，同时返回服务端的结果
	result, err := c.SayHello(context.Background(), p)
	if err !=nil{
		fmt.Println(err)
		return
	}
	fmt.Println(result.Reply)
}