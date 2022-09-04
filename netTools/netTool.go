package netTools

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"time"
	"xkpdaserver/dbTools"
)

const (
	KDKey  = "kdservice-sessionid"
	ASPKey = "ASP.NET_SessionId"
)

var xkLoginUrl = "http://192.168.31.153/k3cloud/Kingdee.BOS.WebApi.ServicesStub.AuthService.ValidateUser.common.kdsvc"

type cookiesManger struct {
	rwLock  *sync.RWMutex
	cookies []*http.Cookie
}

var defCookieManger = cookiesManger{
	rwLock:  &sync.RWMutex{},
	cookies: []*http.Cookie{},
}

type LoginBase struct {
	AcctID   string `json:"acctID"`
	Username string `json:"username"`
	Password string `json:"password"`
	Lcid     int    `json:"lcid"`
}

//select FDATACENTERID from T_BAS_DATACENTER
var defLoginBase = &LoginBase{"627cf0721296c7", "administrator", "kingdee@123", 2052}

func GetConfListenPort() string {
	return dbTools.GetConfFromKey("listenPort")
}

var o sync.Once

var oneInit = func() {
	tempValue := dbTools.GetConfFromKey("acctid")
	if tempValue != "" {
		defLoginBase.AcctID = tempValue
	}
	tempValue = dbTools.GetConfFromKey("username")
	if tempValue != "" {
		defLoginBase.Username = tempValue
	}
	tempValue = dbTools.GetConfFromKey("password")
	if tempValue != "" {
		defLoginBase.Password = tempValue
	}
	tempValue = dbTools.GetConfFromKey("ServerIp")
	if tempValue != "" {
		xkLoginUrl = "http://" + tempValue + "/k3cloud/Kingdee.BOS.WebApi.ServicesStub.AuthService.ValidateUser.common.kdsvc"
	}
	fmt.Println("读取登录信息为:", *defLoginBase)
	fmt.Println("星空登陆数据初始化完成")
}

func Init() {
	o.Do(oneInit)
}

var client = &http.Client{}

type LoginCookie struct {
	Kdservice string
	ASP       string
}

func TryLogin(b *LoginBase) bool {
	if b == nil {
		b = defLoginBase
	}

	j, e := json.Marshal(b)
	if e != nil {
		fmt.Println(e)
		return false
	}

	req, _ := http.NewRequest("POST", xkLoginUrl, bytes.NewBuffer(j))

	req.Header.Set("Content-Type", "application/json")

	resp, e := client.Do(req)
	if e != nil {
		fmt.Println(e)
		return false
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)

	buf := make([]byte, 10240)
	i, e := resp.Body.Read(buf)
	if e != nil && e != io.EOF {
		fmt.Println("登录失败:" + e.Error())
		return false
	}

	if strings.Contains(string(buf[0:i]), "IsSuccessByAPI\":false") {
		fmt.Println("登录失败:")
		fmt.Println(*b)
		return false
	}

	//fmt.Println("登录返回信息:" + string(buf[0:i]))

	defCookieManger.rwLock.Lock()
	defer defCookieManger.rwLock.Unlock()

	defCookieManger.cookies = []*http.Cookie{}

	for _, c := range resp.Cookies() {
		if c.Name == KDKey {
			defCookieManger.cookies = append(defCookieManger.cookies, c)
		}
		if c.Name == ASPKey {
			defCookieManger.cookies = append(defCookieManger.cookies, c)
		}
	}

	if len(defCookieManger.cookies) == 0 {
		fmt.Println("获取cookie失败")
		return false
	}

	return true
}

func PostSome(postUrl string, jsonByte []byte) []byte {
	if postUrl == "" || jsonByte == nil || len(jsonByte) == 0 {
		return nil
	}
	defCookieManger.rwLock.RLock()
	defer defCookieManger.rwLock.RUnlock()
	if len(defCookieManger.cookies) == 0 {
		defCookieManger.rwLock.RUnlock()
		TryLogin(nil)
		defCookieManger.rwLock.RLock()
	}

	req, e := http.NewRequest("POST", postUrl, bytes.NewBuffer(jsonByte))
	if e != nil {
		fmt.Println(e)
		return nil
	}

	req.Header.Set("Content-Type", "application/json")

	for _, c := range defCookieManger.cookies {
		req.AddCookie(c)
	}

	resp, e := client.Do(req)
	if e != nil {
		fmt.Println(e)
		return nil
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)

	body, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		fmt.Println(e)
		return nil
	}

	randNum := getRandNumber()
	fmt.Println(randNum, ":post发送:"+string(jsonByte))
	fmt.Println(randNum, ":post后返回:"+string(body))
	return body
}

func GetSome(postUrl string, jsonByte []byte) []byte {
	if postUrl == "" || jsonByte == nil || len(jsonByte) == 0 {
		return nil
	}
	defCookieManger.rwLock.RLock()
	defer defCookieManger.rwLock.RUnlock()
	if len(defCookieManger.cookies) == 0 {
		defCookieManger.rwLock.RUnlock()
		TryLogin(nil)
		defCookieManger.rwLock.RLock()
	}

	req, e := http.NewRequest("GET", postUrl, bytes.NewBuffer(jsonByte))
	if e != nil {
		fmt.Println(e)
		return nil
	}

	req.Header.Set("Content-Type", "application/json")

	for _, c := range defCookieManger.cookies {
		req.AddCookie(c)
	}

	resp, e := client.Do(req)
	if e != nil {
		fmt.Println(e)
		return nil
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)

	body, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		fmt.Println(e)
		return nil
	}
	return body
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func getRandNumber() uint64 {
	return rand.Uint64()
}
