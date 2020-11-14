package commands

import (
	"errors"
	"fmt"
	"strings"

	"github.com/xiaxiaosheng/go-imap"
)

// ID is a command, as defined in RFC 3501 section 6.2.2.
type ID struct {
	IAM map[string]string
}

func (cmd *ID) Command() *imap.Command {
	arr := make([]interface{}, 0, 2*len(cmd.IAM))
	isFirst := true
	for key, value := range cmd.IAM {
		if isFirst {
			arr = append(arr, "("+key)
			isFirst = false
		} else {
			arr = append(arr, key)
		}
		arr = append(arr, value)
	}
	arr[len(arr)-1] = fmt.Sprintf("%v)", arr[len(arr)-1])

	return &imap.Command{
		Name:      "ID",
		Arguments: arr,
	}
}

func (cmd *ID) Parse(fields []interface{}) error {
	if len(fields) < 2 || len(fields)%2 != 0 {
		return errors.New("Not enough arguments")
	}

	cmd.IAM = make(map[string]string)
	for i := 0; i < len(fields); i++ {
		var s1 string
		var s2 string
		s1 = fmt.Sprintf("%v", fields[i])
		s2 = fmt.Sprintf("%v", fields[i+1])
		if i == 0 && strings.HasPrefix(s1, "(") {
			s1 = strings.TrimPrefix(s1, "(")
		}
		if i == len(fields)-1 && strings.HasSuffix(s2, ")") {
			s2 = strings.TrimSuffix(s2, ")")
		}
		cmd.IAM[s1] = s2
	}

	return nil
}
