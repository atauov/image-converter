package worker

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/atauov/image-converter/internal/s3storage"
	"github.com/hibiken/asynq"
	"log"
)

const (
	TypeImageUpload      = "image:upload"
	TypeImageDelete      = "image:delete"
	TypeLocalImageDelete = "local:image:delete"
)

type ImageUploadPayload struct {
	//ImageData []byte // send and receive img to queue w/o saving
	Filename string
}

type ImageDeletePayload struct {
	ImageURL string
}

type ImageLocalDeletePayload struct {
	Filename string
}

func (w *Worker) HandleImageUploadTask(ctx context.Context, task *asynq.Task) error {
	var p ImageUploadPayload

	if err := json.Unmarshal(task.Payload(), &p); err != nil {
		log.Println(err)

		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	img, err := loadImage(p.Filename)
	if err != nil {
		log.Println(err)
		return fmt.Errorf("can not load image: %v: %w", err, asynq.SkipRetry)
	}

	imgBuffs := make(map[uint][]byte)

	if err = collectBuffs(img, imgBuffs); err != nil {
		log.Println(err)
		return fmt.Errorf("can not load image: %v: %w", err, asynq.SkipRetry)
	}

	for size, buff := range imgBuffs {
		newName := modifyFilename(p.Filename, size)

		sizeOfFile := int64(len(buff))
		reader := bytes.NewReader(buff)

		if err = s3storage.UploadToS3(w.S3, reader, sizeOfFile, newName); err != nil {
			log.Println(err)
			return fmt.Errorf("uploading to s3 failed: %v: %w", err, asynq.SkipRetry)
		}

		log.Println("Successfully uploaded image", newName)
	}

	if err = deleteLocalFile(p.Filename); err != nil {
		log.Println(err)
	} else {
		log.Println("Successfully deleted image", p.Filename)
	}

	return nil
}

func (w *Worker) HandleImageDeleteTask(ctx context.Context, task *asynq.Task) error {
	var p ImageDeletePayload
	if err := json.Unmarshal(task.Payload(), &p); err != nil {
		log.Println(err)
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	files := getFilesFromUrl(p.ImageURL)

	for _, file := range files {
		fmt.Println(file)
		if err := s3storage.DeleteFromS3(w.S3, file); err != nil {
			log.Println(err)
			return fmt.Errorf("deleting from s3 failed: %v: %w", err, asynq.SkipRetry)
		}

		log.Println("Successfully deleted image", file)
	}

	return nil
}
