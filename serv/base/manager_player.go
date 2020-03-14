package base

import (
	"errors"
	"sync"
	"time"
)

type PlayerI interface {
	GetId() interface{}
	IsValid() bool
}

type PlayerManager struct {
	Mu   *sync.Mutex
	M    map[interface{}]PlayerI
	BRun bool
}

//初始化
func (pm *PlayerManager) Init() {
	pm.Mu = new(sync.Mutex)
	pm.M = make(map[interface{}]PlayerI)
	pm.BRun = false
}

//维护
func (pm *PlayerManager) Run(dur time.Duration) {
	if pm.BRun {
		return
	}
	pm.BRun = true
	go func() {
		time.Sleep(dur)
		for id, p := range pm.M {
			if p.IsValid() {
				pm.DelPlayer(id)
			}
			pm.DelPlayer(id)
		}
	}()
}

//拿人
func (pm *PlayerManager) GetPlayer(uid interface{}) PlayerI {
	pm.Mu.Lock()
	defer pm.Mu.Unlock()
	if pm.M == nil {
		return nil
	}
	//拿内存
	if p, ok := pm.M[uid]; ok {
		return p
	}
	//拿db  todo
	return nil
}

//加人
func (pm *PlayerManager) AddPlayer(p PlayerI) error {
	pm.Mu.Lock()
	defer pm.Mu.Unlock()
	if pm.M == nil {
		pm.M = make(map[interface{}]PlayerI)
	}
	if p == nil {
		return errors.New("传入的玩家为空")
	}
	pm.M[p.GetId()] = p
	return nil
}

//减人
func (pm *PlayerManager) DelPlayer(uid interface{}) {
	pm.Mu.Lock()
	defer pm.Mu.Unlock()
	if pm.M == nil {
		return
	}
	delete(pm.M, uid)
}
