package profile

import "time"

var (
	SleepTime  = 1 * time.Millisecond
	SessionKey = []byte{}
	RandomIV   = []byte{}
	SessionID  string
	Header     map[string]string
	ProxyUrl   string

	GetUrl        = ""
	PostUrl       = ""
	SessionFormat = "PHPSession=%v"
	UserAgent     = ""
)
