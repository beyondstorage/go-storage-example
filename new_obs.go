package example

import (
	"fmt"
	"os"

	obs "github.com/beyondstorage/go-service-obs"
	"github.com/beyondstorage/go-storage/v4/pairs"
	"github.com/beyondstorage/go-storage/v4/services"
	"github.com/beyondstorage/go-storage/v4/types"
)

func NewObs() (types.Storager, error) {
	return obs.NewStorager(
		// work_dir: https://beyondstorage.io/docs/go-storage/pairs/work_dir
		//
		// Relative operations will be based on this WorkDir.
		pairs.WithWorkDir(os.Getenv("STORAGE_OBS_WORKDIR")),
		// credential: https://beyondstorage.io/docs/go-storage/pairs/credential
		//
		// Credential could be fetched from service's console: https://console.huaweicloud.com/iam/?region=cn-east-3#/mine/accessKey
		//
		// Example Value: hmac:access_key_id:secret_access_key
		pairs.WithCredential(os.Getenv("STORAGE_OBS_CREDENTIAL")),
		// endpoint: https://beyondstorage.io/docs/go-storage/pairs/endpoint
		//
		// Available endpoint: https://developer.huaweicloud.com/endpoint?OBS
		pairs.WithEndpoint(os.Getenv("STORAGE_OBS_ENDPOINT")),
		// name: https://beyondstorage.io/docs/go-storage/pairs/name
		//
		// name is the bucket name.
		pairs.WithName(os.Getenv("STORAGE_OBS_NAME")),
	)
}

func NewObsFromString() (types.Storager, error) {
	connStr := fmt.Sprintf(
		"obs://%s%s?credential=%s&location=%s",
		os.Getenv("STORAGE_OBS_NAME"),
		os.Getenv("STORAGE_OBS_WORKDIR"),
		os.Getenv("STORAGE_OBS_CREDENTIAL"),
		os.Getenv("STORAGE_OBS_ENDPOINT"),
	)
	return services.NewStoragerFromString(connStr)
}
