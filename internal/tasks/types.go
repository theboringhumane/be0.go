package tasks

import "time"

// Task Types
const (
	// Queue related tasks
	TaskTypeQueueConfig = "queue:config"
)

// Task Queues
const (
	QueueCritical = "critical" // For time-sensitive tasks like email sending
	QueueDefault  = "default"  // For regular tasks
	QueueLow      = "low"      // For background tasks like cleanup
)

// Task Priorities (1-10, higher is more important)
const (
	PriorityCritical = 10
	PriorityHigh     = 8
	PriorityNormal   = 5
	PriorityLow      = 3
	PriorityBG       = 1
)

// Task Timeouts
const (
	TimeoutShort  = 1 * time.Minute
	TimeoutMedium = 5 * time.Minute
	TimeoutLong   = 30 * time.Minute
)

// Task Retry Settings
const (
	RetryMax     = 5
	RetryDefault = 3
	RetryMin     = 1
)
