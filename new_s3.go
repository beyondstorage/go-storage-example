package example

import (
	"fmt"
	"os"

	s3 "go.beyondstorage.io/services/s3/v3"
	"go.beyondstorage.io/v5/pairs"
	"go.beyondstorage.io/v5/services"
	"go.beyondstorage.io/v5/types"
)

func NewS3() (types.Storager, error) {
	return s3.NewStorager(
		// work_dir: https://beyondstorage.io/docs/go-storage/pairs/work_dir
		//
		// Relative operations will be based on this WorkDir.
		pairs.WithWorkDir(os.Getenv("STORAGE_S3_WORKDIR")),
		// credential: https://beyondstorage.io/docs/go-storage/pairs/credential
		//
		// Credential could be fetched from service's console.
		//
		// Example Value: hmac:access_key_id:secret_access_key
		pairs.WithCredential(os.Getenv("STORAGE_S3_CREDENTIAL")),
		// endpoint: https://beyondstorage.io/docs/go-storage/pairs/endpoint
		//
		// endpoint is default to amazon s3's endpoint.
		// If using s3 compatible services, please input their endpoint.
		//
		// Example Value: https:host:port
		pairs.WithEndpoint(os.Getenv("STORAGE_S3_ENDPOINT")),
		// location: https://beyondstorage.io/docs/go-storage/pairs/location
		//
		// For s3, location is the bucket's zone.
		// For s3 compatible services, location could be ignored or has other value,
		// please refer to their documents.
		//
		// Example Value: ap-east-1
		pairs.WithLocation(os.Getenv("STORAGE_S3_LOCATION")),
		// name: https://beyondstorage.io/docs/go-storage/pairs/name
		//
		// name is the bucket name.
		pairs.WithName(os.Getenv("STORAGE_S3_NAME")),
		// features: https://beyondstorage.io/docs/go-storage/pairs/index#feature-pairs
		//
		// virtual_dir feature is designed for a service that doesn't have native dir support but wants to provide simulated operations.
		s3.WithEnableVirtualDir(),
	)
}

func NewS3FromString() (types.Storager, error) {
	connStr := fmt.Sprintf(
		"s3://%s%s?credential=%s&endpoint=%s&location=%s&enbale_virtual_dir",
		os.Getenv("STORAGE_S3_NAME"),
		os.Getenv("STORAGE_S3_WORKDIR"),
		os.Getenv("STORAGE_S3_CREDENTIAL"),
		os.Getenv("STORAGE_S3_ENDPOINT"),
		os.Getenv("STORAGE_S3_LOCATION"),
	)
	return services.NewStoragerFromString(connStr)
}
