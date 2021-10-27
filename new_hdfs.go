package example

import (
	"fmt"
	"os"

	hdfs "go.beyondstorage.io/services/hdfs"
	"go.beyondstorage.io/v5/pairs"
	"go.beyondstorage.io/v5/services"
	"go.beyondstorage.io/v5/types"
)

func NewHDFS() (types.Storager, error) {
	return hdfs.NewStorager(
		// work_dir: https://beyondstorage.io/docs/go-storage/pairs/work_dir
		//
		// Relative operations will be based on this WorkDir.
		pairs.WithWorkDir(os.Getenv("STORAGE_HDFS_WORKDIR")),
		// endpoint: https://beyondstorage.io/docs/go-storage/pairs/endpoint
		//
		// Example Value: tcp:host:port
		pairs.WithEndpoint(os.Getenv("STORAGE_HDFS_ENDPOINT")),
	)
}

func NewHDFSFromString() (types.Storager, error) {
	str := fmt.Sprintf(
		"hdfs://%s?endpoint=%s",
		os.Getenv("STORAGE_HDFS_WORKDIR"),
		os.Getenv("STORAGE_HDFS_ENDPOINT"),
	)
	return services.NewStoragerFromString(str)
}
