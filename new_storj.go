package example

import (
	"fmt"
	"os"

	storj "github.com/beyondstorage/go-service-storj"
	"github.com/beyondstorage/go-storage/v4/pairs"
	"github.com/beyondstorage/go-storage/v4/services"
	"github.com/beyondstorage/go-storage/v4/types"
)

func NewSTORJ() (types.Storager, error) {
	return storj.NewStorager(
		// work_dir: https://beyondstorage.io/docs/go-storage/pairs/work_dir
		//
		// Relative operations will be based on this WorkDir.
		pairs.WithWorkDir(os.Getenv("STORAGE_STORJ_WORKDIR")),
		// endpoint: https://beyondstorage.io/docs/go-storage/pairs/credential
		//
		// Example Value: apikey:apikey_value
		pairs.WithCredential(os.Getenv("STORAGE_STORJ_CREDENTIAL")),
		// name: https://beyondstorage.io/docs/go-storage/pairs/name
		//
		// name is the bucket name.
		pairs.WithName(os.Getenv("STORAGE_STORJ_NAME")),
	)
}

func NewSTORJFromString() (types.Storager, error) {
	connStr := fmt.Sprintf(
		"storj://%s%s?credential=%s",
		os.Getenv("STORAGE_STORJ_NAME"),
		os.Getenv("STORAGE_STORJ_WORKDIR"),
		os.Getenv("STORAGE_STORJ_CREDENTIAL"),
	)
	return services.NewStoragerFromString(connStr)
}
