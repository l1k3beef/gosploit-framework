package command

import (
	"encoding/json"
	"errors"
)

type DynmaicCommand struct {
	*Commnad
}

func (c *DynmaicCommand) Unmarshall(data []byte) error {
	err := json.Unmarshal(data, &c)
	if err != nil {
		return errors.New("unmarshall failed")
	}
	return nil
}
