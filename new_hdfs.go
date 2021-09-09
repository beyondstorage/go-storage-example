package example

import (
	"fmt"
	"os"

	hdfs "github.com/beyondstorage/go-service-hdfs"
	"github.com/beyondstorage/go-storage/v4/pairs"
	"github.com/beyondstorage/go-storage/v4/services"
	"github.com/beyondstorage/go-storage/v4/types"
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