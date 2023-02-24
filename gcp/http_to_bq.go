package telemetry

import (
	"cloud.google.com/go/bigquery"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type Item struct {
	Timestamp   time.Time `json:"timestamp"`
	DeviceID    string    `json:"device_id"`
	MemoryUsage float64   `json:"memory_usage"`
	CPUUsage    float64   `json:"cpu_usage"`
}

func StoreToBigQuery(w http.ResponseWriter, r *http.Request) {
	// Extracting data from the HTTP POST request

	var projectID = os.Getenv("PROJECT_ID")

	var requestData struct {
		Timestamp   time.Time `json:"timestamp"`
		DeviceID    string    `json:"device_id"`
		MemoryUsage float64   `json:"memory_usage"`
		CPUUsage    float64   `json:"cpu_usage"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if requestData.DeviceID == "" {
		http.Error(w, "No data in request body", http.StatusBadRequest)
		return
	}

	//decodedData, err := base64.StdEncoding.Decode()
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusBadRequest)
	//	return
	//}

	var rowToInsert Item
	rowToInsert.Timestamp = requestData.Timestamp
	rowToInsert.CPUUsage = requestData.CPUUsage
	rowToInsert.MemoryUsage = requestData.MemoryUsage
	rowToInsert.DeviceID = requestData.DeviceID

	//if err := json.Unmarshal([]byte(requestData), &rowToInsert); err != nil {
	//	http.Error(w, err.Error(), http.StatusBadRequest)
	//	return
	//}
	//
	// Creating a BigQuery client
	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//
	// Setting up the BigQuery table details
	datasetID := "telemetry_dataset"
	tableID := "telemetry_table"
	//
	//// Creating a BigQuery table reference
	//tableRef := client.Dataset(datasetID).Table(tableID)
	//
	// Defining the BigQuery schema
	schema, err := bigquery.InferSchema(Item{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//
	//// Creating the BigQuery table if it doesn't already exist
	table := client.Dataset(datasetID).Table(tableID)
	if _, err := table.Metadata(ctx); err != nil {
		if err := client.Dataset(datasetID).Create(ctx, &bigquery.DatasetMetadata{}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err := table.Create(ctx, &bigquery.TableMetadata{
			Schema: schema,
		}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	//
	//// Inserting the row into the BigQuery table
	u := table.Inserter()
	if err := u.Put(ctx, &rowToInsert); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, "Data decoded and stored successfully "+requestData.DeviceID)
}
