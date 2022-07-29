package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/Shobhit0403/calculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {

	cc, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)

	//Unary API function - Calculator
	Sum(c)

	// Server-Side Streaming function - PrimeNumbers
	PrimeNumbers(c)

	fmt.Println("Completed.")

}

func Sum(c calculatorpb.CalculatorServiceClient) {

	req := calculatorpb.SumRequest{
		Num1: 5.5,
		Num2: 7.4,
	}

	resp, err := c.Sum(context.Background(), &req)
	if err != nil {
		log.Fatalf("error while calling sum grpc unary call: %v", err)
	}

	log.Printf("Response from Unary Call, Sum : %v", resp.GetSum())

}

func PrimeNumbers(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Staring ServerSide GRPC streaming ....")

	req := calculatorpb.PrimeNumbersRequest{
		Limit: 15,
	}

	respStream, err := c.PrimeNumbers(context.Background(), &req)
	if err != nil {
		log.Fatalf("error while calling Prime Numbers server-side streaming grpc : %v", err)
	}

	for {
		msg, err := respStream.Recv()
		if err == io.EOF {
			//we have reached to the end of the file
			break
		}

		if err != nil {
			log.Fatalf("error while receving server stream : %v", err)
		}

		log.Println("Response From Server, Prime Number : ", msg.GetPrimeNum())
	}
}
