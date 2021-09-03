package example

import (
	"fmt"
	"os"

	gdrive "github.com/beyondstorage/go-service-gdrive"
	"github.com/beyondstorage/go-storage/v4/pairs"
	"github.com/beyondstorage/go-storage/v4/services"
	"github.com/beyondstorage/go-storage/v4/types"
)

func NewGdrive() (types.Storager, error) {
	return gdrive.NewStorager(
		// credential: https://beyondstorage.io/docs/go-storage/pairs/credential
		//
		// Example Value: file:<abs_path_of_credential>
		pairs.WithCredential(os.Getenv("STORAGE_GDRIVE_CREDENTIAL")),
		// name: https://beyondstorage.io/docs/go-storage/pairs/name
		//
		// name is the bucket name.
		pairs.WithName(os.Getenv("STORAGE_GDRIVE_NAME")),
		// work_dir: https://beyondstorage.io/docs/go-storage/pairs/work_dir
		//
		// Relative operations will be based on this WorkDir.
		pairs.WithWorkDir(os.Getenv("STORAGE_GDRIVE_WORKDIR")),
	)
}

func NewGdriveFromString() (types.Storager, error) {
	str := fmt.Sprintf(
		"gdrive://%s/%s?credential=%s",
		os.Getenv("STORAGE_GDRIVE_NAME"),
		os.Getenv("STORAGE_GDRIVE_WORKDIR"),
		os.Getenv("STORAGE_GDRIVE_CREDENTIAL"),
	)
	return services.NewStoragerFromString(str)
}
