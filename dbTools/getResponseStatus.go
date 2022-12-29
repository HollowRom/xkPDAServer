package dbTools

import (
	"fmt"
	"github.com/tidwall/gjson"
)

/*
{
  "name": {"first": "Tom", "last": "Anderson"},
  "age":37,
  "children": ["Sara","Alex","Jack"],
  "fav.movie": "Deer Hunter",
  "friends": [
    {"first": "Dale", "last": "Murphy", "age": 44, "nets": ["ig", "fb", "tw"]},
    {"first": "Roger", "last": "Craig", "age": 68, "nets": ["fb", "tw"]},
    {"first": "Jane", "last": "Murphy", "age": 47, "nets": ["ig", "tw"]}
  ]
}
"name.last"          >> "Anderson"
"age"                >> 37
"children"           >> ["Sara","Alex","Jack"]
"children.#"         >> 3
"children.1"         >> "Alex"
"child*.2"           >> "Jack"
"c?ildren.0"         >> "Sara"
"fav\.movie"         >> "Deer Hunter"
"friends.#.first"    >> ["Dale","Roger","Jane"]
"friends.1.last"     >> "Craig"
*/

func isSuccess(json *gjson.Result) bool {
	if json == nil {
		return false
	}
	return json.Get("Result.ResponseStatus.IsSuccess").Bool()
}

func getReBillNo(json *gjson.Result) []string {
	if json == nil || !isSuccess(json) || !json.Get("Result.ResponseStatus.SuccessEntitys").Exists() || json.Get("Result.ResponseStatus.SuccessEntitys.#").Int() == 0 {
		return nil
	}
	var returnNumberList []string
	reList := json.Get("Result.ResponseStatus.SuccessEntitys.#.Number").Array()
	for _, v := range reList {
		returnNumberList = append(returnNumberList, v.Str)
	}
	return returnNumberList
}

func getReBillId(json *gjson.Result) []int {
	if json == nil || !isSuccess(json) || !json.Get("Result.ResponseStatus.SuccessEntitys").Exists() || json.Get("Result.ResponseStatus.SuccessEntitys.#").Int() == 0 {
		return nil
	}
	var returnNumberList []int
	reList := json.Get("Result.ResponseStatus.SuccessEntitys.#.Id").Array()
	for _, v := range reList {
		returnNumberList = append(returnNumberList, int(v.Int()))
	}
	return returnNumberList
}

func getErrMess(json *gjson.Result) string {
	if isSuccess(json) {
		return ""
	}
	errs := json.Get("Result.ResponseStatus.Errors.#.Message").Array()
	reStr := ""
	for _, e := range errs {
		reStr = reStr + e.Str + "\r\n"
	}
	return reStr
}

type ResponseStatus struct {
	Json *gjson.Result
}

func (r *ResponseStatus) GetErrMess() string {
	return getErrMess(r.Json)
}

func (r *ResponseStatus) ToString() string {
	return fmt.Sprintf("%s", r.Json)
}

func (r *ResponseStatus) SetStr(json string) {
	if r.Json != nil {
		return
	}
	j := gjson.Parse(json)
	r.Json = &j
}

func (r *ResponseStatus) IsSuccess() bool {
	return isSuccess(r.Json)
}

func (r *ResponseStatus) GetReBillNo() []string {
	return getReBillNo(r.Json)
}

func (r *ResponseStatus) GetReBillId() []int {
	return getReBillId(r.Json)
}

//func (r *ResponseStatus) isSuccess() bool {
//	if r == nil {
//		return false
//	}
//	return (r.Map["Result"]).(map[string]interface{})["ResponseStatus"].(map[string]interface{})["isSuccess"].(bool)
//}
//
//func (r *ResponseStatus) getReBillNo() []string {
//	if r == nil || (r.Map["Result"]).(map[string]interface{})["ResponseStatus"].(map[string]interface{})["SuccessEntitys"] == nil || len((r.Map["Result"]).(map[string]interface{})["ResponseStatus"].(map[string]interface{})["SuccessEntitys"].([]map[string]interface{})) == 0 {
//		return nil
//	}
//	var returnNumberList []string
//	reList := (r.Map["Result"]).(map[string]interface{})["ResponseStatus"].(map[string]interface{})["SuccessEntitys"].([]map[string]interface{})
//	for _, v := range reList {
//		returnNumberList = append(returnNumberList, v["Number"].(string))
//	}
//	return returnNumberList
//}
//
//func (r *ResponseStatus) getReBillId() []int {
//	if r == nil || (r.Map["Result"]).(map[string]interface{})["ResponseStatus"].(map[string]interface{})["SuccessEntitys"] == nil || len((r.Map["Result"]).(map[string]interface{})["ResponseStatus"].(map[string]interface{})["SuccessEntitys"].([]map[string]interface{})) == 0 {
//		return nil
//	}
//	var returnNumberList []int
//	reList := (r.Map["Result"]).(map[string]interface{})["ResponseStatus"].(map[string]interface{})["SuccessEntitys"].([]map[string]interface{})
//	for _, v := range reList {
//		//it, err := strconv.Atoi(v.Id)
//		//if err != nil {
//		//	return nil
//		//}
//		//returnNumberList = append(returnNumberList, it)
//		returnNumberList = append(returnNumberList, v["Id"].(int))
//	}
//	return returnNumberList
//}
//------------------------

//type ResponseStatus struct {
//	Result struct {
//		ResponseStatus struct {
//			isSuccess      bool          `json:"isSuccess"`
//			Errors         []interface{} `json:"Errors"`
//			SuccessEntitys []*struct {
//				Id     int    `json:"Id.,string"`
//				Number string `json:"Number"`
//				DIndex int    `json:"DIndex"`
//			} `json:"SuccessEntitys"`
//			SuccessMessages []interface{} `json:"SuccessMessages"`
//			MsgCode         int           `json:"MsgCode"`
//		} `json:"ResponseStatus"`
//		Id             int    `json:"Id,,string"`
//		Number         string `json:"Number"`
//		NeedReturnData []struct {
//		} `json:"NeedReturnData"`
//	} `json:"Result"`
//}
//
//func (r *ResponseStatus) isSuccess() bool {
//	if r == nil {
//		return false
//	}
//	return r.Result.ResponseStatus.isSuccess
//}
//
//func (r *ResponseStatus) getReBillNo() []string {
//	if r == nil || r.Result.ResponseStatus.SuccessEntitys == nil || len(r.Result.ResponseStatus.SuccessEntitys) == 0 {
//		return nil
//	}
//	var returnNumberList []string
//	reList := r.Result.ResponseStatus.SuccessEntitys
//	for _, v := range reList {
//		returnNumberList = append(returnNumberList, v.Number)
//	}
//	return returnNumberList
//}
//
//func (r *ResponseStatus) getReBillId() []int {
//	if r == nil || r.Result.ResponseStatus.SuccessEntitys == nil || len(r.Result.ResponseStatus.SuccessEntitys) == 0 {
//		return nil
//	}
//	var returnNumberList []int
//	reList := r.Result.ResponseStatus.SuccessEntitys
//	for _, v := range reList {
//		//it, err := strconv.Atoi(v.Id)
//		//if err != nil {
//		//	return nil
//		//}
//		//returnNumberList = append(returnNumberList, it)
//		returnNumberList = append(returnNumberList, v.Id)
//	}
//	return returnNumberList
//}
