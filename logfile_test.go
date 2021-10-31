package logfile_test

import (
	"fmt"

	"github.com/tonnytg/logfile"

	"os"
	"strings"
	"testing"
)

func TestLogging(t *testing.T) {
	logfile.Msg("CRITICAL", "system stop working")
	logfile.Msg("WARNING", "timezone is not correct")

	// remember to check if you change name of file
	filename := "activity.log"

	// try to read log file
	file, _ := os.ReadFile(filename)
	if file == nil {
		t.Error("File not found")
	}
	// print message like this: 2021/10/31 20:20:37 UTC CRITICAL: Error
	fmt.Println(string(file))

	// read content and check if it contains the string
	if !strings.Contains(string(file), "system stop working") {
		t.Error("system message not found in logfile")
	}
	if !strings.Contains(string(file), "timezone is not correct") {
		t.Error("timezone message error not found in logfile")
	}
}
