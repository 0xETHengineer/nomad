package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestSentinelDeleteCommand_Implements(t *testing.T) {
	var _ cli.Command = &SentinelDeleteCommand{}
}
