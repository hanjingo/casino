package lib

import "time"

type PlayerBase struct {
	id         int64
	sex        bool
	status     int
	name       string
	createTime time.Time
}

func (base *PlayerBase) GetId() int64 {
	return base.id
}

func (base *PlayerBase) SetId(id int64) {
	base.id = id
}

func (base *PlayerBase) GetSex() bool {
	return base.sex
}

func (base *PlayerBase) SetSex(sex bool) {
	base.sex = sex
}

func (base *PlayerBase) GetStatus() int {
	return base.status
}

func (base *PlayerBase) SetStatus(status int) {
	base.status = status
}

func (base *PlayerBase) GetName() string {
	return base.name
}

func (base *PlayerBase) SetName(name string) {
	base.name = name
}

func (base *PlayerBase) GetCreateTime() time.Time {
	return base.createTime
}

func (base *PlayerBase) SetCreateTime(t time.Time) {
	base.createTime = t
}
