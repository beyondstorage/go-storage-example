# Examples for go-storage

## Init Service/Storage

- [Create bos Storager](new_bos.go) (Baidu Object Storage)
- [Create cos Storager](new_cos.go) (Tencent Cloud Object Storage)
- [Create fs Storager](new_fs.go)
- [Create ftp Storager](new_ftp.go) (File Transfer Protocol)
- [Create hdfs Storager](new_hdfs.go) (Hadoop Distributed File System)
- [Create ipfs Storager](new_ipfs.go) (InterPlanetary File System)
- [Create minio Storager](new_minio.go) (MinIO)
- [Create gdrive Storager](new_gdrive.go) (Google Drive)
- [Create obs Storager](new_obs.go) (Huawei Object Storage Service)
- [Create s3 Storager](new_s3.go) (Amazon S3)
- [Create storj Storager](new_storj.go) (Storj)
## Basic Operations

### List files

- [List all files](list.go)
- [List a directory](list.go)
- [List with prefix](list.go)
- [List multipart uploads](list.go)

### Read file

- [Read a file](read.go)
- [Read a range of a file](read.go)
- [Read a file with callback](read.go)
- [Read a file using signed URL](read.go)

### Write file

Write a single file.

- [Write a file](write.go)
- [Write a file with callback](write.go)
- [Write a file using signed URL](write.go)

Write file via append.

- [Append to a new file](append.go)
- [Append to an existing file](append.go)

Write file via multipart.

- [Multipart upload](multipart.go)
- [Resume a multipart upload](multipart.go)
- [Cancel a multipart upload](multipart.go)
