package bot

import (
	"os"
	"testing"

	"github.com/xanzy/go-gitlab"
)

func TestPrintAvailCmds(t *testing.T) {
	// Want to ensure that we are rendering a teplnate without errors.
	// Super basic test, just want to make sure something comes out.
	cwd, _ := os.Getwd()
	err := os.Chdir(os.Dir(cwd))
	msg, err := formatAvailCommands()
	if err != nil {
		t.Errorf("Error printing available commands: %s", err)
	}

	if len(msg) == 0 {
		t.Errorf("Available commands template not rendered correctly")
	}
}
