package cobaltstrike

import (
	"testing"
	"time"
)

func TestHttpReverseBeacon(t *testing.T) {
	cc := &HttpReverseBeacon{}
	cc.Beacon = &Beacon{}
	cc.SessionKey = []byte("1234567812345678")
	cc.RandomIV = []byte("1234567812345678")
	cc.Profile = &HttpReverseProfile{
		BeaconAddr:  "localhost:80",
		PostUrl:     "/post",
		GetUrl:      "/get",
		SessionName: "JSESSIONID",
		SleepTime:   200,
	}
	cc.Run()
	for {
		time.Sleep(200)
	}
}
