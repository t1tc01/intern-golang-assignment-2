package model

type ApiReq struct {
	ID          int
	ReqTime     int64  `json:"time"`
	ReqParams   string `json:"req_param"`
	ReqBody     string `json:"req_body"`
	ReqHeaders  string `json:"req_headers"`
	RespTime    int64  `json:"resp_time"`
	RespStatus  int    `json:"resp_status"`
	RespBody    string `json:"resp_body"`
	RespHeaders string `json:"resp_headers"`
	CreatedAt   int64  `json:"createdTime"`
	UpdatedAt   int64  `json:"updatedTime"`
	DeletedAt   int64  `json:"deletedTime"`
}
