package basic

import (
	"os"

	cos "github.com/aos-dev/go-service-cos"
	"github.com/aos-dev/go-storage/v3/pairs"
	"github.com/aos-dev/go-storage/v3/types"
)

func NewCos() (types.Storager, error) {
	return cos.NewStorager(
		// work_dir: https://aos.dev/docs/go-storage/pairs/work_dir
		//
		// Relative operations will be based on this WorkDir.
		pairs.WithWorkDir(os.Getenv("STORAGE_COS_WORKDIR")),
		// credential: https://aos.dev/docs/go-storage/pairs/credential
		//
		// Credential could be fetched from service's console: https://console.cloud.tencent.com/cam/capi
		//
		// Example Value: hmac:access_key_id:secret_access_key
		pairs.WithCredential(os.Getenv("STORAGE_COS_CREDENTIAL")),
		// location: https://aos.dev/docs/go-storage/pairs/location
		//
		// Available location: https://cloud.tencent.com/document/product/436/6224
		pairs.WithLocation(os.Getenv("STORAGE_COS_LOCATION")),
		// name: https://aos.dev/docs/go-storage/pairs/name
		//
		// name is the bucket name.
		pairs.WithName(os.Getenv("STORAGE_COS_NAME")),
	)
}
