package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/hibiken/asynq"

	"go-async/task"
)

func main() {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}

	scheduler := asynq.NewScheduler(
		&asynq.RedisClientOpt{
			Addr:     "127.0.0.1:36379",
			Password: "G62m50oigInC30sf",
			DB:       0,
		},
		&asynq.SchedulerOpts{
			Location:        loc,
			PostEnqueueFunc: handleEnqueueError,
		},
	)

	t1, err := task.NewTaskScheduleTest(time.Now())
	if err != nil {
		log.Fatal(err)
	}

	// 每分钟执行一次任务 t1 参数固定，会保存到redis中
	entryID, err := scheduler.Register("* * * * *", t1, asynq.Queue("cron_test"))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("registered an entry: %q\n", entryID)
	// 运行
	if err := scheduler.Run(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func handleEnqueueError(task *asynq.TaskInfo, err error) {
	fmt.Println("task:", task)
	fmt.Println("err:", err)
}
