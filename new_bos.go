package example

import (
	"fmt"
	"os"

	bos "go.beyondstorage.io/services/bos/v2"
	"go.beyondstorage.io/v5/pairs"
	"go.beyondstorage.io/v5/services"
	"go.beyondstorage.io/v5/types"
)

func NewBos() (types.Storager, error) {
	return bos.NewStorager(
		// work_dir: https://beyondstorage.io/docs/go-storage/pairs/work_dir
		//
		// Relative operations will be based on this WorkDir.
		pairs.WithWorkDir(os.Getenv("STORAGE_BOS_WORKDIR")),
		// credential: https://beyondstorage.io/docs/go-storage/pairs/credential
		//
		// Credential could be fetched from service's console: https://console.bce.baidu.com/iam/#/iam/accesslist
		//
		// Example Value: hmac:access_key_id:secret_access_key
		pairs.WithCredential(os.Getenv("STORAGE_BOS_CREDENTIAL")),
		// endpoint: https://beyondstorage.io/docs/go-storage/pairs/endpoint
		//
		// Available endpoint: https://cloud.baidu.com/doc/BOS/s/akrqd2wcx
		pairs.WithEndpoint(os.Getenv("STORAGE_BOS_ENDPOINT")),
		// name: https://beyondstorage.io/docs/go-storage/pairs/name
		//
		// name is the bucket name.
		pairs.WithName(os.Getenv("STORAGE_BOS_NAME")),
	)
}

func NewBosFromString() (types.Storager, error) {
	connStr := fmt.Sprintf(
		"bos://%s%s?credential=%s&location=%s",
		os.Getenv("STORAGE_BOS_NAME"),
		os.Getenv("STORAGE_BOS_WORKDIR"),
		os.Getenv("STORAGE_BOS_CREDENTIAL"),
		os.Getenv("STORAGE_BOS_ENDPOINT"),
	)
	return services.NewStoragerFromString(connStr)
}
