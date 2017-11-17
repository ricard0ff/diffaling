package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"cloud.google.com/go/storage"
	"golang.org/x/net/context"
	"google.golang.org/api/iterator"
)

func main() {
	ctx := context.Background()

	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
	if projectID == "" {
		fmt.Fprintf(os.Stderr, "GOOGLE_CLOUD_PROJECT environment variable must be set.\n")
		os.Exit(1)
	}

	// [START setup]
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	// [END setup]

	// Give the bucket a unique name.
	name := fmt.Sprintf("golang-example-buckets-%d", time.Now().Unix())
	if err := create(client, projectID, name); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("created bucket: %v\n", name)

	// list buckets from the project
	buckets, err := list(client, projectID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("buckets: %+v\n", buckets)

}

func create(client *storage.Client, projectID, bucketName string) error {
	ctx := context.Background()
	// [START create_bucket]
	if err := client.Bucket(bucketName).Create(ctx, projectID, nil); err != nil {
		return err
	}
	// [END create_bucket]
	return nil
}
func list(client *storage.Client, projectID string) ([]string, error) {
	ctx := context.Background()
	// [START list_buckets]
	var buckets []string
	it := client.Buckets(ctx, projectID)
	for {
		battrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		buckets = append(buckets, battrs.Name)
	}
	// [END list_buckets]
	return buckets, nil
}
