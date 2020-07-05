- *Go语言QQ群: 102319854, 1055927514*
- *凹语言(凹读音“Wa”)(The Wa Programming Language): https://github.com/wa-lang/wa*

----

# BPG for Go

[![Build Status](https://travis-ci.org/chai2010/bpg.svg)](https://travis-ci.org/chai2010/bpg)
[![GoDoc](https://godoc.org/github.com/chai2010/bpg?status.svg)](https://godoc.org/github.com/chai2010/bpg)

BPG is defined at:
http://bellard.org/bpg/

# Install

Install `GCC` or `MinGW` (http://tdm-gcc.tdragon.net/download) at first,
and then run these commands:

	1. Assure set the `CGO_ENABLED` environment variable to `1` to enable `CGO` (Default is enabled).
	2. `go get github.com/chai2010/bpg`
	3. `go run hello.go`


# Examples

This is a simple example:

```Go
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
```


# BUGS

Report bugs to <chaishushan@gmail.com>.

Thanks!
