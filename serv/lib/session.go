package lib

import (
	"errors"

	"github.com/hanjingo/network"

	v1 "github.com/hanjingo/gate/cli/golang/v1"
)

type Session struct {
	Gate *v1.GateCli
	Id   uint64
}

func NewSession() *Session {
	back := &Session{}
	back.Gate = v1.NewGateCli()
	return back
}

func (session *Session) Bind(id uint64) {
	session.Id = id
}

func (session *Session) Dial(addr, token string, conf *network.SessionConfig) error {
	if session.Gate == nil {
		return errors.New("网关客户端为空")
	}
	return session.Gate.Dial(addr, token, conf)
}
