package task

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/hibiken/asynq"
)

const (
	TypeCacheTest    = "test:cache"
	TypeDeleteTest   = "test:reminder"
	TypeScheduleTest = "test:schedule"
)

// PayloadCacheTest 定时任务所需参数
type PayloadCacheTest struct {
	Time time.Time
}

func NewTaskScheduleTest(t time.Time) (*asynq.Task, error) {
	payload, err := json.Marshal(PayloadCacheTest{Time: t})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeScheduleTest, payload), nil
}

func NewTaskCacheTest(t time.Time) (*asynq.Task, error) {
	payload, err := json.Marshal(PayloadCacheTest{Time: t})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeCacheTest, payload), nil
}

func NewTaskDeleteTest(t time.Time) (*asynq.Task, error) {
	payload, err := json.Marshal(PayloadCacheTest{Time: t})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeDeleteTest, payload), nil
}

func HandleTaskScheduleTask(ctx context.Context, t *asynq.Task) error {
	var p PayloadCacheTest
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}
	log.Printf(" [*] Run Schedule 任务 task time [%s]", p.Time.Format("2006-01-02 15:04:05"))
	return nil
}

func HandleTaskCacheTask(ctx context.Context, t *asynq.Task) error {
	var p PayloadCacheTest
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}
	log.Printf(" [*] Run 任务1 task time [%s]", p.Time.Format("2006-01-02 15:04:05"))
	return nil
}

func HandleTaskDeleteTest(ctx context.Context, t *asynq.Task) error {
	var p PayloadCacheTest
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}
	log.Printf(" [*] Run 任务2 task time [%s]", p.Time.Format("2006-01-02 15:04:05"))
	return nil
}
