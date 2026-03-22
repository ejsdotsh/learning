# SPDX-FileCopyrightText: 2021-present j e.j. sahala <jejs@sahala.org>
#
# SPDX-License-Identifier: MPL-2.0

.PHONY: run
run:
	go run main.go

.PHONY: test
test:
	go test -v -race ./... | sed ''/PASS/s//$$(printf "\033[32mPASS\033[0m")/'' | sed ''/FAIL/s//$$(printf "\033[31mFAIL\033[0m")/''