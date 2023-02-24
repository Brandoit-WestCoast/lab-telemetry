
## Deploy Google Cloud Function
```
cd gcp/

gcloud functions deploy http_to_bq \
--trigger-http \
--allow-unauthenticated
```

## Test Google Cloud Function
```
curl -X POST -H "Content-Type: application/json" -d '{
  "timestamp": "2023-01-31T12:34:56.789Z",
  "device_id": "abc123",
  "memory_usage": 0.45,
  "cpu_usage": 0.23
}' "FUNCTION_URL"
```