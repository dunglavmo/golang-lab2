package main

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	redis "github.com/go-redis/redis/v8"
)

type TaskCache struct {
	client *redis.Client
}

func NewTaskCache() *TaskCache {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	return &TaskCache{client: client}
}

func (tc *TaskCache) AddTask(task Task) error {
	ctx := context.Background()
	taskKey := fmt.Sprintf("task:%d", task.ID)
	taskValue := fmt.Sprintf("%s|%d|%s", task.Name, task.Status, task.CreatedAt)

	err := tc.client.Set(ctx, taskKey, taskValue, 60*time.Second).Err()
	return err
}

func (tc *TaskCache) GetTask(id int) (Task, error) {
	ctx := context.Background()
	taskKey := fmt.Sprintf("task:%d", id)

	taskValue, err := tc.client.Get(ctx, taskKey).Result()
	if err == redis.Nil {
		return Task{}, fmt.Errorf("task with ID %d not found", id)
	} else if err != nil {
		return Task{}, err
	}

	parts := strings.Split(taskValue, "|")
	if len(parts) != 3 {
		return Task{}, fmt.Errorf("invalid task data in cache for ID %d", id)
	}

	status, _ := strconv.Atoi(parts[1])

	task := Task{
		ID:        id,
		Name:      parts[0],
		Status:    status,
		CreatedAt: parts[2],
	}

	return task, nil
}

func (tc *TaskCache) ListTasks() ([]Task, error) {
	ctx := context.Background()
	keys, err := tc.client.Keys(ctx, "task:*").Result()
	if err != nil {
		return nil, err
	}

	var tasks []Task
	for _, key := range keys {
		id, _ := strconv.Atoi(strings.TrimPrefix(key, "task:"))
		task, err := tc.GetTask(id)
		if err == nil {
			tasks = append(tasks, task)
		}
	}

	return tasks, nil
}

func (tc *TaskCache) UpdateTask(task Task) error {
	return tc.AddTask(task)
}

func (tc *TaskCache) DeleteTask(id int) error {
	ctx := context.Background()
	taskKey := fmt.Sprintf("task:%d", id)
	return tc.client.Del(ctx, taskKey).Err()
}
