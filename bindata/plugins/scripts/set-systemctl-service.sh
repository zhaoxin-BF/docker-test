#!/bin/bash

set -e
set -o pipefail

setSystemCtlService() {
  local EVERAI_NODE_HOME="$1"
  echo "EVERAI_NODE_HOME: $EVERAI_NODE_HOME"

# Define the service template
  SERVICE_TEMPLATE='[Unit]
Description=EverAI Resource Node Server
After=network.target

[Service]
WorkingDirectory='"${EVERAI_NODE_HOME}"'
ExecStart='"${EVERAI_NODE_HOME}"'/everai-resource-node
Restart=always
User=root

[Install]
WantedBy=multi-user.target'

  # Write the service template to a file
  # echo "$SERVICE_TEMPLATE" > "$EVERAI_NODE_HOME"/temp/everai-resource-node.service
  echo "$SERVICE_TEMPLATE" > /etc/systemd/system/everai-resource-node.service
}

# Call the setSystemCtlService function with the resource path
if [ "$#" -ne 1 ]; then
  echo "Usage: $0 <EVERAI_NODE_HOME>"
  exit 1
fi

setSystemCtlService "$1"