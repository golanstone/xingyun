package xingyun

import (
	"net/http"
	"time"
)

func (ctx *Context) SetCookie(name string, value interface{}) {
	ctx.SetExpireCookie(name, value, 0)
}

func (ctx *Context) SetExpireCookie(name string, value interface{}, sec int64) {
	ctx.checkHeaderWrite()
	cookier := ctx.Server.SecureCookie
	encoded, err := cookier.Encode(name, value)
	if err != nil {
		ctx.Logger.Errorf(err.Error())
		return
	}
	cookie := &http.Cookie{
		Name:  name,
		Value: encoded,
		Path:  "/",
	}
	if sec != 0 {
		cookie.Expires = time.Unix(time.Now().Unix()+sec, 0)
	}
	if ctx.Config.CookieDomain != "" {
		cookie.Domain = ctx.Config.CookieDomain
	}
	http.SetCookie(ctx.ResponseWriter, cookie)
}

func (ctx *Context) GetCookie(name string, value interface{}) error {
	cookier := ctx.Server.SecureCookie
	r := ctx.Request
	cookie, err := r.Cookie(name)
	if err != nil {
		return err
	}
	return cookier.Decode(name, cookie.Value, value)
}

func (ctx *Context) GetStringCookie(name string) (value string, err error) {
	var s string
	err = ctx.GetCookie(name, &s)
	return s, err
}

func (ctx *Context) RemoveCookie(name string) {
	ctx.SetExpireCookie(name, "", -1)
}
