package main

import (
	"fmt"
	"log"

	"github.com/hibiken/asynq"

	"go-async/task"
)

func main() {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: "127.0.0.1:36379", Password: "G62m50oigInC30sf", DB: 0},
		asynq.Config{Concurrency: 20, Queues: map[string]int{
			"critical":  4,
			"normal":    2,
			"cron_test": 2,
			"default":   1,
			"low":       1,
		}, IsFailure: func(err error) bool {
			fmt.Printf("asynq server exec task IsFailure ======== >>>>>>>>>>>  err : %+v \n", err)
			return true
		}},
	)

	mux := asynq.NewServeMux()
	mux.HandleFunc(task.TypeCacheTest, task.HandleTaskCacheTask)
	mux.HandleFunc(task.TypeDeleteTest, task.HandleTaskDeleteTest)
	mux.HandleFunc(task.TypeScheduleTest, task.HandleTaskScheduleTask)

	if err := srv.Run(mux); err != nil {
		log.Fatal(err)
	}
}
