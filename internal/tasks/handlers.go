package tasks

import (
	"be0/internal/config"
	"be0/internal/utils"
	"be0/internal/utils/logger"

	"gorm.io/gorm"
)

var (
	cfg, _ = config.Load()
)

// TaskHandler handles task processing with improved error handling and logging
type TaskHandler struct {
	db             *gorm.DB
	logger         *logger.Logger
	taskClient     *TaskClient
	storageHandler *utils.StorageHandler
}

// NewTaskHandler creates a new TaskHandler
func NewTaskHandler(db *gorm.DB) *TaskHandler {
	return &TaskHandler{
		db:             db,
		logger:         logger.New("task_handler"),
		taskClient:     NewTaskClient(cfg.Redis.Addr, cfg.Redis.Username, cfg.Redis.Password, cfg.Redis.DB),
		storageHandler: utils.NewStorageHandler(),
	}
}
