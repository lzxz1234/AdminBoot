package utils

// Result 实体
type Result struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
	Count int         `json:"count"`
	Page  int         `json:"page"`
	Limit int         `json:"limit"`
}

// NewPage to create Page
func NewPage(count int, pageNo int, pageSize int, list interface{}) Result {

	return Result{Code: 0, Msg: "请求成功", Data: list, Count: count, Page: pageNo, Limit: pageSize}
}

// NewResult 构造结果实体
func NewResult(code int, msg string, data interface{}) Result {
	return Result{Code: code, Msg: msg, Data: data}
}
