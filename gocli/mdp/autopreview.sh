#!/bin/bash

# TODO set flags

FHASH=$(md5sum $1)

# watch the file until interrupted
while true; do
  NHASH=$(md5sum $1)
  if [ "${NHASH}" != "${FHASH}" ]; then
    go run main.go -file $1
    FHASH=${NHASH}
  fi
  sleep 5
done
