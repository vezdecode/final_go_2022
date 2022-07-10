package lib

import (
	"bufio"
	"flag"
	"log"
	"os"
	"time"
)

func GetDuration() (*[]time.Duration, error) {
	durations := make([]time.Duration, 0)

	r, err := ScanFile()

	if err != nil {
		return nil, err
	}

	for _, v := range r {
		duration, err := time.ParseDuration(v)
		if err != nil {
			return nil, err
		}
		durations = append(durations, duration)
	}

	return &durations, nil
}

func RunTask(duration time.Duration) {
	log.Printf("start task with duration %v\n", duration)
	time.Sleep(duration)
	log.Printf("complete task with duration %v\n", duration)
}

func ScanFile() ([]string, error) {
	fileName := flag.String("file", "", "file for get source, by default ./examples/data.txt")
	flag.Parse()
	if *fileName == "" {
		*fileName = "examples/data.txt"
	}

	file, err := os.Open(*fileName)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	lines := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			lines = append(lines, line)
		}
	}

	err = scanner.Err()
	if err != nil {
		return nil, err
	}
	return lines, nil
}
