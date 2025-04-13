package keeper_datasource

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

// RunPackerAcceptanceTest runs a Packer acceptance test and checks the logs for expected lines.
func RunPackerAcceptanceTest(t *testing.T, buildCommand *exec.Cmd, logfile string, expectedLogLines []string) error {
	if buildCommand.ProcessState != nil {
		if buildCommand.ProcessState.ExitCode() != 0 {
			return fmt.Errorf("Bad exit code. Logfile: %s", logfile)
		}
	}

	logsBytes, err := os.ReadFile(logfile)
	if err != nil {
		return fmt.Errorf("Unable to read %s", logfile)
	}

	logsString := string(logsBytes)
	for _, line := range expectedLogLines {
		matched, _ := regexp.MatchString(line+".*", logsString)
		assert.True(t, matched, "logs doesn't contain expected value %s", line)
	}

	return nil
}