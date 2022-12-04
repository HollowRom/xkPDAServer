package netTools

import (
	"crypto/md5"
	"sync"
	"time"
)

type cacheMap struct {
	*sync.RWMutex
	c map[string]bool
}

var defTimeOut time.Duration = 60 * 1000

var defCacheMap = &cacheMap{c: map[string]bool{}}

func AddCache(key string) {
	defCacheMap.Lock()
	defer defCacheMap.Unlock()
	k := strToMd5(key)
	defCacheMap.c[k] = true
	defer deleteKey(k)
}

func GetCache(key string) bool {
	defCacheMap.RLock()
	defer defCacheMap.RUnlock()
	_, ok := defCacheMap.c[strToMd5(key)]
	return ok
}

func deleteKey(key string) {
	time.Sleep(defTimeOut)
	defCacheMap.Lock()
	defer defCacheMap.Unlock()
	delete(defCacheMap.c, key)
}

func strToMd5(key string) string {
	var o = md5.Sum([]byte(key))
	var n = make([]byte, 0, 16)
	for _, b := range o {
		n = append(n, b)
	}
	return string(n)
}
