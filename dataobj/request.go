package dataobj

import "time"

type RequestItem struct {
	Url string `json:"url"`
	ReqTime time.Time `json:"req_time"`
	IsCheck bool `json:"is_check"`
	Deepth int8 `json:"deepth"`
	Domain string `json:"domain"`
}

type ReponseItem struct{
	RequestUrl string `json:"request_url"`
	RespCode int `json:"resp_code"`
	RespTime time.Duration `json:"resp_time"`
	RespUrls []RequestItem `json:"resp_urls"`
	IsCheck bool `json:"is_check"`
	Deepth int8 `json:"deepth"`
	Domain string `json:"domain"`
}
