package example

import (
	"fmt"
	"os"

	cos "go.beyondstorage.io/services/cos/v3"
	"go.beyondstorage.io/v5/pairs"
	"go.beyondstorage.io/v5/services"
	"go.beyondstorage.io/v5/types"
)

func NewCos() (types.Storager, error) {
	return cos.NewStorager(
		// work_dir: https://beyondstorage.io/docs/go-storage/pairs/work_dir
		//
		// Relative operations will be based on this WorkDir.
		pairs.WithWorkDir(os.Getenv("STORAGE_COS_WORKDIR")),
		// credential: https://beyondstorage.io/docs/go-storage/pairs/credential
		//
		// Credential could be fetched from service's console: https://console.cloud.tencent.com/cam/capi
		//
		// Example Value: hmac:access_key_id:secret_access_key
		pairs.WithCredential(os.Getenv("STORAGE_COS_CREDENTIAL")),
		// location: https://beyondstorage.io/docs/go-storage/pairs/location
		//
		// Available location: https://cloud.tencent.com/document/product/436/6224
		pairs.WithLocation(os.Getenv("STORAGE_COS_LOCATION")),
		// name: https://beyondstorage.io/docs/go-storage/pairs/name
		//
		// name is the bucket name.
		pairs.WithName(os.Getenv("STORAGE_COS_NAME")),
	)
}

func NewCosFromString() (types.Storager, error) {
	connStr := fmt.Sprintf(
		"cos://%s%s?credential=%s&location=%s",
		os.Getenv("STORAGE_COS_NAME"),
		os.Getenv("STORAGE_COS_WORKDIR"),
		os.Getenv("STORAGE_COS_CREDENTIAL"),
		os.Getenv("STORAGE_COS_LOCATION"),
	)
	return services.NewStoragerFromString(connStr)
}
