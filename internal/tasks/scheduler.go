package tasks

import (
	"fmt"

	"be0/internal/utils/logger"

	"github.com/hibiken/asynq"
)

// Scheduler handles periodic task scheduling
type Scheduler struct {
	scheduler *asynq.Scheduler
	logger    *logger.Logger
}

// NewScheduler creates a new task scheduler
func NewScheduler(redisAddr, username, password string, db int, logger *logger.Logger) *Scheduler {
	scheduler := asynq.NewScheduler(
		asynq.RedisClientOpt{
			Addr:     redisAddr,
			Username: username,
			Password: password,
			DB:       db,
		},
		&asynq.SchedulerOpts{},
	)

	return &Scheduler{
		scheduler: scheduler,
		logger:    logger,
	}
}

// Start starts the scheduler
func (s *Scheduler) Start() error {
	if err := s.registerTasks(); err != nil {
		return fmt.Errorf("failed to register tasks: %w", err)
	}

	s.logger.Info("starting task scheduler")
	return s.scheduler.Run()
}

// Stop stops the scheduler
func (s *Scheduler) Stop() {
	s.scheduler.Shutdown()
	s.logger.Info("task scheduler stopped")
}

// registerTasks registers all periodic tasks
func (s *Scheduler) registerTasks() error {
	s.logger.Info("registered all periodic tasks")
	return nil
}

// RegisterCustomTask registers a custom periodic task
func (s *Scheduler) RegisterCustomTask(spec string, taskType string, payload []byte, opts ...asynq.Option) error {
	entryID, err := s.scheduler.Register(spec, asynq.NewTask(taskType, payload, opts...))
	if err != nil {
		return fmt.Errorf("failed to register custom task: %w", err)
	}

	s.logger.Info("registered custom task %s %s %s", taskType, spec, entryID)
	return nil
}
