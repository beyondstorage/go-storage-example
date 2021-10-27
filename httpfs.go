package example

import (
	"io"
	"log"

	"go.beyondstorage.io/v5/pkg/fswrap"
	"go.beyondstorage.io/v5/types"
)

func HttpFSOpen(store types.Storager, path string) {
	fsys := fswrap.HttpFs(store)

	f, err := fsys.Open(path)
	if err != nil {
		log.Fatalf("Open %v: %v", path, err)
	}

	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		log.Fatalf("Stat %v: %v", path, err)
	}

	log.Printf("file name: %v", info.Name())
}

func HttpFSReadDir(store types.Storager, path string) {
	fsys := fswrap.HttpFs(store)

	f, err := fsys.Open(path)
	if err != nil {
		log.Fatalf("Open %v: %v", path, err)
	}

	defer f.Close()

	list, err := f.Readdir(-1)
	if err != nil {
		log.Fatalf("Readdir %v: %v", path, err)
	}

	for _, info := range list {
		log.Printf("file name: %v", info.Name())
	}
}

func HttpFsRead(store types.Storager, path string) {
	fsys := fswrap.HttpFs(store)

	f, err := fsys.Open(path)
	if err != nil {
		log.Fatalf("Open %v: %v", path, err)
	}

	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		log.Fatalf("Stat %v: %v", path, err)
	}

	size := info.Size()
	data := make([]byte, 0, size+1)

	n, err := f.Read(data)
	if err != nil {
		log.Fatalf("Read %v: %v", path, err)
	}

	log.Printf("read data length: %d", n)
}

func HttpFsSeek(store types.Storager, path string) {
	fsys := fswrap.HttpFs(store)

	f, err := fsys.Open(path)
	if err != nil {
		log.Fatalf("Open %v: %v", path, err)
	}

	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		log.Fatalf("Stat %v: %v", path, err)
	}

	size := info.Size()

	got, err := f.Seek(0, io.SeekEnd)
	if err != nil || got != size {
		log.Fatalf("Seek %v: %v", path, err)
	}
}
