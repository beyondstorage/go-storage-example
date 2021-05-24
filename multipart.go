package example

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"math/rand"

	"github.com/beyondstorage/go-storage/v4/pairs"
	"github.com/beyondstorage/go-storage/v4/pkg/randbytes"
	"github.com/beyondstorage/go-storage/v4/types"
)

func MultipartUploadTest(store types.Storager, path string) {
	// content to write
	size := rand.Int63n(4 * 1024 * 1024)
	content, _ := ioutil.ReadAll(io.LimitReader(randbytes.NewRand(), size))
	r := bytes.NewReader(content)

	// `store` should implement `Multiparter`
	multiparter, ok := store.(types.Multiparter)
	if !ok {
		log.Fatalf("Multiparter unimplemented")
	}

	// CreateMultipart needs at least one argument.
	//
	// `path` is the path of object.
	// If path is relative path, the real path will be `store.WorkDir + path`.
	// If path is absolute path, the real path will be `path`.
	//
	// CreateMultipart will return two values.
	// `o` is the created multipart object.
	// `err` is the error during this operation.
	o, err := multiparter.CreateMultipart(path)
	if err != nil {
		log.Fatalf("CreateMultipart %v: %v", path, err)
	}

	// `isValidMultipartID` indicates whether the multipartId is valid.
	var isValidMultipartID = true

	// Delete with multipartId could be called when you want to abort the multipart upload or error occurred.
	//
	// Delete with multipartId needs at least two arguments.
	//
	// `path` is the path of the multipart object.
	// `pairs` is the optional argument and should take multipartId.
	//
	// Delete with multipartId will return one value.
	// `err` is the error during this operation.
	defer func(hasMultipartID bool) {
		if hasMultipartID {
			err := store.Delete(path, pairs.WithMultipartID(o.MustGetMultipartID()))
			if err != nil {
				log.Fatalf("DeleteWithMultipartID %v: %v", path, err)
			}
		}
	}(isValidMultipartID)

	// WriteMultipart could be called concurrently.
	//
	// WriteMultipart needs at least four arguments.
	//
	// `o` is the object returned by CreateMultipart.
	// `r` the read instance for reading the data to upload.
	// `size` is the size of content to upload.
	// `index` is the part number. It's zero-based and should be [0, 9,999] for the current supported services.
	//
	// WriteMultipart will return three values.
	// `n` is the size of write part operation. It's valid when `err` is nil.
	// `part` is the part information, include `Index`, `Size` and `ETag`.
	// `err` is the error during this operation.
	n, part, err := multiparter.WriteMultipart(o, r, size, 0)
	if err != nil {
		log.Fatalf("WriteMultipart %v: %v", path, err)
	}

	// ListMultipart needs at least one argument.
	//
	// `o` is the object returned by CreateMultipart.
	//
	// CompleteMultipart will return two values.
	// `pi` is the part information iterator.
	// `err` is the error during this operation.
	it, err := multiparter.ListMultipart(o)
	if err != nil {
		log.Fatalf("ListMultipart %v: %v", path, err)
	}

	// You can traverse the list of parts information through `Next()`.
	for {
		p, err := it.Next()
		if err != nil && !errors.Is(err, types.IterateDone) {
			log.Fatalf("Next %v: %v", path, err)
		}

		if err != nil {
			log.Println("Next completed: ", path)
			break
		}

		log.Printf("Index: %d, Size: %d, ETag: %v", p.Index, p.Size, p.ETag)
	}

	// CompleteMultipart needs at least two arguments.
	//
	// `o` is the object returned by CreateMultipart.
	// `parts` is the list of parts information consist of the return value of WriteMultipart.
	//
	// CompleteMultipart will return one value.
	// `err` is the error during this operation.
	err = multiparter.CompleteMultipart(o, []*types.Part{part})
	if err != nil {
		log.Fatalf("CompleteMultipart %v: %v", path, err)
	}

	// The `multipartId` is invalid after completing multipart upload successfully.
	isValidMultipartID = false

	log.Printf("multipart upload size: %d", n)
}

func MultiparterWithUploadIdTest(store types.Storager, path string) {
	size := rand.Int63n(4 * 1024 * 1024)
	content, _ := ioutil.ReadAll(io.LimitReader(randbytes.NewRand(), size))
	r := bytes.NewReader(content)

	multiparter, ok := store.(types.Multiparter)
	if !ok {
		log.Fatalf("Multiparter unimplemented")
	}

	o, err := multiparter.CreateMultipart(path)
	if err != nil {
		log.Fatalf("CreateMultipart %v: %v", path, err)
	}

	multipartId := o.MustGetMultipartID()

	var isValidMultipartID = true
	defer func(hasMultipartID bool) {
		if hasMultipartID {
			err := store.Delete(path, pairs.WithMultipartID(multipartId))
			if err != nil {
				log.Fatalf("DeleteWithMultipartID %v: %v", path, err)
			}
		}
	}(isValidMultipartID)

	// Create with multipartId could be called when you want to create a multipart object with a known multipartId.
	//
	// Create with multipartId needs at least two arguments.
	//
	// `path` is the path of object..
	// `pairs` is the optional argument and should take multipartId obtained from CreateMultipart.
	//
	// Create with multipartId will return one value.
	// `mo` is the created multipart object.
	mo := store.Create(path, pairs.WithMultipartID(multipartId))

	n, part, err := multiparter.WriteMultipart(mo, r, size, 0)
	if err != nil {
		log.Fatalf("WriteMultipart %v: %v", path, err)
	}

	err = multiparter.CompleteMultipart(mo, []*types.Part{part})
	if err != nil {
		log.Fatalf("CompleteMultipart %v: %v", path, err)
	}

	isValidMultipartID = false
	log.Printf("multipart upload size: %d", n)
}
