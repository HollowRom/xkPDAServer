package netTools

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
	"xkpdaserver/dbTools"
)

const (
	KDKey  = "kdservice-sessionid"
	ASPKey = "ASP.NET_SessionId"

	jumpApiUrl = "http://127.0.0.1:8080/jumpApi/"
)

var (
	loginUrl          = "/k3cloud/Kingdee.BOS.WebApi.ServicesStub.AuthService.ValidateUser.common.kdsvc"
	saveBillUrlTail   = "/k3cloud/Kingdee.BOS.WebApi.ServicesStub.DynamicFormService.Save.common.kdsvc"
	pushBillUrlTail   = "/k3cloud/Kingdee.BOS.WebApi.ServicesStub.DynamicFormService.Push.common.kdsvc"
	selectBillUrlTail = "/k3cloud/Kingdee.BOS.WebApi.ServicesStub.DynamicFormService.View.common.kdsvc"
	defHost           = "127.0.0.1"
)

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

const versionFilePath = ".\\version\\version"

type ObjectVersionFieldJson struct {
	Version string           `json:"version"`
	Data    *ObjectFieldData `json:"data"`
}

type ObjectFieldData struct {
	ListName []string        `json:"listName"`
	List     []*ObjectFields `json:"list"`
}

type ObjectFields struct {
	ObjectName string         `json:"objectName"`
	Fields     []*ObjectField `json:"fields"`
}

type ObjectField struct {
	OrderIdx    int    `json:"orderIdx"`
	BodyTitle   string `json:"bodyTitle"`
	IsEdit      bool   `json:"isEdit"`
	IsHidden    bool   `json:"isHidden"`
	IsCheckNum  bool   `json:"isCheckNum"`
	IsPostField bool   `json:"isPostField"`
	Width       int    `json:"width"`
}

func setDefObjectFieldJson(o *ObjectVersionFieldJson) {
	if o.Data != nil && o.Data.List != nil {
		t := o.Data.List
		for _, of := range t {
			if of != nil && of.Fields != nil {
				for _, off := range of.Fields {
					if off.Width == 0 {
						off.Width = 300
					}
				}
			}
		}
	}

	//return &ObjectField{
	//	FieldName:   "",
	//	OrderIdx:    0,
	//	BodyTitle:   "",
	//	IsEdit:      false,
	//	IsHidden:    true,
	//	IsCheckNum:  false,
	//	IsPostField: false,
	//	Width:       300,
	//}
}

func ReadVersionFieldJson() (version string, fieldJson *ObjectFieldData) {
	s, e := os.ReadFile(versionFilePath)
	if e != nil {
		return "", nil
	}
	parseStruct := &ObjectVersionFieldJson{}
	e = json.Unmarshal(s, parseStruct)
	if e != nil {
		return "", nil
	}
	setDefObjectFieldJson(parseStruct)
	//???????????????
	return parseStruct.Version, parseStruct.Data
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
		defHost = tempValue
	}
	loginUrl = "http://" + defHost + loginUrl
	saveBillUrlTail = "http://" + defHost + saveBillUrlTail
	selectBillUrlTail = "http://" + defHost + selectBillUrlTail
	pushBillUrlTail = "http://" + defHost + pushBillUrlTail
	fmt.Println("?????????????????????:", *defLoginBase)
	fmt.Println("?????????????????????????????????")
	if !tryLogin(nil) {
		panic("????????????????????????")
	}

	//test()
	//
	//panic("????????????")
}

func GetLoginUrl() string {
	return loginUrl
}
func GetSaveBillUrl() string {
	return saveBillUrlTail
}
func GetSelectBillUrl() string {
	return selectBillUrlTail
}
func GetPushBillUrl() string {
	return pushBillUrlTail
}

func GetJumpApi() string {
	return jumpApiUrl
}

func Init() {
	o.Do(oneInit)
}

var client = &http.Client{}

type LoginCookie struct {
	Kdservice string
	ASP       string
}

