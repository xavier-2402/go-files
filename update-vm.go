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

	// Detener la VM
	stopReq := &compute.StopInstanceRequest{
		Project:  "steel-totality-430001-n0",
		Zone:     "us-central1-a",
		Instance: "instance-20240902-035008",
	}
	_, err = instancesClient.Stop(ctx, stopReq)
	if err != nil {
		log.Fatalf("Failed to stop instance: %v", err)
	}

	fmt.Println("Instancia detenida. Actualizando...")

	// Actualizar la configuración de la VM (cambiar el tipo de máquina)
	setMachineReq := &compute.SetMachineTypeInstanceRequest{
		Project:  "steel-totality-430001-n0",
		Zone:     "us-central1-a",
		Instance: "instance-20240902-035008",
		InstancesSetMachineTypeRequestResource: &compute.InstancesSetMachineTypeRequest{
			MachineType: "zones/us-central1-a/machineTypes/n1-standard-2",
		},
	}
	_, err = instancesClient.SetMachineType(ctx, setMachineReq)
	if err != nil {
		log.Fatalf("Failed to update instance: %v", err)
	}

	// Reiniciar la VM
	startReq := &compute.StartInstanceRequest{
		Project:  "steel-totality-430001-n0",
		Zone:     "us-central1-a",
		Instance: "instance-20240902-035008",
	}
	_, err = instancesClient.Start(ctx, startReq)
	if err != nil {
		log.Fatalf("Failed to start instance: %v", err)
	}

	fmt.Println("Instancia actualizada y reiniciada.")
}
