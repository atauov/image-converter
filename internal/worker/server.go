package worker

import (
	"github.com/atauov/image-converter/internal/config"
	"github.com/atauov/image-converter/internal/s3storage"
	"github.com/hibiken/asynq"
	"github.com/minio/minio-go/v7"
	"log"
)

type Worker struct {
	Router   *asynq.ServeMux
	Client   *asynq.Client
	AsynqSrv *asynq.Server
	S3       *minio.Client
}

func NewServer(cfg *config.Config) *Worker {
	redisClient := asynq.NewClient(asynq.RedisClientOpt{
		Addr: cfg.RedisServer.Address,
	})

	asynqServer := asynq.NewServer(
		asynq.RedisClientOpt{
			Addr: cfg.RedisServer.Address,
		},
		asynq.Config{
			Concurrency: cfg.Workers,
		})

	s3Client, err := s3storage.S3Connection(&cfg.S3Server)
	if err != nil {
		log.Println(err)
	}

	return &Worker{
		Client:   redisClient,
		AsynqSrv: asynqServer,
		S3:       s3Client,
	}
}

func (w *Worker) InitRoutes() *asynq.ServeMux {
	mux := asynq.NewServeMux()
	mux.HandleFunc(TypeImageUpload, w.HandleImageUploadTask)
	mux.HandleFunc(TypeImageDelete, w.HandleImageDeleteTask)

	return mux
}

func (w *Worker) Start(mux *asynq.ServeMux) error {
	return w.AsynqSrv.Run(mux)
}
