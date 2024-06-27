package s3storage

import (
	"github.com/atauov/image-converter/internal/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"golang.org/x/net/context"
	"io"
)

const CONTENT_TYPE = "image/webp"

var bucketName string //TODO change on cfg

func S3Connection(cfg *config.S3Server) (*minio.Client, error) {
	s3Client, err := minio.New(cfg.Endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(cfg.AccessKey, cfg.SecretKey, ""),
	})
	if err != nil {
		return nil, err
	}
	bucketName = cfg.Bucket

	return s3Client, nil
}

func UploadToS3(s3Client *minio.Client, file io.Reader, size int64, name string) error {
	_, err := s3Client.PutObject(context.Background(), bucketName, name, file,
		size, minio.PutObjectOptions{ContentType: CONTENT_TYPE})

	if err != nil {
		return err
	}

	return nil
}

func DeleteFromS3(s3Client *minio.Client, filename string) error {
	if err := s3Client.RemoveObject(context.Background(), bucketName, filename,
		minio.RemoveObjectOptions{ForceDelete: true}); err != nil {
		return err
	}

	return nil
}
