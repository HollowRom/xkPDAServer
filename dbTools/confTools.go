package dbTools

import (
	"fmt"
	"os"
	"strings"
)

var confMap = map[string]string{}

func GetConfFromKey(key string) string {
	return confMap[key]
}

func init() {
	c, e := os.ReadFile("conf.txt")
	if e != nil {
		panic(e)
		return
	}
	confSplit := strings.Split(string(c), "\r\n")
	for i := 0; i < len(confSplit); i++ {
		keyValue := strings.Split(confSplit[i], "-")
		if len(keyValue) > 1 && keyValue[0] != "" {
			confMap[keyValue[0]] = keyValue[1]
		}
	}
	fmt.Printf("配置文件读取完成，共%d项配置信息\n", len(confMap))
	fmt.Println(confMap)
}
