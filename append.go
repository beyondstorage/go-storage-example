package example

import (
	"bytes"
	"github.com/aos-dev/go-storage/v3/pkg/randbytes"
	"github.com/aos-dev/go-storage/v3/types"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
)

func Append(appender types.Appender, path string) {
	// content to append
	size := rand.Int63n(4 * 1024 * 1024)
	content, _ := ioutil.ReadAll(io.LimitReader(randbytes.NewRand(), size))
	r := bytes.NewReader(content)

	// CreateAppend needs at least one argument.
	//
	// `path` is the path of object.
	// If path is relative path, the real path will be `store.WorkDir + path`.
	// If path is absolute path, the real path will be `path`.
	//
	// CreateAppend will return two values.
	// `o` is the created appendable object.
	// `err` is the error during this operation.
	o, err := appender.CreateAppend(path)
	if err != nil {
		log.Fatalf("CreateAppend #{path}: #{err}")
	}

	// WriteAppend could be called many times. The maximum size of the final appendable object ups to different service.
	//
	// WriteAppend needs at least three arguments.
	//
	// `o` is the appendable object returned by CreateAppend.
	// `r` the read instance for reading the data to append.
	// `size` is the size of content to append.
	//
	// WriteAppend will return two values.
	// `n` is the next append position. It's valid when `err` is nil.
	// `err` is the error during this operation.
	n, err := appender.WriteAppend(o, r, size)
	if err != nil {
		log.Fatalf("WriteAppend %v: %v", path, err)
	}

	log.Printf("append size: %d", n)
}
