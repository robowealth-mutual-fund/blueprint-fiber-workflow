#!/bin/sh -eu
cat <<EOF
window.appConfig = {
    API_URL: "${API_URL}${BASE_PATH}/swagger-ui"
    }
EOF