package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jessevdk/go-flags"
)

type config struct {
	LogSourceFilePath string        `long:"log_source_file_path" env:"LOG_SOURCE_FILE_PATH" default:"example_logs.txt" description:"Location to load the file containing the data we want to log"`
	LogDelay          time.Duration `long:"log_delay" env:"LOG_DELAY" default:"1s" description:"Time to wait before printing a log line"`
}

func main() {
	conf := &config{}
	parser := flags.NewParser(conf, flags.Default)
	if _, err := parser.Parse(); err != nil {
		log.Fatalf("Could not load config: %v\n", err)
	}

	ticker := time.NewTicker(conf.LogDelay)
	logLines, err := loadLogs(conf.LogSourceFilePath)
	if err != nil {
		log.Fatalf("Could not load logs: %v\n", err)
	}
	line := 0
	for range ticker.C {
		fmt.Println(logLines[line])
		line++
		line %= len(logLines)
	}
}

func loadLogs(logSourceFilePath string) ([]string, error) {
	f, err := os.Open(logSourceFilePath)
	if err != nil {
		return nil, fmt.Errorf("could not open file '%s': %w", logSourceFilePath, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	logLines := []string{}
	for scanner.Scan() {
		logLines = append(logLines, scanner.Text())
	}

	return logLines, nil
}
