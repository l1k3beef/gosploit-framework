package profile

import "time"

var (
	SleepTime  = 200 * time.Millisecond
	SessionKey = []byte{}
	RandomIV   = []byte{}
	SessionID  string
	Header     map[string]string
	ProxyUrl   string

	GetUrl        = ""
	PostUrl       = ""
	SessionFormat = "JSESSIONID=%v"
	UserAgent     = ""
)
