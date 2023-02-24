package telemetry

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type MyStruct struct {
	FieldName1 string    `json:"field_name_1"`
	FieldName2 int       `json:"field_name_2"`
	FieldName3 time.Time `json:"field_name_3"`
}

func StoreToBigQuery(w http.ResponseWriter, r *http.Request) {
	// Extracting data from the HTTP POST request
	var requestData struct {
		Timestamp   string  `json:"timestamp"`
		DeviceID    string  `json:"device_id"`
		MemoryUsage float64 `json:"memory_usage"`
		CPUUsage    float64 `json:"cpu_usage"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if requestData.Timestamp == "" {
		http.Error(w, "No data in request body", http.StatusBadRequest)
		return
	}

	//decodedData, err := base64.StdEncoding.DecodeString(requestData.Data)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusBadRequest)
	//	return
	//}

	//var rowToInsert MyStruct
	//if err := json.Unmarshal(decodedData, &rowToInsert); err != nil {
	//	http.Error(w, err.Error(), http.StatusBadRequest)
	//	return
	//}
	//
	//// Creating a BigQuery client
	//ctx := context.Background()
	//client, err := bigquery.NewClient(ctx, "my-project-id")
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
	//
	//// Setting up the BigQuery table details
	//datasetID := "my_dataset"
	//tableID := "my_table"
	//
	//// Creating a BigQuery table reference
	//tableRef := client.Dataset(datasetID).Table(tableID)
	//
	//// Defining the BigQuery schema
	//schema, err := bigquery.InferSchema(MyStruct{})
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
	//
	//// Creating the BigQuery table if it doesn't already exist
	//table := client.Dataset(datasetID).Table(tableID)
	//if _, err := table.Metadata(ctx); err != nil {
	//	if _, err := client.Dataset(datasetID).Create(ctx, &bigquery.DatasetMetadata{}); err != nil {
	//		http.Error(w, err.Error(), http.StatusInternalServerError)
	//		return
	//	}
	//	if err := table.Create(ctx, &bigquery.TableMetadata{
	//		Schema: schema,
	//	}); err != nil {
	//		http.Error(w, err.Error(), http.StatusInternalServerError)
	//		return
	//	}
	//}
	//
	//// Inserting the row into the BigQuery table
	//u := table.Uploader()
	//if err := u.Put(ctx, &rowToInsert); err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}

	fmt.Fprint(w, "Data decoded successfully"+requestData.DeviceID)
}
