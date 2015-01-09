// Copyright 2015 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"flag"
	"log"
	"net/http"
)

var (
	flagPort      = flag.Int("port", 9527, "server port")
	flagStaticDir = flag.String("file-dir", ".", "static file dir")
)

func main() {
	flag.Parse()
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(*flagStaticDir))))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", *flagPort), nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
