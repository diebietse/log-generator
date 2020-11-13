package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_loadLogs(t *testing.T) {
	expectedLogLines := []string{"A plaintext log", `{"message": "A JSON log"}`, `{"message": "{\"escaped\":\"An escaped JSON log\"}"}`}
	logLines, err := loadLogs("example_logs.txt")
	assert.NoError(t, err)
	assert.Equal(t, expectedLogLines, logLines)
}

func Test_loadLogsFileNotFound(t *testing.T) {
	_, err := loadLogs("some_non_existing_file")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no such file or directory")
}
