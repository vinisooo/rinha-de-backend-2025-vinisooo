package services

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type QueueService struct {
	client *redis.Client
}

type JobPayload struct {
	Type string
	Data interface{}
}

func NewQueueService(client *redis.Client) *QueueService {
	return &QueueService{
		client: client,
	}
}

func (q *QueueService) AddJob(ctx context.Context, queueName string, jobType string, data interface{}) error {
	payload := JobPayload{
		Data: data,
		Type: jobType,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("error on parsing json: %w", err)
	}

	err = q.client.LPush(ctx, queueName, jsonData).Err()
	if err != nil {
		return fmt.Errorf("something went wrong when adding data to queue %w", err)
	}

	return nil
}

func (q *QueueService) GetJob(ctx context.Context, queueName string) (*JobPayload, error) {
	result, err := q.client.BRPop(ctx, 0, queueName).Result()

	if err != nil {
		return nil, fmt.Errorf("error on geting result from queue: %w", err)
	}

	if len(result) < 2 {
		return nil, fmt.Errorf("invalid result from queue")
	}

	var payload JobPayload
	err = json.Unmarshal([]byte(result[1]), &payload)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal job payload: %w", err)
	}

	return &payload, nil
}
