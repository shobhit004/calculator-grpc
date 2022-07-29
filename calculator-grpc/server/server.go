package main

import (
	"context"
	"log"
	"math"
	"net"
	"time"

	"github.com/Shobhit0403/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct {
	calculatorpb.UnimplementedCalculatorServiceServer
}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (resp *calculatorpb.SumResponse, err error) {
	num1 := req.GetNum1()
	num2 := req.GetNum2()

	resp = &calculatorpb.SumResponse{
		Sum: num1 + num2,
	}
	return resp, nil
}

func (*server) PrimeNumbers(req *calculatorpb.PrimeNumbersRequest, resp calculatorpb.CalculatorService_PrimeNumbersServer) error {

	isPrime := func(num int64) bool {
		if num <= 1 {
			return false
		}
		limit := int64(math.Sqrt(float64(num)))
		for i := int64(2); i <= limit; i++ {
			if num%i == 0 {
				return false
			}
		}
		return true
	}

	limit := req.GetLimit()

	for i := int64(0); i <= limit; i++ {
		if isPrime(i) {
			res := calculatorpb.PrimeNumbersResponse{
				PrimeNum: i,
			}
			time.Sleep(1000 * time.Millisecond)
			resp.Send(&res)
		}
	}
	return nil
}

func main() {

	listen, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatalf("Failed to Listen: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	log.Println("Initiating Server")
	if err = s.Serve(listen); err != nil {
		log.Fatalf("failed to serve : %v", err)
	}
}
