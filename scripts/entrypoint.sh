#!/bin/sh

/scripts/generate_config_js.sh > /www/swagger-ui/env.js

echo "Start generate swagger JSON file:"
./app/bin/swagger

echo "Start Server:"
./app/bin/server
