package main
 
import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "tGrpc/hello"
	"net"
)
 
const (
	port = ":8088"
)
 
type server struct {
}
 
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	result := &pb.HelloResponse{}
	fmt.Println(in.Greeting)
	result.Reply = "这里是服务器，请求服务成功！"
	return result, nil
 
}
/*
    1. 首先我们必须实现我们自定义rpc服务,例如:rpc SayHello()-在此我们可以实现我们自己的逻辑
    2. 创建监听listener
    3. 创建grpc的服务
    4. 将我们的服务注册到grpc的server中
    5. 启动grpc服务,将我们自定义的监听信息传递给grpc客户端
*/
func main() {
	//创建server端监听端口
	listener, err := net.Listen("tcp", port)
	if err!=nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()
	myServer:=grpc.NewServer()//创建grpc的service
	pb.RegisterHelloMyServiceServer(myServer,&server{})//注册服务
	fmt.Println("启动grpc服务...")
	myServer.Serve(listener)//启动监听服务
}