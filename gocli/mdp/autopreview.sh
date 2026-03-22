#!/bin/bash

# SPDX-FileCopyrightText: 2021-present j e.j. sahala <jejs@sahala.org>
#
# SPDX-License-Identifier: MPL-2.0

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
