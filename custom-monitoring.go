package main

import (
	"context"
	"log"

	monitoring "cloud.google.com/go/monitoring/apiv3"
	"google.golang.org/genproto/googleapis/monitoring/v3"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	ctx := context.Background()
	client, err := monitoring.NewMetricClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	// Crear una métrica personalizada
	req := &monitoring.CreateTimeSeriesRequest{
		Name: "projects/Practica docker",
		TimeSeries: []*monitoring.TimeSeries{
			{
				Metric: &monitoring.Metric{
					Type: "custom.googleapis.com/mi_metrica_personalizada",
				},
				Resource: &monitoring.MonitoredResource{
					Type: "global",
					Labels: map[string]string{
						"project_id": "steel-totality-430001-n0",
					},
				},
				Points: []*monitoring.Point{
					{
						Interval: &monitoring.TimeInterval{
							EndTime: timestamppb.Now(),
						},
						Value: &monitoring.TypedValue{
							Value: &monitoring.TypedValue_Int64Value{Int64Value: 123},
						},
					},
				},
			},
		},
	}

	if err := client.CreateTimeSeries(ctx, req); err != nil {
		log.Fatalf("Failed to create time series: %v", err)
	}
	log.Println("Métrica personalizada enviada correctamente")
}
