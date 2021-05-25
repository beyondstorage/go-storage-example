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

func Multipart(store types.Storager, path string) {
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

	log.Printf("multipart upload size: %d", n)
}

func ResumeMultipart(store types.Storager, path string) {
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

	// Create with multipartId could be called when you want to resume multipart upload.
	//
	// Create with multipartId needs at least two arguments.
	//
	// `path` is the path of object.
	// `pairs` is the optional argument and should take multipartId obtained from CreateMultipart.
	//
	// Create with multipartId will return one value.
	// `mo` is the created multipart object.
	mo := store.Create(path, pairs.WithMultipartID(multipartId))

	// `partNumber` indicates the last uploaded part number.
	var partNumber = -1
	// `totalSize` indicates the total upload size.
	var totalSize int64 = 0

	// List all parts that have been uploaded for the specific multipartId.
	//
	// ListMultipart needs at least one argument.
	//
	// `mo` is the object returned by Create.
	//
	// ListMultipart will return two values.
	// `it` is the part information iterator.
	// `err` is the error during this operation.
	it, err := multiparter.ListMultipart(mo)
	if err != nil {
		log.Fatalf("ListMultipart %v: %v", path, err)
	}

	// Traverse the iterator through `Next()` to get all the uploaded parts.
	var parts []*types.Part
	for {
		p, err := it.Next()
		if err != nil && !errors.Is(err, types.IterateDone) {
			log.Fatalf("Next %v: %v", path, err)
		}

		if err != nil {
			log.Printf("ListMultipart completed: %v", path)
			break
		}

		partNumber = p.Index
		totalSize += p.Size
		parts = append(parts, p)
	}

	n, part, err := multiparter.WriteMultipart(mo, r, size, partNumber+1)
	if err != nil {
		log.Fatalf("WriteMultipart %v: %v", path, err)
	}

	totalSize += n
	parts = append(parts, part)

	err = multiparter.CompleteMultipart(mo, parts)
	if err != nil {
		log.Fatalf("CompleteMultipart %v: %v", path, err)
	}

	log.Printf("total upload size: %d", totalSize)
}

func CancelMultipart(store types.Storager, path string) {
	multiparter, ok := store.(types.Multiparter)
	if !ok {
		log.Fatalf("Multiparter unimplemented")
	}

	o, err := multiparter.CreateMultipart(path)
	if err != nil {
		log.Fatalf("CreateMultipart %v: %v", path, err)
	}

	// Delete with multipartId could be called when you want to abort the multipart upload or error occurred.
	//
	// Delete with multipartId needs at least two arguments.
	//
	// `path` is the path of the multipart object.
	// `pairs` is the optional argument and should take multipartId.
	//
	// Delete with multipartId will return one value.
	// `err` is the error during this operation.
	err = store.Delete(path, pairs.WithMultipartID(o.MustGetMultipartID()))
	if err != nil {
		log.Fatalf("Delete with multipartId %v: %v", path, err)
	}

	log.Printf("cancel multipart: %v", path)
}