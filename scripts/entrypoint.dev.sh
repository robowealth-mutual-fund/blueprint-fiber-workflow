#!/usr/bin/env bash

export SWAGGER_FILE_NAME=api.json
export BASE_PATH=

set -eu
echo "First, let's init your gRPC port [or blank]: "
read -p 'gRPC port : ' GRPC_SERVER_PORT
echo "done ✅ "
echo "Now, let's init your HTTP port [or blank]: "
read -p 'HTTP port : ' HTTP_SERVER_PORT
echo "done ✅ "

export GRPC_SERVER_PORT
export HTTP_SERVER_PORT
export OTEL_EXPORTER_OTLP_PROTOCOL=grpc
# FIXME Change your OTEL_EXPORTER_OTLP_ENDPOINT
export OTEL_EXPORTER_OTLP_ENDPOINT=http://localhost:4317


if [[ $GRPC_SERVER_PORT == "" ]]; then
  export GRPC_SERVER_PORT=4001
  echo "No gRPC port provided, use ${GRPC_SERVER_PORT} by default"
fi

if [[ $HTTP_SERVER_PORT == "" ]]; then
  export HTTP_SERVER_PORT=4002
  echo "No HTTP port provided, use ${HTTP_SERVER_PORT} by default"
fi

export API_URL=http://localhost:${HTTP_SERVER_PORT}

./scripts/generate_config_js.sh > ./www/swagger-ui/env.js

echo "Blueprint Workflow is starting..."
. ./scripts/setenv.sh
go run cmd/server/server.go
