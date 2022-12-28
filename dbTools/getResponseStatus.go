package dbTools

type ResponseStatus struct {
	Result struct {
		ResponseStatus struct {
			IsSuccess      bool          `json:"IsSuccess"`
			Errors         []interface{} `json:"Errors"`
			SuccessEntitys []struct {
				Id     int    `json:"Id"`
				Number string `json:"Number"`
				DIndex int    `json:"DIndex"`
			} `json:"SuccessEntitys"`
			SuccessMessages []interface{} `json:"SuccessMessages"`
			MsgCode         int           `json:"MsgCode"`
		} `json:"ResponseStatus"`
		Id             int    `json:"Id"`
		Number         string `json:"Number"`
		NeedReturnData []struct {
		} `json:"NeedReturnData"`
	} `json:"Result"`
}

func (r *ResponseStatus) IsSuccess() bool {
	if r == nil {
		return false
	}
	return r.Result.ResponseStatus.IsSuccess
}

func (r *ResponseStatus) GetReBillNo() []string {
	if r == nil || r.Result.ResponseStatus.SuccessEntitys == nil || len(r.Result.ResponseStatus.SuccessEntitys) == 0 {
		return nil
	}
	var returnNumberList []string
	reList := r.Result.ResponseStatus.SuccessEntitys
	for _, v := range reList {
		returnNumberList = append(returnNumberList, v.Number)
	}
	return returnNumberList
}

func (r *ResponseStatus) GetReBillId() []int {
	if r == nil || r.Result.ResponseStatus.SuccessEntitys == nil || len(r.Result.ResponseStatus.SuccessEntitys) == 0 {
		return nil
	}
	var returnNumberList []int
	reList := r.Result.ResponseStatus.SuccessEntitys
	for _, v := range reList {
		//it, err := strconv.Atoi(v.Id)
		//if err != nil {
		//	return nil
		//}
		returnNumberList = append(returnNumberList, v.Id)
	}
	return returnNumberList
}
