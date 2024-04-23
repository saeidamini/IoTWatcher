aws dynamodb create-table --endpoint-url http://localhost:8000 --cli-input-json file://schema/devices.create.json --profile default
aws dynamodb batch-write-item --endpoint-url http://localhost:8000 --request-items file://schema/devices.seed.json --profile default
aws dynamodb scan --table-name  saeid-amn-Devices --profile default