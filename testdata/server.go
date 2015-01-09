// Copyright 2015 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"log"
	"net/http"
	"os/exec"
	"runtime"
	"strconv"
	"time"
)

var (
	flagPort      = flag.Int("port", 9527, "server port")
	flagStaticDir = flag.String("file-dir", ".", "static file dir")
)

func main() {
	flag.Parse()

	go func() {
		url := "http://127.0.0.1:" + strconv.Itoa(*flagPort)
		if waitServer(url) && startBrowser(url) {
			log.Printf("A browser window should open. If not, please visit %s", url)
		} else {
			log.Printf("Please open your web browser and visit %s", url)
		}
	}()

	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(*flagStaticDir))))
	if err := http.ListenAndServe("127.0.0.1:"+strconv.Itoa(*flagPort), nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func waitServer(url string) bool {
	tries := 20
	for tries > 0 {
		resp, err := http.Get(url)
		if err == nil {
			resp.Body.Close()
			return true
		}
		time.Sleep(100 * time.Millisecond)
		tries--
	}
	return false
}

func startBrowser(url string) bool {
	// try to start the browser
	var args []string
	switch runtime.GOOS {
	case "darwin":
		args = []string{"open"}
	case "windows":
		args = []string{"cmd", "/c", "start"}
	default:
		args = []string{"xdg-open"}
	}
	cmd := exec.Command(args[0], append(args[1:], url)...)
	return cmd.Start() == nil
}
