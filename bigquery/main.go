package main

import (
	"cloud.google.com/go/bigquery"
	"context"
	"flag"
	"log"
	"os"
)

const (
	projectID = "smedia-business-intelligence"
	datasetID = "octopus_dev"
	tableID   = "test"
)

func main() {
	// gcloud --project smedia-business-intelligence iam service-accounts keys create OUTPUT-FILE /tmp/key.json --iam-account octopus-dev@smedia-business-intelligence.iam.gserviceaccount.com
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "/tmp/key.json")

	ctx := context.Background()

	flag.Parse()

	if err := run(ctx); err != nil {
		log.Fatalf("failed: %v\n", err)
	}
	log.Printf("finished\n")
}

func run(ctx context.Context) error {
	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		return err
	}
	defer client.Close()

	query := client.Query("SELECT orderNumber FROM `smedia-business-intelligence.octopus_dev.atlassian_invoice_v1` LIMIT 20")
	query.Dst = client.Dataset(datasetID).Table(tableID)
	query.WriteDisposition = bigquery.WriteTruncate
	query.CreateDisposition = bigquery.CreateIfNeeded

	job, err := query.Run(ctx)
	if err != nil {
		return err
	}
	log.Printf("job ID: %s", job.ID())

	status, err := job.Wait(ctx)
	if err != nil {
		return err
	}
	log.Printf("wait finish with state %s", stateToString(*status))
	return nil
}

func stateToString(status bigquery.JobStatus) string {
	switch status.State {
	case bigquery.Done:
		return "Done"
	case bigquery.Running:
		return "Running"
	case bigquery.Pending:
		return "Pending"
	case bigquery.StateUnspecified:
		return "StateUnspecified"
	default:
		return "unkown"
	}
}
