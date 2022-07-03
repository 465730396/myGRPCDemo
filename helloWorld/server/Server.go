package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"myGRPCDemo/helloWorld/proto"
	"net"
	"strconv"
)

var (
	port = flag.Int("port", 50052, "The server port")
)

type server struct {
	proto.UnimplementedSimpleServer
}

func (*server) Route(context context.Context, request *proto.SimpleRequest) (*proto.SimpleResponse, error) {

	return &proto.SimpleResponse{Code: 11, Value: "hello alex"}, nil
}
func (*server) Conversations(srv proto.Simple_ConversationsServer) error {
	n := 1
	for {
		req, err := srv.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		err = srv.Send(&proto.StreamResponse{
			Answer: "from stream server answer: the " + strconv.Itoa(n) + " question is " + req.Question,
		})
		if err != nil {
			return err
		}
		n++
		log.Printf("from stream client question: %s", req.Question)
	}
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// 从输入证书文件和密钥文件为服务端构造TLS凭证
	//cert, _ := tls.LoadX509KeyPair("helloWorld/tls/server.pem", "helloWorld/tls/server.key")
	//certPool := x509.NewCertPool()
	//ca, _ := ioutil.ReadFile(".helloWorld/tls/ca.pem")
	//certPool.AppendCertsFromPEM(ca)
	//
	//creds := credentials.NewTLS(&tls.Config{
	//	Certificates: []tls.Certificate{cert},
	//	ClientAuth:   tls.RequireAndVerifyClientCert,
	//	ClientCAs:    certPool,
	//})
	//
	//// 新建gRPC服务器实例,并开启TLS认证
	//s := grpc.NewServer(grpc.Creds(creds))
	s := grpc.NewServer()
	proto.RegisterSimpleServer(s, &server{})
	log.Printf("server whth TLS and token... listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
