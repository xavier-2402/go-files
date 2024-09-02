package main

import (
	"context"
	"log"

	"cloud.google.com/go/logging"
)

func main() {
	ctx := context.Background()

	// Crear un cliente para Cloud Logging
	client, err := logging.NewClient(ctx, "steel-totality-430001-n0")
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	// Crear un logger
	logger := client.Logger("mi-logger")

	// Enviar un mensaje de log
	logger.Log(logging.Entry{
		Payload:  "Este es un mensaje de log desde Go",
		Severity: logging.Info,
	})

	// Forzar a que se envíen los logs
	client.Flush()
}
