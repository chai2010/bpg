// Copyright 2015 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ingore

package main

import (
	"bytes"
	"fmt"
	"image/png"
	"io/ioutil"
	"log"

	"github.com/chai2010/bpg"
)

func main() {
	var buf bytes.Buffer
	var data []byte
	var err error

	data, err = ioutil.ReadFile("./testdata/lena512color.bpg")
	if err != nil {
		log.Println(err)
	}
	info, err := bpg.DecodeInfo(data)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("info: %v\n", info)

	// Decode bpg
	m, err := bpg.Decode(bytes.NewReader(data))
	if err != nil {
		log.Println(err)
	}

	// save as png
	if err = png.Encode(&buf, m); err != nil {
		log.Println(err)
	}
	if err = ioutil.WriteFile("output.png", buf.Bytes(), 0666); err != nil {
		log.Println(err)
	}

	fmt.Println("Save as output.png")
}
