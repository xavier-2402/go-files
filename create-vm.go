package main

import (
	"context"
	"fmt"
	"log"

	compute "cloud.google.com/go/compute/apiv1"
	"google.golang.org/genproto/googleapis/cloud/compute/v1"
)

func main() {
	ctx := context.Background()
	instancesClient, err := compute.NewInstancesRESTClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer instancesClient.Close()

	req := &compute.InsertInstanceRequest{
		Project: "tu-proyecto",
		Zone:    "us-central1-a",
		InstanceResource: &compute.Instance{
			Name:        "mi-nueva-vm",
			MachineType: "zones/us-central1-a/machineTypes/n1-standard-1",
			Disks: []*compute.AttachedDisk{
				{
					InitializeParams: &compute.AttachedDiskInitializeParams{
						SourceImage: "projects/debian-cloud/global/images/family/debian-10",
					},
					Boot: true,
				},
			},
			NetworkInterfaces: []*compute.NetworkInterface{
				{
					Name: "global/networks/default",
				},
			},
		},
	}

	op, err := instancesClient.Insert(ctx, req)
	if err != nil {
		log.Fatalf("Failed to create instance: %v", err)
	}
	fmt.Printf("Instance creation started: %v\n", op)
}
