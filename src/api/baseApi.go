package api

import (
	"dst-admin-go/session"
	_ "dst-admin-go/session/memory"
)

var sessions *session.Manager

const CookieName = "token"

func init() {
	sessions = session.NewManager("memory", CookieName, 3600*5)
	go sessions.GC()
}

func Sessions() *session.Manager {
	return sessions
}
