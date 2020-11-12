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
	LogSourceFilePath string `long:"log_source_file_path" env:"LOG_SOURCE_FILE_PATH" default:"example_logs.txt" description:"Location to load the file containing the data we want to log"`
	LogDelaySeconds   int    `long:"log_delay_seconds" env:"LOG_DELAY_SECONDS" default:"1" description:"Time to wait before printing a log line"`
}

func main() {
	conf := &config{}
	parser := flags.NewParser(conf, flags.Default)
	if _, err := parser.Parse(); err != nil {
		log.Fatalf("Could not load config: %v\n", err)
	}

	ticker := time.NewTicker(time.Duration(conf.LogDelaySeconds) * time.Second)
	logLines := loadLogs(conf.LogSourceFilePath)
	line := 0
	for range ticker.C {
		fmt.Println(logLines[line])
		line++
		line %= len(logLines)
	}
}

func loadLogs(logSourceFilePath string) []string {
	f, err := os.Open(logSourceFilePath) // os.OpenFile has more options if you need them
	if err != nil {
		log.Fatalf("Could not open file '%s': %v\n", logSourceFilePath, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	logLines := []string{}
	for scanner.Scan() {
		logLines = append(logLines, scanner.Text())
	}

	return logLines
}
