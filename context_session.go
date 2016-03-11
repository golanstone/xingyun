package xingyun

import (
	"time"
)

var (
	SessionKey   string = "ZQSESSID"
	sessionIDLen int    = 36
)

func newSessionID() string {
	return GenRandString(sessionIDLen)
}

func (ctx *Context) SetSession(key string, data []byte) {
	ctx.Server.SessionStorage.SetSession(ctx.GetSessionID(), key, data)
}

func (ctx *Context) GetSession(key string) []byte {
	return ctx.Server.SessionStorage.GetSession(ctx.GetSessionID(), key)
}

func (ctx *Context) ClearSession(key string) {
	ctx.Server.SessionStorage.ClearSession(ctx.GetSessionID(), key)
}

func (ctx *Context) setNewSessionID() (sessionID string) {
	sessionID = newSessionID()
	_timeout := int64(ctx.Config.SessionTimeout) * int64(time.Minute.Seconds())
	ctx.SetExpireCookie(SessionKey, sessionID, _timeout)
	return
}

// SetCookie adds a cookie header to the response.
func (ctx *Context) GetSessionID() (sessionID string) {
	var cookieValue string
	ctx.GetCookie(SessionKey, &cookieValue)

	if cookieValue == "" {
		return ctx.setNewSessionID()
	}
	return cookieValue
}
