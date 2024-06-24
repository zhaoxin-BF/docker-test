#!/bin/bash

set -e
set -o pipefail

# dst /etc
#DST_DIR="/etc"
DST_DIR="/Users/zhaoxin/Workspace/github.com/zhaoxin-BF/docker-test/bindata/temp"

syncDcgmExporter() {
  local EVERAI_NODE_HOME="$1"
  echo "EVERAI_NODE_HOME: $EVERAI_NODE_HOME"

  SRC_DIR="$EVERAI_NODE_HOME/dcgm-exporter"
  echo "SRC_DIR"

  # cp dir
  sudo cp -r "$SRC_DIR" "$DST_DIR"

  if [ $? -eq 0 ]; then
      echo "Successfully copied '$SRC_DIR' to '$DST_DIR'."
  else
      echo "Error occurred while copying '$SRC_DIR' to '$DST_DIR'."
  fi
  return 0
}

# Call the syncDcgmExporter function with the resource path
if [ "$#" -ne 1 ]; then
  echo "Usage: $0 <EVERAI_NODE_HOME>"
  exit 1
fi

syncDcgmExporter "$1"