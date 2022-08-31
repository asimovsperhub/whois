package controller

import (
	"context"
	v1 "dewhois/api/v1"
	"dewhois/internal/service"
	"encoding/json"
)

var (
	Search = cSearch{}
)

type cSearch struct{}

func (c *cSearch) Search(ctx context.Context, req *v1.SearchReq) (res *v1.SearchRes, err error) {
	res = &v1.SearchRes{}
	sData := service.Search().Search(req.Key)
	j, _ := json.Marshal(sData)
	//fmt.Println(string(j))
	res.Result = string(j)
	return
}
