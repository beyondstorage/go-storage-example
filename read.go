package example

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/beyondstorage/go-storage/v4/pairs"
	"github.com/beyondstorage/go-storage/v4/types"
)

func ReadWhole(store types.Storager, path string) {
	var buf bytes.Buffer

	// Read needs at least two arguments.
	//
	// `path` is the path of object.
	// If path is relative path, the real path will be `store.WorkDir + path`.
	// If path is absolute path, the real path will be `path`.
	//
	// `w`, the `&buf` here is the writer of this operation.
	// storage will write all content that read into this writer.
	// It's caller's duty to make sure the writer has been closed.
	//
	// Read will return two values.
	// `n` is the size of read operation.
	// `err` is the error during this operation.
	n, err := store.Read(path, &buf)
	if err != nil {
		log.Fatalf("read %v: %v", path, err)
	}

	log.Printf("read size: %d", n)
	log.Printf("read content: %s", buf.Bytes())
}

func ReadRange(store types.Storager, path string, offset, size int64) {
	var buf bytes.Buffer

	// Offset is the read operation's offset.
	// Size is the read operation's size.
	//
	// In this read operation, we will read content in [offset, offset+size).
	n, err := store.Read(path, &buf,
		pairs.WithOffset(offset),
		pairs.WithSize(size),
	)
	if err != nil {
		log.Fatalf("read %v: %v", path, err)
	}

	log.Printf("read size: %d", n)
	log.Printf("read content: %s", buf.Bytes())
}

func ReadWithCallback(store types.Storager, path string) {
	var buf bytes.Buffer

	cur := int64(0)
	fn := func(bs []byte) {
		cur += int64(len(bs))
		log.Printf("read %d bytes already", cur)
	}

	// If IoCallback is specified, the storage will call it in every I/O operation.
	// User could use this feature to implement progress bar.
	n, err := store.Read(path, &buf, pairs.WithIoCallback(fn))
	if err != nil {
		log.Fatalf("read %v: %v", path, err)
	}

	log.Printf("read size: %d", n)
	log.Printf("read content: %s", buf.Bytes())
}

func ReadWithSignedURL(store types.Storager, path string, expire time.Duration) {
	signer, ok := store.(types.StorageHTTPSigner)
	if !ok {
		log.Fatalf("StorageHTTPSigner unimplemented")
	}

	// QuerySignHTTPRead needs at least two arguments.
	//
	// `path` is the path of object.
	// `expire` provides the time period, with type time.Duration, for which the generated req.URL is valid.
	//
	// QuerySignHTTPRead will return two values.
	// `req` is the generated `*http.Request`, `req.URL` specifies the URL to access with signature in the query string. And `req.Header` specifies the HTTP headers included in the signature.
	// `err` is the error during this operation.
	req, err := signer.QuerySignHTTPRead(path, expire)
	if err !=nil {
		log.Fatalf("read %v: %v", path, err)
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("send HTTP request for reading %v: %v", path, err)
	}

	defer func() {
		err = resp.Body.Close()
		if err != nil {
			log.Fatalf("close HTTP response body for reading %v: %v", path, err)
		}
	}()

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("read from HTTP response body for reading %v: %v", path, err)
	}

	log.Printf("read size: %d", resp.ContentLength)
	log.Printf("read content: %s", buf)
}
