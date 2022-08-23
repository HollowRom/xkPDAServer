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
