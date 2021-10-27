package example

import (
	"fmt"
	"os"

	ipfs "go.beyondstorage.io/services/ipfs"
	"go.beyondstorage.io/v5/pairs"
	"go.beyondstorage.io/v5/services"
	"go.beyondstorage.io/v5/types"
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

func NewIPFSFromString() (types.Storager, error) {
	connStr := fmt.Sprintf(
		"ipfs://%s?endpoint=%s",
		os.Getenv("STORAGE_IPFS_WORKDIR"),
		os.Getenv("STORAGE_IPFS_ENDPOINT"),
	)
	return services.NewStoragerFromString(connStr)
}
