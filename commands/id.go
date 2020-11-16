package commands

import (
	"errors"

	"github.com/xiaxiaosheng/go-imap"
)

// ID is a command, as defined in RFC 3501 section 6.2.2.
type ID struct {
	Params []interface{}
}

func (cmd *ID) Command() *imap.Command {
	return &imap.Command{
		Name:      "ID",
		Arguments: []interface{}{cmd.Params},
	}
}

func (cmd *ID) Parse(fields []interface{}) error {
	if len(fields) < 1 && len(fields)%2 != 0 {
		return errors.New("Not enough arguments")
	}

	cmd.Params = fields
	return nil
}
