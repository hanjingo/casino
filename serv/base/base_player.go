package base

import (
	"time"

	"github.com/hanjingo/casino/serv/lib"
)

type Player struct {
	Id         int64
	Sex        bool //true 男 false 女
	Status     int
	Name       string
	CreateTime time.Time
	Session    *lib.Session
}

//构造函数
func (p *Player) Player() {
	p.Id = 0
	p.Sex = false
	p.Status = 0
	p.Name = ""
	p.Session = nil
}

func (p *Player) GetId() interface{} {
	return p.Id
}

func (p *Player) IsValid() bool {
	//todo
	return true
}
