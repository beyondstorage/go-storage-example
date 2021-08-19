package example

import (
	"fmt"
	"os"

	ftp "github.com/beyondstorage/go-service-ftp"
	"github.com/beyondstorage/go-storage/v4/pairs"
	"github.com/beyondstorage/go-storage/v4/services"
	"github.com/beyondstorage/go-storage/v4/types"
)

func NewFTP() (types.Storager, error) {
	return ftp.NewStorager(
		// credential: https://beyondstorage.io/docs/go-storage/pairs/credential
		//
		// Credential could be fetched from service's console.
		//
		// Example Value: basic:user:password
		pairs.WithCredential(os.Getenv("STORAGE_FTP_CREDENTIAL")),
		// endpoint: https://beyondstorage.io/docs/go-storage/pairs/endpoint
		//
		// Example Value: tcp:host:port
		pairs.WithEndpoint(os.Getenv("STORAGE_FTP_ENDPOINT")),
		// work_dir: https://beyondstorage.io/docs/go-storage/pairs/work_dir
		//
		// Relative operations will be based on this WorkDir.
		pairs.WithWorkDir(os.Getenv("STORAGE_FTP_WORKDIR")),
	)
}

func NewFTPFromString() (types.Storager, error) {
	str := fmt.Sprintf(
		"ftp:///%s?credential=%s&endpoint=%s",
		os.Getenv("STORAGE_FTP_WORKDIR"),
		os.Getenv("STORAGE_FTP_CREDENTIAL"),
		os.Getenv("STORAGE_FTP_ENDPOINT"),
	)
	return services.NewStoragerFromString(str)
}
