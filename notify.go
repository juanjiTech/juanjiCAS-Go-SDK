package juanjiCAS

import (
	"errors"
	"github.com/guonaihong/gout"
)

var (
	ErrNotSetSendKey  = errors.New("not set send key")
	ErrInvalidSendKey = errors.New("invalid send key")
	ErrUnknown        = errors.New("unknown error")
	ErrServerErr      = errors.New("server error")
)

func (c *CAS) Notify(notice *Notice) error {
	if c.sendKey == "" {
		return ErrNotSetSendKey
	}
	resp := Resp{}
	err := gout.POST(c.domain + "/api/notify/send/" + c.sendKey).
		SetJSON(notice).
		BindJSON(&resp).Do()
	if err != nil {
		return err
	}
	if resp.Code == 40100 {
		return ErrInvalidSendKey
	} else if resp.Code == 50001 {
		return ErrServerErr
	} else if resp.Code != 20000 {
		return ErrUnknown
	}
	return nil
}
