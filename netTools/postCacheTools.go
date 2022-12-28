package netTools

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
	"time"
)

type cacheMap struct {
	rw *sync.RWMutex
	c  map[string]bool
}

var defTimeOut time.Duration = 60 * 1000

var defCacheMap = &cacheMap{rw: &sync.RWMutex{}, c: map[string]bool{}}

func AddCache(key string) {
	defCacheMap.rw.Lock()
	defer defCacheMap.rw.Unlock()
	k := strToMd5(key)
	defCacheMap.c[k] = true
	go deleteKey(k)
}

//func GetCache(key string) bool {
//	defCacheMap.rw.RLock()
//	defer defCacheMap.rw.RUnlock()
//	_, ok := defCacheMap.c[strToMd5(key)]
//	return ok
//}

func GetCacheWithCTX(key string, ctx *gin.Context) bool {
	defCacheMap.rw.RLock()
	defer defCacheMap.rw.RUnlock()
	_, ok := defCacheMap.c[strToMd5(key)]
	if ok {
		ctx.JSON(http.StatusBadRequest, &gin.H{"err": fmt.Sprintf("重复提交数据,%d秒后重试", defTimeOut/1000)})
		fmt.Println(fmt.Sprintf("重复提交数据,%d秒后重试", defTimeOut/1000))
	}

	return ok
}

func deleteKey(key string) {
	time.Sleep(defTimeOut)
	defCacheMap.rw.Lock()
	defer defCacheMap.rw.Unlock()
	delete(defCacheMap.c, key)
}

func strToMd5(key string) string {
	var o = md5.Sum([]byte(key))
	return fmt.Sprintf("%x", o)
	//var n = make([]byte, 0, 16)
	//for _, b := range o {
	//	n = append(n, b)
	//}
	//return string(n)
}
