package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"myGRPCDemo/helloWorld/proto"
	"strconv"
	"time"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:50052", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)
var c proto.SimpleClient

func main() {
	flag.Parse()
	//从输入的证书文件中为客户端构造TLS凭证

	//cert, _ := tls.LoadX509KeyPair("helloWorld/tls/client.pem", "helloWorld/tls/client.key")
	//certPool := x509.NewCertPool()
	//ca, _ := ioutil.ReadFile("helloWorld/tls/ca.pem")
	//certPool.AppendCertsFromPEM(ca)
	//
	//creds := credentials.NewTLS(&tls.Config{
	//	Certificates: []tls.Certificate{cert},
	//	ServerName:   "localhost",
	//	RootCAs:      certPool,
	//})

	// 连接服务器
	//conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(creds))
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c = proto.NewSimpleClient(conn)

	route()
	conversations()

}
func route() {

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Route(ctx, &proto.SimpleRequest{Data: "hello world"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r)
}

// conversations 调用服务端的Conversations方法
func conversations() {
	//调用服务端的Conversations方法，获取流
	stream, err := c.Conversations(context.Background())
	if err != nil {
		log.Fatalf("get conversations stream err: %v", err)
	}
	for n := 0; n < 5; n++ {
		err := stream.Send(&proto.StreamRequest{Question: "stream client rpc " + strconv.Itoa(n)})
		if err != nil {
			log.Fatalf("stream request err: %v", err)
		}
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Conversations get stream err: %v", err)
		}
		// 打印返回值
		log.Println(res.Answer)
	}
	//最后关闭流
	err = stream.CloseSend()
	if err != nil {
		log.Fatalf("Conversations close stream err: %v", err)
	}
}
