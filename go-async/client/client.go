package main

import (
	"log"
	"time"

	"github.com/hibiken/asynq"

	"go-async/task"
)

func main() {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}

	// 立即执行任务和延时执行任务案例
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: "127.0.0.1:36379", Password: "G62m50oigInC30sf", DB: 0})

	// 任务一：立即执行任务 t2
	t2, err := task.NewTaskCacheTest(time.Now())
	if err != nil {
		log.Fatal(err)
	}

	info, err := client.Enqueue(t2, asynq.Queue("normal"))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf(" [*] Successfully enqueued task1: %+v", info)

	// 任务二：3秒后执行任务 t3
	t3, err := task.NewTaskDeleteTest(time.Now())
	if err != nil {
		log.Fatal(err)
	}

	info, err = client.Enqueue(t3, asynq.Queue("critical"), asynq.ProcessIn(3*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf(" [*] Successfully enqueued task2: %+v", info)

	// 任务三：指定时间执行任务 t4
	t4, err := task.NewTaskDeleteTest(time.Now())
	if err != nil {
		log.Fatal(err)
	}

	info, err = client.Enqueue(t4, asynq.Queue("low"), asynq.ProcessAt(time.Date(2024, time.January, 30, 13, 49, 00, 0, loc)))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf(" [*] Successfully enqueued task3: %+v", info)
}
