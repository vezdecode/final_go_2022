package main

import (
	"sync"
	"time"
	"vk/internal/lib"
)

var wg sync.WaitGroup

func main() {
	durations, err := lib.GetDuration()
	if err != nil {
		return
	}

	wg.Add(len(*durations))

	for _, v := range *durations {
		go func(vv time.Duration) {
			defer wg.Done()
			lib.RunTask(vv)
		}(v)
	}

	wg.Wait()
}
