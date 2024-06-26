package service

import (
	"github.com/atauov/image-converter/internal/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type S3Service struct {
	S3Client *minio.Client
	S3Bucket string
}

func S3Connection(cfg *config.S3Server) (*minio.Client, error) {
	s3Client, err := minio.New(cfg.Endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(cfg.AccessKey, cfg.SecretKey, ""),
	})
	if err != nil {
		return nil, err
	}

	return s3Client, nil
}

func (s3 *S3Service) PutObject() error {
	return nil
}

func (s3 *S3Service) DeleteObject(filename string) error {
	return nil
}
