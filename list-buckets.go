package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/storage"
)

func main() {
	ctx := context.Background()

	// Crear un cliente de Google Cloud Storage
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	// Listar todos los buckets en el proyecto
	it := client.Buckets(ctx, "steel-totality-430001-n0")
	for {
		bucketAttrs, err := it.Next()
		if err != nil {
			log.Fatalf("Bucket listing error: %v", err)
		}
		fmt.Println(bucketAttrs.Name)
	}
}
