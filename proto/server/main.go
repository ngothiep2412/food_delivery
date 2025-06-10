package main

import (
	"fmt"
	"g05-food-delivery/proto/pb/demo"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"
)

type server struct {
	demo.UnimplementedHelloServiceServer
}

func (*server) Hello(ctx context.Context, request *demo.HelloRequest) (*demo.HelloResponse, error) {
	name := request.Name
	response := &demo.HelloResponse{
		Message: "Hello " + name,
	}

	return response, nil
}

func (*server) GetRestaurantLikeStat(ctx context.Context, request *demo.RestaurantLikeStatRequest) (*demo.RestaurantLikeStatResponse, error) {
	request.ResIds = []int32{1, 2, 3}
	// call db
	return &demo.RestaurantLikeStatResponse{
		Result: map[int32]int32{
			1: 3,
			2: 4,
		},
	}, nil
}

func (*server) HelloStream(req *demo.HelloRequest, serv demo.HelloService_HelloStreamServer) error {
	for i := 1; i <= 10; i++ {
		fmt.Println("Waiting client ready ...")
		_ = serv.Send(&demo.HelloResponse{
			Message: "Hello " + fmt.Sprintf("%d", i),
		})
	}

	return nil
}

func main() {
	address := "0.0.0.0:50051"
	lis, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	fmt.Printf("server listening at %v\n", lis.Addr())

	s := grpc.NewServer()

	demo.RegisterHelloServiceServer(s, &server{})

	//s.Serve(lis)

	log.Println("Serving gRPC on 0.0.0.0:50051")
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	// http
	conn, err := grpc.NewClient(
		"0.0.0.0:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()

	err = demo.RegisterHelloServiceHandler(context.Background(), gwmux, conn)

	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: gwmux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
	log.Fatalln(gwServer.ListenAndServe())
}
