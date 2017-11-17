package main

import "cloud.google.com/go/storage"
import (
	"context"
	"fmt"
)

func main() {
	fmt.Println("hello world")
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}
