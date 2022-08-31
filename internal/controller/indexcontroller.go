package controller

import (
	"context"
	v1 "dewhois/api/v1"
	"dewhois/internal/model"
	"dewhois/internal/service"
)

var (
	Index = cIndex{}
)

type cIndex struct {
}

func (*cIndex) Index(ctx context.Context, req *v1.IndexReq) (res *v1.IndexRes, err error) {
	service.View().Render(ctx, model.View{
		Title: "首页",
	})
	return
}
