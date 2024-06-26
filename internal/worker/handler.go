package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
)

const (
	TypeImageUpload = "image:upload"
	TypeImageDelete = "image:delete"
)

type ImageUploadPayload struct {
	Filename string
}

type ImageDeletePayload struct {
	ImageURL string
}

func HandleImageUploadTask(ctx context.Context, task *asynq.Task) error {
	var p ImageUploadPayload
	if err := json.Unmarshal(task.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	return nil
}

func HandleImageDeleteTask(ctx context.Context, task *asynq.Task) error {
	var p ImageDeletePayload
	if err := json.Unmarshal(task.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	return nil
}
