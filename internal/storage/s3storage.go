package storage

import (
	"github.com/atauov/image-converter/internal/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func S3Connection(cfg *config.S3Server) (*minio.Client, error) {
	s3Client, err := minio.New(cfg.Endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(cfg.AccessKey, cfg.SecretKey, ""),
	})
	if err != nil {
		return nil, err
	}

	return s3Client, nil
}
