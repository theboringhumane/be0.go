package tasks

import (
	"fmt"
	"time"

	"be0/internal/utils/logger"

	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"
)

// TaskClient handles task enqueuing with improved error handling and context support
type TaskClient struct {
	client       *asynq.Client
	logger       *logger.Logger
	redisOptions *redis.Options
	redisClient  *redis.Client
}

type RateLimiter struct {
	Rate   int
	Burst  int
	Period time.Duration
}

func (c *TaskClient) GetClient() *asynq.Client {
	return c.client
}

// NewTaskClient creates a new TaskClient with the given Redis configuration
func NewTaskClient(redisAddr, username, password string, db int) *TaskClient {
	redisOpt := asynq.RedisClientOpt{
		Addr:     redisAddr,
		Username: username,
		Password: password,
		DB:       db,
	}

	redisClient := redis.NewClient(
		&redis.Options{
			Addr:     redisAddr,
			Username: username,
			Password: password,
			DB:       db,
		},
	)

	return &TaskClient{
		client: asynq.NewClient(redisOpt),
		redisOptions: &redis.Options{
			Addr:     redisAddr,
			Username: username,
			Password: password,
			DB:       db,
		},
		redisClient: redisClient,
		logger:      logger.New("TASKS"),
	}
}

// Close closes the underlying asynq client
func (c *TaskClient) Close() error {
	return c.client.Close()
}

func GetEmailQueueName(smtpSettingsID string) string {
	return fmt.Sprintf("email:smtp:%s", smtpSettingsID)
}
