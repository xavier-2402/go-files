package main

import (
    "context"
    "fmt"
    "log"

    compute "cloud.google.com/go/compute/apiv1"
    computepb "cloud.google.com/go/compute/apiv1/computepb"
    "google.golang.org/api/option"
)

// Funciones auxiliares para convertir valores a punteros
func stringPtr(s string) *string {
    return &s
}

func boolPtr(b bool) *bool {
    return &b
}

func int64Ptr(i int64) *int64 {
    return &i
}

func main() {
    ctx := context.Background()

    // Crea un cliente de Compute Engine
    c, err := compute.NewInstancesRESTClient(ctx, option.WithEndpoint("https://compute.googleapis.com/compute/v1/projects/"))
    if err != nil {
        log.Fatalf("Failed to create client: %v", err)
    }
    defer c.Close()

    projectID := "steel-totality-430001-n0"
    zone := "us-central1-a"

    // Crea una solicitud para insertar una instancia
    req := &computepb.InsertInstanceRequest{
        Project: projectID,
        Zone:    zone,
        InstanceResource: &computepb.Instance{
            Name: stringPtr("example-instance"),
            Disks: []*computepb.AttachedDisk{
                {
                    InitializeParams: &computepb.AttachedDiskInitializeParams{
                        SourceImage: stringPtr("projects/debian-cloud/global/images/family/debian-11"),
                    },
                    AutoDelete: boolPtr(true),
                    Boot:       boolPtr(true),
                },
            },
            MachineType: stringPtr(fmt.Sprintf("zones/%s/machineTypes/n1-standard-1", zone)),
            NetworkInterfaces: []*computepb.NetworkInterface{
                {
                    Name: stringPtr("global/networks/default"),
                },
            },
        },
    }

    // Llama a la API para insertar la instancia
    op, err := c.Insert(ctx, req)
    if err != nil {
        log.Fatalf("Failed to create instance: %v", err)
    }

    fmt.Printf("Instance creation started: %v\n", op)
}
