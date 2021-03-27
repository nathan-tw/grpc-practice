package main

import (
	"context"
	"fmt"
	"net"

	calculatorPB "github.com/nathan-tw/grpc-practice/proto/calculator"
	"google.golang.org/grpc"
)

type Server struct{
	calculatorPB.UnimplementedCalculatorServiceServer
}

func (*Server) Sum(ctx context.Context, req *calculatorPB.CalculatorRequest) (*calculatorPB.CalculatorResponse, error) {
	fmt.Printf("Sum function is invoked with %v \n", req)

	a := req.GetA()
	b := req.GetB()

	res := &calculatorPB.CalculatorResponse{
		Result: a + b,
	}

	return res, nil

}

func main() {
	fmt.Println("starting gRPC server")
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	calculatorPB.RegisterCalculatorServiceServer(grpcServer, &Server{})

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}

}
