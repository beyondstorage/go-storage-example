package example

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"math/rand"

	"github.com/aos-dev/go-storage/v3/pkg/randbytes"
	"github.com/aos-dev/go-storage/v3/types"
)

func AppendToNewFile(appender types.Appender, path string) {
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

	// WriteAppend could be called many times. The maximum size of the final appendable object ups to different services.
	//
	// WriteAppend needs at least three arguments.
	//
	// `o` is the appendable object returned by CreateAppend. It specifies the next call's append position, so the caller need not to maintain this information.
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

func AppendToExistingFile(store types.Storager, path string) {
	// content to append
	size := rand.Int63n(4 * 1024 * 1024)
	content, _ := ioutil.ReadAll(io.LimitReader(randbytes.NewRand(), size))
	r := bytes.NewReader(content)

	// `store` should implement `Appender`
	appender, ok := store.(types.Appender)
	if !ok {
		log.Fatalf("Appender unimplemented")
	}

	// Use `Stat` to get an appendable object.
	//
	// Stat needs at least one argument.
	//
	// `path` is the path of object.
	// If path is relative path, the real path will be `store.WorkDir + path`.
	// If path is absolute path, the real path will be `path`.
	//
	// Stat will return two values.
	// `o` is the existing object.
	// `err` is the error during this operation.
	o, err := store.Stat(path)
	if err != nil {
		log.Fatalf("Stat %v: %v", path, err)
	}

	// `o` is the object returned by Stat.
	// The service should check if the object `isAppend` and maintains the next call's append position.
	n, err := appender.WriteAppend(o, r, size)
	if err != nil {
		log.Fatalf("WriteAppend %v: %v", path, err)
	}

	log.Printf("append size: %d", n)
}
