// SPDX-FileCopyrightText: 2021-present j e.j. sahala <jejs@sahala.org>
//
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"log"

	"github.com/joshuaejs/godcls/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
