package command

import (
	"encoding/json"
	"errors"
)

type ShellCommand struct {
	*Commnad

	Argv []string
	Envp []string
}

func (c *ShellCommand) Unmarshall(data []byte) error {
	err := json.Unmarshal(data, &c)
	if err != nil {
		return errors.New("unmarshall failed")
	}
	return nil
}