func tryLogin(b *LoginBase) bool {
	if b == nil {
		b = defLoginBase
	}

	j, e := json.Marshal(b)
	if e != nil {
		fmt.Println(e)
		return false
	}

	req, _ := http.NewRequest("POST", loginUrl, bytes.NewBuffer(j))

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
		fmt.Println("????????????:" + e.Error())
		return false
	}

	if strings.Contains(string(buf[0:i]), "IsSuccessByAPI\":false") {
		fmt.Println("????????????:")
		fmt.Println(*b)
		return false
	}

	//fmt.Println("??????????????????:" + string(buf[0:i]))

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
		fmt.Println("??????cookie??????")
		return false
	}

	return true
}

func PostSaveSomeBill(jsonByte []byte) []byte {
	if jsonByte == nil || len(jsonByte) == 0 {
		return nil
	}
	return postSomeBill(GetSaveBillUrl(), jsonByte)
}

func PostPushSomeBill(jsonByte []byte) []byte {
	if jsonByte == nil || len(jsonByte) == 0 {
		return nil
	}
	return postSomeBill(GetPushBillUrl(), jsonByte)
}

func PostJumpApi(jsonByte []byte) []byte {
	if jsonByte == nil || len(jsonByte) == 0 {
		return nil
	}

	m := &map[string]interface{}{}

	e := json.Unmarshal(jsonByte, m)
	if e != nil {
		fmt.Println(e)
		return nil
	}

	if _, ok := (*m)["api"]; !ok {
		fmt.Println("?????????api??????")
		return nil
	}
	//11
	if _, ok := (*m)["jsp"]; !ok {
		fmt.Println("?????????jsp??????")
		return nil
	}

	m = nil

	return postSomeBill(GetJumpApi(), jsonByte)
}

func PostSelectSomeBill(jsonByte []byte) *map[string]interface{} {
	if jsonByte == nil || len(jsonByte) == 0 {
		return nil
	}
	reMap := &map[string]interface{}{}
	e := json.Unmarshal(postSomeBill(GetSaveBillUrl(), jsonByte), reMap)
	if e != nil {
		fmt.Println(e)
		return nil
	}
	return reMap
}

func postSomeBill(postUrl string, jsonByte []byte) []byte {
	defCookieManger.rwLock.RLock()
	defer defCookieManger.rwLock.RUnlock()
	if len(defCookieManger.cookies) == 0 {
		defCookieManger.rwLock.RUnlock()
		tryLogin(nil)
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
	fmt.Println(randNum, ":post??????:"+string(jsonByte))
	fmt.Println(randNum, ":post?????????:"+string(body))
	return body
}

//func GetSome(getUrl string, jsonByte []byte) []byte {
//	if getUrl == "" || jsonByte == nil || len(jsonByte) == 0 {
//		return nil
//	}
//	defCookieManger.rwLock.RLock()
//	defer defCookieManger.rwLock.RUnlock()
//	if len(defCookieManger.cookies) == 0 {
//		defCookieManger.rwLock.RUnlock()
//		tryLogin(nil)
//		defCookieManger.rwLock.RLock()
//	}
//
//	req, e := http.NewRequest("GET", getUrl, bytes.NewBuffer(jsonByte))
//	if e != nil {
//		fmt.Println(e)
//		return nil
//	}
//
//	req.Header.Set("Content-Type", "application/json")
//
//	for _, c := range defCookieManger.cookies {
//		req.AddCookie(c)
//	}
//
//	resp, e := client.Do(req)
//	if e != nil {
//		fmt.Println(e)
//		return nil
//	}
//	defer func(Body io.ReadCloser) {
//		err := Body.Close()
//		if err != nil {
//			fmt.Println(err)
//		}
//	}(resp.Body)
//
//	body, e := ioutil.ReadAll(resp.Body)
//	if e != nil {
//		fmt.Println(e)
//		return nil
//	}
//	return body
//}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func getRandNumber() uint64 {
	return rand.Uint64()
}
