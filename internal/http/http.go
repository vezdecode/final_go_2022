package http

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
	"vk/internal/lib"
)

type Task struct {
	chDone chan int
	d      time.Duration
	isSync bool
}

var ch = make(chan int)
var schedule []Task

var CircleCtxCancel context.CancelFunc

func testCircle(ctx context.Context) {
	println("aasasa")
	for {
		if len(schedule) == 0 {
			log.Println("await new task")
			select {
			case _ = <-ch:
			}
		}

		if len(schedule) > 0 {
			lib.RunTask(schedule[0].d)

			if schedule[0].isSync {
				schedule[0].chDone <- 1
			}

			schedule = schedule[1:]
		}

	}
}

func ScheduleHandler(ctx *gin.Context) {
	keys := make([]time.Duration, len(schedule))

	for i, k := range schedule {
		keys[i] = k.d
		i++
	}

	ctx.JSON(200, keys)
}

func AddHandler(ctx *gin.Context) {
	isSync := ctx.Request.URL.Query().Has("sync")

	duration, err := time.ParseDuration(ctx.Request.URL.Query().Get("timeDuration"))
	if err != nil {
		return
	}
	var chDone chan int
	if isSync {
		chDone = make(chan int)
	}

	schedule = append(schedule, Task{chDone, duration, isSync})

	if len(schedule) == 1 {
		ch <- 1
	}

	if isSync {
		select {
		case _ = <-chDone:
		}
	}
}

func TimeHandler(ctx *gin.Context) {
	sum := time.Duration(0)
	for _, v := range schedule {
		sum += v.d
	}
	ctx.JSON(200, gin.H{
		"elevated_time": sum,
	})

}

func GetRoute() *gin.Engine {
	app := gin.Default()
	println("asa")
	schedule = make([]Task, 0)
	println("asa2")

	fmt.Println(schedule)
	if CircleCtxCancel != nil {
		CircleCtxCancel()
	}
	var ctx context.Context
	ctx, CircleCtxCancel = context.WithCancel(context.Background())
	go testCircle(ctx)

	app.POST("/add", AddHandler)
	app.GET("/schedule", ScheduleHandler)
	app.GET("/time", TimeHandler)

	return app
}
