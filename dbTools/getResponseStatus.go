package dbTools


type ResponseStatus struct {
	Map map[string]interface{}
}

func (r *ResponseStatus) IsSuccess() bool {
	if r == nil {
		return false
	}
	return (r.Map["Result"]).(map[string]interface{})["ResponseStatus"].(map[string]interface{})["IsSuccess"].(bool)
}

func (r *ResponseStatus) GetReBillNo() []string {
	if r == nil || (r.Map["Result"]).(map[string]interface{})["ResponseStatus"].(map[string]interface{})["SuccessEntitys"] == nil || len((r.Map["Result"]).(map[string]interface{})["ResponseStatus"].(map[string]interface{})["SuccessEntitys"].([]map[string]interface{})) == 0 {
		return nil
	}
	var returnNumberList []string
	reList := (r.Map["Result"]).(map[string]interface{})["ResponseStatus"].(map[string]interface{})["SuccessEntitys"].([]map[string]interface{})
	for _, v := range reList {
		returnNumberList = append(returnNumberList, v["Number"].(string))
	}
	return returnNumberList
}

func (r *ResponseStatus) GetReBillId() []int {
	if r == nil || (r.Map["Result"]).(map[string]interface{})["ResponseStatus"].(map[string]interface{})["SuccessEntitys"] == nil || len((r.Map["Result"]).(map[string]interface{})["ResponseStatus"].(map[string]interface{})["SuccessEntitys"].([]map[string]interface{})) == 0 {
		return nil
	}
	var returnNumberList []int
	reList := (r.Map["Result"]).(map[string]interface{})["ResponseStatus"].(map[string]interface{})["SuccessEntitys"].([]map[string]interface{})
	for _, v := range reList {
		//it, err := strconv.Atoi(v.Id)
		//if err != nil {
		//	return nil
		//}
		//returnNumberList = append(returnNumberList, it)
		returnNumberList = append(returnNumberList, v["Id"].(int))
	}
	return returnNumberList
}

//type ResponseStatus struct {
//	Result struct {
//		ResponseStatus struct {
//			IsSuccess      bool          `json:"IsSuccess"`
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
//func (r *ResponseStatus) IsSuccess() bool {
//	if r == nil {
//		return false
//	}
//	return r.Result.ResponseStatus.IsSuccess
//}
//
//func (r *ResponseStatus) GetReBillNo() []string {
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
//func (r *ResponseStatus) GetReBillId() []int {
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
