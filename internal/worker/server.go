package worker

import (
	"github.com/atauov/image-converter/internal/config"
	"github.com/hibiken/asynq"
)

type Worker struct {
	Router   *asynq.ServeMux
	Client   *asynq.Client
	AsynqSrv *asynq.Server
}

func NewServer(cfg *config.RedisServer) *Worker {
	redisClient := asynq.NewClient(asynq.RedisClientOpt{
		Addr: cfg.Address,
	})

	asynqServer := asynq.NewServer(
		asynq.RedisClientOpt{
			Addr: cfg.Address,
		},
		asynq.Config{
			Concurrency: cfg.Workers,
		})

	mux := asynq.NewServeMux()
	mux.HandleFunc(TypeImageUpload, HandleImageUploadTask)
	mux.HandleFunc(TypeImageDelete, HandleImageDeleteTask)

	return &Worker{
		Router:   mux,
		Client:   redisClient,
		AsynqSrv: asynqServer,
	}
}

func (w *Worker) Start() error {
	return w.AsynqSrv.Run(asynq.NewServeMux())
}
