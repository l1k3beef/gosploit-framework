package command

import "time"

type Command struct {
	Operate string
}

type ShellCommand struct {
	*Command
	Argv []string
	Envp []string
}

type SleepCommand struct {
	*Command
	SleepTime time.Duration
}
