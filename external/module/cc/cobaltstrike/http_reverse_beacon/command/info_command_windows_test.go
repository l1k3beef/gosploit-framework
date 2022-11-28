package command

import (
	"os"
	"testing"
)

func TestLogin(t *testing.T) {
	initTest(t)
	cmd := &InfoCommand{}
	resp := cmd.Login()
	if resp == nil {
		println("login failed")
	} else {
		println("login sucess")
	}
}

func TestInfo(t *testing.T) {
	hostname, _ := os.Hostname()
	println(hostname)
}
