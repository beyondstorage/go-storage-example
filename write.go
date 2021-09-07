package example

import (
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/beyondstorage/go-storage/v4/pairs"
	"github.com/beyondstorage/go-storage/v4/pkg/randbytes"
	"github.com/beyondstorage/go-storage/v4/types"
)

func WriteData (store types.Storager, path string) {
	// content to write
	size := rand.Int63n(4 * 1024 * 1024)
	r := io.LimitReader(randbytes.NewRand(), size)

	// Write needs at least three arguments.
	// `path` is the path of object.
	// `r` is io.Reader instance for reading the data for uploading.
	// `size` is the length, in bytes, of the data for uploading.
	//
	// Write will return two values.
	// `n` is the size of write operation.
	// `err` is the error during this operation.
	n, err := store.Write(path, r, size)
	if err != nil {
		log.Fatalf("write %v: %v", path, err)
	}

	log.Printf("write size: %d", n)
}

func WriteWithCallback (store types.Storager, path string) {
	// content to write
	size := rand.Int63n(4 * 1024 * 1024)
	r := io.LimitReader(randbytes.NewRand(), size)

	cur := int64(0)
	fn := func(bs []byte) {
		cur += int64(len(bs))
		log.Printf("write %d bytes already", cur)
	}

	// If IoCallback is specified, the storage will call it in every I/O operation.
	// User could use this feature to implement progress bar.
	n, err := store.Write(path, r, size, pairs.WithIoCallback(fn))
	if err != nil {
		log.Fatalf("write %v: %v", path, err)
	}

	log.Printf("write size: %d", n)
}

func WriteWithSignedURL (store types.Storager, path string, expire time.Duration) {
	signer, ok := store.(types.StorageHTTPSigner)
	if !ok {
		log.Fatalf("StorageHTTPSigner unimplemented")
	}

	size := rand.Int63n(4 * 1024 * 1024)
	r := io.LimitReader(randbytes.NewRand(), size)

	// QuerySignHTTPWrite needs at least three arguments.
	// `path` is the path of object.
	// `size` is the length, in bytes, of the data for uploading.
	// `expire` provides the time period, with type time.Duration, for which the generated req.URL is valid.
	//
	// QuerySignHTTPWrite will return two values.
	//
	// `req` is the generated `*http.Request`:
	// `req.URL` specifies the URL to access with signature in the query string.
	// `req.Header` specifies the HTTP headers included in the signature.
	// `req.ContentLength` records the length of the associated content, the value equals to `size`.
	//
	// `err` is the error during this operation.
	req, err := signer.QuerySignHTTPWrite(path, size, expire)
	if err !=nil {
		log.Fatalf("write %v: %v", path, err)
	}

	// Set request body.
	req.Body = ioutil.NopCloser(r)

	client := http.Client{}
	_, err = client.Do(req)
	if err != nil {
		log.Fatalf("send HTTP request for writing %v: %v", path, err)
	}

	log.Printf("write size: %d", size)
}
