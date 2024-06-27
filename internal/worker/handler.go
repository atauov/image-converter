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
	TypeImageUpload = "image:upload"
	TypeImageDelete = "image:delete"
)

type ImageUploadPayload struct {
	//ImageData []byte // send and receive img to queue w/o saving
	Filename string
}

type ImageDeletePayload struct {
	ImageURL string
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
		var newName string
		if size != 0 {
			newName = modifyFilename(p.Filename, size)
		}

		sizeOfFile := int64(len(buff))
		reader := bytes.NewReader(buff)

		if err = s3storage.UploadToS3(w.S3, reader, sizeOfFile, newName); err != nil {
			log.Println(err)
			return fmt.Errorf("upload to s3 failed: %v: %w", err, asynq.SkipRetry)
		}

		log.Println("Successfully uploaded image", newName)
	}

	return nil
}

func (w *Worker) HandleImageDeleteTask(ctx context.Context, task *asynq.Task) error {
	var p ImageDeletePayload
	if err := json.Unmarshal(task.Payload(), &p); err != nil {
		log.Println(err)

		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	return nil
}
