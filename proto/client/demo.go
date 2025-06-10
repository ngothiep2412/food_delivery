package main

import (
	"fmt"
	"g05-food-delivery/proto/pb/demo"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	fmt.Println("Hello Client")

	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer cc.Close()

	client := demo.NewHelloServiceClient(cc)
	request := &demo.HelloRequest{Name: "Thiep"}

	resp, _ := client.Hello(context.Background(), request)
	fmt.Println("Receive response => [%v]", resp.Message)

	result, _ := client.GetRestaurantLikeStat(context.Background(), &demo.RestaurantLikeStatRequest{
		ResIds: []int32{1, 2, 3},
	})
	fmt.Printf("%+v\n", result.Result)

	streamClient, err := client.HelloStream(context.Background(), request)

	for i := 1; i <= 10; i++ {
		time.Sleep(time.Second)
		res, _ := streamClient.Recv()
		fmt.Println("Receive response => [%v]", res.Message)
	}

}
