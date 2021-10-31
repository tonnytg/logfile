package logfile_test

import (
	"fmt"
	"github.com/tonnytg/logfile"
	"os"
	"strings"
	"testing"
)

func TestLogging(t *testing.T) {
	logfile.Msg("CRITICAL", "Error")

	file, _ := os.ReadFile("activity.log")
	if file == nil {
        t.Error("File not found")
    }
	fmt.Println(string(file))
	// read content and check if it contains the string
	if !strings.Contains(string(file), "Error") {
        t.Error("Error not found in logfile")
    }
}
