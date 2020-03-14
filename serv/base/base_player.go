package base

import (
	"github.com/hanjingo/casino/serv/lib"
	"time"
)

type Player struct {
	Id         int64
	Sex        bool
	Status     int
	Name       string
	CreateTime time.Time
	Session *lib.Session
}

func (p *Player) GetId() interface{} {
	return p.Id
}

func (p *Player) IsValid() bool {
	//todo
	return true
}
