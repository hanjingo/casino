package player_manager

import (
	"sync"
)

type Player1I interface {
	GetId() int
	IsOffline() bool
}

type PlayerManager1 struct {
	mu *sync.Mutex
	m  map[interface{}]Player1I
}

//初始化
func (pm *PlayerManager1) Init() {
	//todo
}

//维护
func (pm *PlayerManager1) Run() {
	for id, p := range pm.m {
		if p.IsOffline() {
			pm.DelPlayer(id)
		}
	}
}

//拿人
func (pm *PlayerManager1) GetPlayer(uid interface{}) Player1I {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	//拿内存
	if p, ok := pm.m[uid]; ok {
		return p
	}
	return nil
}

//加人
func (pm *PlayerManager1) AddPlayer(p Player1I) {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	if pm.m == nil {
		pm.m = make(map[interface{}]Player1I)
	}
	if p == nil {
		return
	}
	pm.m[p.GetId()] = p
}

//减人
func (pm *PlayerManager1) DelPlayer(uid interface{}) {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	if pm.m == nil {
		return
	}
	delete(pm.m, uid)
}
