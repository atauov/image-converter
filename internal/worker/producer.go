package worker

import (
	"encoding/json"
	"github.com/hibiken/asynq"
)

func NewImageUploadTask(filename string) (*asynq.Task, error) {
	payload, err := json.Marshal(ImageUploadPayload{
		Filename: filename},
	)
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TypeImageUpload, payload), nil
}

func NewImageDeleteTask(url string) (*asynq.Task, error) {
	payload, err := json.Marshal(ImageDeletePayload{ImageURL: url})
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TypeImageDelete, payload), nil
}

func NewLocalImageDeleteTask(filename string) (*asynq.Task, error) {
	payload, err := json.Marshal(ImageLocalDeletePayload{filename})
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TypeLocalImageDelete, payload), nil
}
