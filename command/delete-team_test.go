package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestDeleteTeamCommand_implement(t *testing.T) {
	var _ cli.Command = &DeleteTeamCommand{}
}
