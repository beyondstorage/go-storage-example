package example

import (
	"errors"
	"log"

	"go.beyondstorage.io/v5/pairs"
	"go.beyondstorage.io/v5/types"
)

func ListAll(store types.Storager) {
	// List all objects or files under the work dir.
	//
	// List needs at least one parameter.
	// `path` is the directory path for file system, or a file hosting service like dropbox, also it could be a prefix filter for object storage.
	//
	// List will return two values.
	// `oi` is an object iterator.
	// `err` is the error during this operation.
	it, err := store.List("")
	if err != nil {
		log.Fatalf("list: %v", err)
	}

	for {
		// User can retrieve all the objects by `Next`. `types.IterateDone` will be returned while there is no item anymore.
		o, err := it.Next()
		if err != nil && !errors.Is(err, types.IterateDone) {
			log.Fatalf("Next: %v", err)
		}

		if err != nil {
			log.Printf("list completed")
			break
		}

		log.Printf("object path: %v", o.Path)
	}
}

func ListDir(store types.Storager, path string) {
	// List with `types.ListModeDir` will list files or objects hierarchically.
	// `path` is the directory path, or a file hosting service, also it could be a prefix filter(usually combined with delimiter `/` internally).
	it, err := store.List(path, pairs.WithListMode(types.ListModeDir))
	if err != nil {
		log.Fatalf("list %v: %v", path, err)
	}

	for {
		o, err := it.Next()
		if err != nil && !errors.Is(err, types.IterateDone) {
			log.Fatalf("Next %v: %v", path, err)
		}

		if err != nil {
			log.Printf("list directory completed: %v", path)
			break
		}

		log.Printf("object path: %v", o.Path)
	}
}

func ListPrefix(store types.Storager, path string) {
	// List with `types.ListModePrefix` will list files or objects with names contain the prefix.
	// `path` is the prefix that the returned object names must contain.
	it, err := store.List(path, pairs.WithListMode(types.ListModePrefix))
	if err != nil {
		log.Fatalf("list %v: %v", path, err)
	}

	for {
		o, err := it.Next()
		if err != nil && !errors.Is(err, types.IterateDone) {
			log.Fatalf("Next %v: %v", path, err)
		}

		if err != nil {
			log.Printf("list with prefix completed: %v", path)
			break
		}

		log.Printf("object path: %v", o.Path)
	}
}

func ListPart(store types.Storager, path string) {
	// List with `types.ListModePart` could retrieve in-progress multipart uploads.
	// `path` is the prefix that the returned object names must contain.
	it, err := store.List(path, pairs.WithListMode(types.ListModePart))
	if err != nil {
		log.Fatalf("list %v: %v", path, err)
	}

	for {
		o, err := it.Next()
		if err != nil && !errors.Is(err, types.IterateDone) {
			log.Fatalf("Next %v: %v", path, err)
		}

		if err != nil {
			log.Printf("list multipart uploads completed: %v", path)
			break
		}

		log.Printf("object path: %v", o.Path)
		log.Printf("object multipartID: %v", o.MustGetMultipartID())
	}
}
