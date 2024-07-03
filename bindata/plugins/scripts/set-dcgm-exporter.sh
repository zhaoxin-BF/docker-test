#!/bin/bash

set -e
set -o pipefail

# dst /etc
#DST_DIR="/etc"
DST_DIR="/Users/zhaoxin/Workspace/github.com/zhaoxin-BF/docker-test/bindata/temp"

syncDcgmExporter() {
  local EVERAI_EXEC_PATH="/Users/zhaoxin/Workspace/github.com/zhaoxin-BF/docker-test/bindata/plugins"
  echo "EVERAI_EXEC_PATH: $EVERAI_EXEC_PATH"

  sudo mkdir -p "${DST_DIR}"


  file1="${EVERAI_EXEC_PATH}/scripts/1.x-compatibility-metrics.csv"
  file2="${EVERAI_EXEC_PATH}/scripts/dcp-metrics-included.csv"
  file3="${EVERAI_EXEC_PATH}/scripts/default-counters.csv"

  sudo cp "$file1" "$file2" "$file3"  "$DST_DIR"
  return 0
}

# Call the syncDcgmExporter function with the resource path
if [ "$#" -ne 1 ]; then
  echo "Usage: $0 <EVERAI_EXEC_PATH>"
  exit 1
fi

syncDcgmExporter "$1"