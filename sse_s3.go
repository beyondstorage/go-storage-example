package example

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"math/rand"

	s3 "github.com/aos-dev/go-service-s3"
	"github.com/aos-dev/go-storage/v3/pkg/randbytes"
	"github.com/aos-dev/go-storage/v3/types"
)

func WriteSseS3(store types.Storager, path string) {
	// content to write
	size := rand.Int63n(4 * 1024 * 1024)
	content, _ := ioutil.ReadAll(io.LimitReader(randbytes.NewRand(), size))
	r := bytes.NewReader(content)

	n, err := store.Write(path, r, size,
		// Required, must be AES256
		s3.WithServerSideEncryption(s3.ServerSideEncryptionAes256),
	)
	if err != nil {
		log.Fatalf("write %v: %v", path, err)
	}

	log.Printf("write size: %d", n)
}

func WriteSseKms(store types.Storager, path string, kmsKeyId string, context map[string]string, bucketKeyEnabled bool) {
	// content to write
	size := rand.Int63n(4 * 1024 * 1024)
	content, _ := ioutil.ReadAll(io.LimitReader(randbytes.NewRand(), size))
	r := bytes.NewReader(content)

	ctx, _ := json.Marshal(context)

	n, err := store.Write(path, r, size,
		// Required, must be aws:kms
		s3.WithServerSideEncryption(s3.ServerSideEncryptionAwsKms),
		// Required
		//
		// Example value: 1234abcd-12ab-34cd-56ef-1234567890ab
		s3.WithServerSideEncryptionAwsKmsKeyID(kmsKeyId),
		// Optional
		//
		// An encryption context is an optional set of key-value pairs that can contain additional contextual information about the data. https://docs.aws.amazon.com/AmazonS3/latest/userguide/UsingKMSEncryption.html#encryption-context
		s3.WithServerSideEncryptionContext(base64.StdEncoding.EncodeToString(ctx)),
		// Optional, S3 Bucket Key settings will be used if this is not specified.
		//
		// S3 Bucket Keys can reduce your AWS KMS request costs by decreasing the request traffic from Amazon S3 to AWS KMS. https://docs.aws.amazon.com/AmazonS3/latest/userguide/UsingKMSEncryption.html#sse-kms-bucket-keys
		s3.WithServerSideEncryptionBucketKeyEnabled(bucketKeyEnabled),
	)
	if err != nil {
		log.Fatalf("write %v: %v", path, err)
	}

	log.Printf("write size: %d", n)
}

func WriteSseC(store types.Storager, path string, key []byte) {
	// content to write
	size := rand.Int63n(4 * 1024 * 1024)
	content, _ := ioutil.ReadAll(io.LimitReader(randbytes.NewRand(), size))
	r := bytes.NewReader(content)

	n, err := store.Write(path, r, size,
		// Required, must be AES256
		s3.WithServerSideEncryptionCustomerAlgorithm(s3.ServerSideEncryptionAes256),
		// Required, your AES-256 key, a 32-byte binary value
		s3.WithServerSideEncryptionCustomerKey(key),
	)
	if err != nil {
		log.Fatalf("write %v: %v", path, err)
	}

	log.Printf("write size: %d", n)
}

func ReadSseC(store types.Storager, path string, key []byte) {
	var buf bytes.Buffer

	n, err := store.Read(path, &buf,
		// Required, must be AES256
		s3.WithServerSideEncryptionCustomerAlgorithm(s3.ServerSideEncryptionAes256),
		// Required, your AES-256 key, a 32-byte binary value
		s3.WithServerSideEncryptionCustomerKey(key),
	)
	if err != nil {
		log.Fatalf("read %v: %v", path, err)
	}

	log.Printf("read size: %d", n)
	log.Printf("read content: %s", buf.Bytes())
}
