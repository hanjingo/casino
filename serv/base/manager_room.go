package base

import (
	"errors"
	"sync"
	"time"
)

type RoomI interface {
	GetId() interface{}
	IsValid() bool
	Close()
}

type RoomManager struct {
	Mu   *sync.RWMutex
	M    map[interface{}]RoomI
	BRun bool
}

//构造函数
func (rm *RoomManager) RoomManager() {
	rm.Mu = new(sync.RWMutex)
	rm.M = make(map[interface{}]RoomI)
	rm.BRun = false
}

//初始化
func (rm *RoomManager) Init() {

}

//跑起来
func (rm *RoomManager) Run(dur time.Duration) {
	if rm.BRun {
		return
	}
	rm.BRun = true
	go func() {
		time.Sleep(dur)
		for id, room := range rm.M {
			if !room.IsValid() {
				room.Close()
			}
			rm.DelRoom(id)
		}
	}()
}

//拿房子
func (rm *RoomManager) GetRoom(id interface{}) RoomI {
	rm.Mu.Lock()
	defer rm.Mu.Unlock()
	if rm.M == nil {
		return nil
	}
	//先找内存
	if room, ok := rm.M[id]; ok {
		return room
	}
	//todo 找磁盘
	return nil
}

//放房子
func (rm *RoomManager) AddRoom(r RoomI) error {
	if r == nil {
		return errors.New("房子不能为空")
	}
	rm.Mu.Lock()
	defer rm.Mu.Unlock()
	if rm.M == nil {
		rm.M = make(map[interface{}]RoomI)
	}
	rm.M[r.GetId()] = r
	return nil
}

//减房
func (rm *RoomManager) DelRoom(id interface{}) {
	rm.Mu.Lock()
	defer rm.Mu.Unlock()
	if rm.M == nil {
		return
	}
	delete(rm.M, id)
}
