package example

import (
	"os"

	ipfs "github.com/beyondstorage/go-service-ipfs"
	"github.com/beyondstorage/go-storage/v4/pairs"
	"github.com/beyondstorage/go-storage/v4/types"
)

func NewIPFS() (types.Storager, error) {
	return ipfs.NewStorager(
		// work_dir: https://beyondstorage.io/docs/go-storage/pairs/work_dir
		//
		// Relative operations will be based on this WorkDir.
		pairs.WithWorkDir(os.Getenv("STORAGE_IPFS_WORKDIR")),
		// endpoint: https://beyondstorage.io/docs/go-storage/pairs/endpoint
		//
		// Example Value: https:host:port
		pairs.WithEndpoint(os.Getenv("STORAGE_IPFS_ENDPOINT")),
	)
}
