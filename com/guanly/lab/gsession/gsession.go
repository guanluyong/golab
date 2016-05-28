// Copyright 2016 By GuanLuyong.  All rights reserved.

// Package gsession is a Session Util for web
package gsession

import (
	. "com/guanly/lab/glog"
	"sync"
	"time"
)

type SessionManager struct {
	Lock     sync.Mutex
	Sessions map[string]Session
	Expires  int
}

type Session struct {
	Key     string
	Expires int
	Value   interface{}
}

var SM *SessionManager

func init() {
	Logger.Debug("Init SessionManager")

	SM = &SessionManager{Sessions: make(map[string]Session), Expires: 3600}

	SM.Listen()
}

func (s *SessionManager) Listen() {
	time.AfterFunc(time.Second, func() { s.Listen() })

	s.Lock.Lock()
	defer s.Lock.Unlock()

	for k, v := range s.Sessions {
		if v.Expires == 0 {
			delete(s.Sessions, k)
		} else {
			v.Expires--
		}
		Logger.Debug(v.Expires)
	}
}

func (s *SessionManager) Set(k string, v interface{}) {
	session := Session{Key: k, Expires: s.Expires, Value: v}
	delete(s.Sessions, k)
	s.Sessions[k] = session
}

func (s *SessionManager) Get(k string) interface{} {
	if session, ok := s.Sessions[k]; ok {
		return session.Value
	} else {
		return nil
	}
}
