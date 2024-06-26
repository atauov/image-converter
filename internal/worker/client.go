package worker

import (
	"github.com/atauov/image-converter/internal/config"
	"github.com/hibiken/asynq"
)

type Client struct {
	Client *asynq.Client
}

func NewClient(cfg *config.RedisServer) *Client {
	client := asynq.NewClient(asynq.RedisClientOpt{
		Addr: cfg.Address,
	})

	return &Client{Client: client}
}

func (c *Client) Close() {
	c.Client.Close()
}
