package main

import (
	"vk/internal/lib"
)

func main() {
	durations, err := lib.GetDuration()
	if err != nil {
		return
	}

	for _, v := range *durations {
		lib.RunTask(v)
	}
}
