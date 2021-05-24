package example

import (
	"os"

	cos "github.com/beyondstorage/go-service-cos/v2"
	"github.com/beyondstorage/go-storage/v4/pairs"
	"github.com/beyondstorage/go-storage/v4/types"
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
