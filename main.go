package main

import (
	"fmt"
	pb "go-proto/pb"

	"google.golang.org/protobuf/proto"
)

func main() {
	Products := &pb.Products{
		Data: []*pb.Product{
			{
				Id:    1,
				Name:  "Adidas Black T-Shirt",
				Price: 100.00,
				Stock: 100,
				Category: &pb.Category{
					Id:   1,
					Name: "T-Shirt",
				},
			},
		},
	}

	data, err := proto.Marshal(Products)
	if err != nil {
		fmt.Println(err)
	}

	testProducts := &pb.Products{}
	if err = proto.Unmarshal(data, testProducts); err != nil {
		fmt.Println(err)
	}

	fmt.Println(testProducts)
}
