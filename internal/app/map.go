package app

import (
	"proxxy/internal/client"
	"sync"
)

type mapItem struct {
	rq  client.Request
	rsp client.Response
}

//NOTE: use sync.Map for 64-core, google: cache contention
type appMap struct {
	mx sync.RWMutex
	m  map[int64]mapItem
}

func newAppMap() *appMap {
	return &appMap{
		m: make(map[int64]mapItem),
	}
}

/*
func (am *appMap) load(id int64) (mapItem, bool) {
	am.mx.RLock()
	val, ok := am.m[id]
	am.mx.RUnlock()
	return val, ok
}
*/
func (am *appMap) store(key int64, mi mapItem) {
	am.mx.Lock()
	am.m[key] = mi
	am.mx.Unlock()
}
