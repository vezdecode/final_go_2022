package main

import (
	"fmt"
	"time"
	"vk/internal/lib"
)

func main() {
	var maxTask int

	_, err := fmt.Scanf("%d", &maxTask)
	if err != nil {
		return
	}

	durations, err := lib.GetDuration()
	if err != nil {
		return
	}

	ch := make(chan int)

	ranTask := 0

	for _, v := range *durations {
		if ranTask == maxTask {
			select {
			case <-ch:
				ranTask -= 1
			}
		}

		ranTask += 1
		go func(vv time.Duration) {
			lib.RunTask(vv)
			ch <- 1
		}(v)
	}
}
