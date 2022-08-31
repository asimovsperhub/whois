package v1

import "github.com/gogf/gf/v2/frame/g"

type SearchReq struct {
	g.Meta `path:"/api/search" tags:"Search" method:"get" summary:"搜索"`
	Key    string `json:"key" v:"required#请输入搜索关键字" dc:"关键字"`
}
type SearchRes struct {
	Result string `json:"result" dc:"result"`
}
