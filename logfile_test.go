package logfile_test

import (
	"fmt"

	"github.com/tonnytg/logfile"

	"os"
	"strings"
	"testing"
)

func TestLogging(t *testing.T) {
	logfile.Msg("CRITICAL", "Error test logfile")

	// remember to check if you change name of file
	filename := "activity.log"

	// try to read log file
	if file, _ := os.ReadFile(filename); file == nil {
		t.Error("File not found")
	}
	// print message like this: 2021/10/31 20:20:37 UTC CRITICAL: Error
	fmt.Println(string(file))

	// read content and check if it contains the string
	if !strings.Contains(string(file), "Error test logfile") {
		t.Error("Error not found in logfile")
	}
}
