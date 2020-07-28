package main

import (
	"context"
	"fmt"
	"log"
	"reflect"

	"github.com/the-go/blog/blogpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Printf("Blog Client \n")

	opts := grpc.WithInsecure()

	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := blogpb.NewBlogServiceClient(cc)

	doUnaryCreate(c)

}

func doUnaryCreate(c blogpb.BlogServiceClient) {
	fmt.Printf("starting unary function \n \n")
	blog := &blogpb.Blog{
		AuthorId: "Muhammmad",
		Title:    "Yes Title",
		Content:  "The Lorem Dummu Dummy",
	}
	createdBlog, err := c.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{Blog: blog})
	checkTypeOf := reflect.TypeOf(createdBlog.GetBlog().GetId())
	log.Printf("typeof created blog id: ===> %v \n\n", checkTypeOf)
	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}

	log.Println("Blog has been created: %v \n\n", createdBlog)

	doUnaryReadFailed(c)
	doUnaryReadSuccess(c, createdBlog.GetBlog().GetId())
}

func doUnaryReadFailed(c blogpb.BlogServiceClient) {
	fmt.Println("Reading the blog client failed \n\n")

	request := &blogpb.ReadBlogRequest{BlogId: "12334113"}
	_, err := c.ReadBlog(context.Background(), request)
	if err != nil {
		fmt.Printf("Error happened while reading, %v\n\n", err)
	}
}

func doUnaryReadSuccess(c blogpb.BlogServiceClient, blogID string) {
	fmt.Printf("reading the blog client success")

	request := &blogpb.ReadBlogRequest{BlogId: blogID}
	response, err := c.ReadBlog(context.Background(), request)
	if err != nil {
		fmt.Printf("Error happened while reading, %v\n", err)
	}

	fmt.Printf(" FOUND IT, %v", response)
}
