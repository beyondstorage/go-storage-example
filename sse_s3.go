package example

import (
	"encoding/base64"
	"encoding/json"
	"os"

	s3 "github.com/beyondstorage/go-service-s3/v2"
	"github.com/beyondstorage/go-storage/v4/pairs"
	"github.com/beyondstorage/go-storage/v4/types"
)

func S3Pairs() []types.Pair {
	return []types.Pair{
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
		pairs.WithName(os.Getenv("STORAGE_S3_NAME"))}
}

func NewS3SseS3() (types.Storager, error) {
	return s3.NewStorager(
		append(S3Pairs(),
			s3.WithDefaultStoragePairs(s3.DefaultStoragePairs{
				Write: []types.Pair{
					// Required, must be AES256
					s3.WithServerSideEncryption(s3.ServerSideEncryptionAes256),
				},
			}))...,
	)
}

func NewS3SseKms(keyId string, context map[string]string, bucketKeyEnabled bool) (types.Storager, error) {
	ctx, _ := json.Marshal(context)

	return s3.NewStorager(
		append(S3Pairs(),
			s3.WithDefaultStoragePairs(s3.DefaultStoragePairs{
				Write: []types.Pair{
					// Required, must be aws:kms
					s3.WithServerSideEncryption(s3.ServerSideEncryptionAwsKms),
					// Required
					//
					// Example value: 1234abcd-12ab-34cd-56ef-1234567890ab
					s3.WithServerSideEncryptionAwsKmsKeyID(keyId),
					// Optional
					//
					// An encryption context is an optional set of key-value pairs that can contain additional contextual information about the data. https://docs.aws.amazon.com/AmazonS3/latest/userguide/UsingKMSEncryption.html#encryption-context
					s3.WithServerSideEncryptionContext(base64.StdEncoding.EncodeToString(ctx)),
					// Optional, S3 Bucket Key settings will be used if this is not specified.
					//
					// S3 Bucket Keys can reduce your AWS KMS request costs by decreasing the request traffic from Amazon S3 to AWS KMS. https://docs.aws.amazon.com/AmazonS3/latest/userguide/UsingKMSEncryption.html#sse-kms-bucket-keys
					s3.WithServerSideEncryptionBucketKeyEnabled(bucketKeyEnabled),
				},
			}))...,
	)
}

func NewS3SseC(key []byte) (types.Storager, error) {
	return s3.NewStorager(
		append(S3Pairs(),
			s3.WithDefaultStoragePairs(s3.DefaultStoragePairs{
				Write: []types.Pair{
					// Required, must be AES256
					s3.WithServerSideEncryptionCustomerAlgorithm(s3.ServerSideEncryptionAes256),
					// Required, your AES-256 key, a 32-byte binary value
					s3.WithServerSideEncryptionCustomerKey(key),
				},
				// Now you have to provide customer key to read encrypted data
				Read: []types.Pair{
					// Required, must be AES256
					s3.WithServerSideEncryptionCustomerAlgorithm(s3.ServerSideEncryptionAes256),
					// Required, your AES-256 key, a 32-byte binary value
					s3.WithServerSideEncryptionCustomerKey(key),
				},
			}))...,
	)
}
